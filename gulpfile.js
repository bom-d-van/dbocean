var gulp = require('gulp');
var react = require('gulp-react');
var rename = require('gulp-rename');
var browserify = require('gulp-browserify');
var less = require('gulp-less');
var sass = require('gulp-sass')
// var autoprefixer = require('gulp-autoprefixer');

gulp.task('default', function () {
	gulp.watch('ui/app.jsx', ['app'])
	gulp.watch('ui/app.scss', ['app'])
});

gulp.task('app', function () {
	gulp.src(['ui/app.jsx'])
		.pipe(react())
		.pipe(rename('app.js'))
		.pipe(gulp.dest('ui/'));

	gulp.src(['ui/app.scss'])
		.pipe(sass())
		.pipe(rename('app.css'))
		.pipe(gulp.dest('ui/'));
});

gulp.task('lib', function() {
	gulp.src(['ui/lib.jsx'])
		.pipe(browserify()) .pipe(react())
		.pipe(rename('lib.js'))
		.pipe(gulp.dest('ui/'));

	gulp.src(['ui/lib.less'])
		.pipe(less())
		// .pipe(autoprefixer({cascade: false, browsers: ['last 2 versions']}))
		.pipe(rename('lib.css'))
		.pipe(gulp.dest('ui/'));
})
