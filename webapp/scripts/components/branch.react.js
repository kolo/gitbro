var React = require('react');
var ReactRouter = require('react-router');

var Link = ReactRouter.Link;

var BranchStore = require('../stores/branch_store');

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
  componentDidMount: function() {
    BranchStore.addChangeListener(this._onChange);
  },

  getInitialState: function() {
    return getStateFromStore();
  },

  render: function() {
    var branches = this.state.branches;

    var links = branches.map(function(name) {
      return <Link key={name} to="branch" params={{branchName: name}}>{name}</Link>;
    });

    return (
      <div>{links}</div>
    );
  },

  _onChange: function() {
    this.setState(getStateFromStore());
  }
});

function getStateFromStore() {
  return {
    branches: BranchStore.getAll()
  }
};

module.exports = {
  Branch: Branch,
  BranchList: BranchList
}
