var Console = React.createClass({displayName: "Console",
	handleKeyUp: function(event) {
		// console.log(event, event.metaKey, event.altKey, event.ctrlKey, event.shiftKey, event.keyCode);
		if (!(event.keyCode == 13 && (event.metaKey || event.ctrlKey))) return;

		// TODO: consider a more native way to retrieve textarea value
		$.post('/db/exec', {cmd: $(event.target).val()}).done(function() {
			console.log('done');
		});
	},
	render: function() {
		return (
			React.createElement("div", {className: "console"}, 
				React.createElement("div", {className: "log"}), 
				React.createElement(mui.TextField, {className: "input", onKeyDown: this.handleKeyUp, hintText: "Type your command", multiLine: true})
			)
		);
	}
});

var DBList = React.createClass({displayName: "DBList",
	currentDB: function() { return ""; },
	render: function() {
		return (
			React.createElement("div", {className: "db-list"}, 
				React.createElement(mui.Toolbar, null, 
					React.createElement(mui.ToolbarGroup, null, 
						React.createElement(mui.RaisedButton, {label: "New DB", secondary: true})
					)
				)

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
