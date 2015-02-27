var React = require('react');
var ReactRouter = require('react-router');

var Link = ReactRouter.Link;

var Branch = React.createClass({
  mixins: [ReactRouter.State],
  render: function() {
    var name = this.getParams().branchName;

    return (
      <div key={name}>
      <div>Branch: {name}</div>
      <Link to="app">Home</Link>
      </div>
    )
  }
});

var BranchList = React.createClass({
  getInitialState: function() {
    return {
      branches: ["master"]
    };
  },

  render: function() {
    var links = this.state.branches.map(function(name) {
      return <Link key={name} to="branch" params={{branchName: name}}>{name}</Link>;
    });

    return (
      <div>{links}</div>
    );
  }
});

module.exports = {
  Branch: Branch,
  BranchList: BranchList
}
