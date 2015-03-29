var Console = React.createClass({
	render: function() {
		return (
			<div className="console">
				<div className="log">
				</div>
				<textarea className="input">
				</textarea>
			</div>
		);
	}
});

var DBList = React.createClass({
	render: function() {
		return (
			<div className="db-list">
			</div>
		);
	}
});

var DataPanel = React.createClass({
	render: function() {
		return (
			<div className="data-panel">
			</div>
		);
	}
});

var App = React.createClass({
	render: function() {
		return (
			<div className="app">
				<DBList />
				<DataPanel />
				<Console />
			</div>
		);
	}
});

React.render(
	 <App />,
	document.getElementById('content')
);
