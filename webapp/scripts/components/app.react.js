var React = require('react');
var ReactRouter = require('react-router');

var Route = ReactRouter.Route;
var DefaultRoute = ReactRouter.DefaultRoute;
var RouteHandler = ReactRouter.RouteHandler;

var App = React.createClass({
  render: function() {
    return (
      <div>
        <RouteHandler/>
      </div>
    );
  }
});

var BranchList = require('./branch.react').BranchList;
var Branch = require('./branch.react').Branch;

var _routes = (
  <Route name='app' path='/' handler={App}>
    <DefaultRoute handler={BranchList} />
    <Route name='branch' path='/b/:branchName' handler={Branch} />
  </Route>
);

module.exports = { 
  render: function(el) {
    ReactRouter.run(_routes, function(Handler) {
      React.render(<Handler/>, el);
    });
  }
};
