var Console = React.createClass({
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
			<div className="console">
				<div className="log"></div>
				<mui.TextField className="input" onKeyDown={this.handleKeyUp} hintText="Type your command" multiLine={true} />
			</div>
		);
	}
});

var DBList = React.createClass({
	currentDB: function() { return ""; },
	render: function() {
		return (
			<div className="db-list">
				<mui.Toolbar>
					<mui.ToolbarGroup>
						<mui.RaisedButton label="New DB" secondary={true} />
					</mui.ToolbarGroup>
				</mui.Toolbar>

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
