var Console = React.createClass({displayName: "Console",
	render: function() {
		return (
			React.createElement("div", {className: "console"}, 
				React.createElement("div", {className: "log"}
				), 
				React.createElement("textarea", {className: "input"}
				)
			)
		);
	}
});

var DBList = React.createClass({displayName: "DBList",
	render: function() {
		return (
			React.createElement("div", {className: "db-list"}
			)
		);
	}
});

var DataPanel = React.createClass({displayName: "DataPanel",
	render: function() {
		return (
			React.createElement("div", {className: "data-panel"}
			)
		);
	}
});

var App = React.createClass({displayName: "App",
	render: function() {
		return (
			React.createElement("div", {className: "app"}, 
				React.createElement(DBList, null), 
				React.createElement(DataPanel, null), 
				React.createElement(Console, null)
			)
		);
	}
});

React.render(
	 React.createElement(App, null),
	document.getElementById('content')
);
