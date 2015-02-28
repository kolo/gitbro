var EventEmitter = require('events').EventEmitter;
var assign = require('object-assign');

var AppDispatcher = require('../dispatcher');

var _branches = [];

var BranchStore = assign({}, EventEmitter.prototype, {
  emitChange: function() {
    this.emit("CHANGE");
  },

  addChangeListener: function(callback) {
    this.on("CHANGE", callback);
  },

  removeChangeListener: function(callback) {
    this.removeListener("CHANGE", callback);
  },

  getAll: function() {
    return _branches;
  }
});

BranchStore.dispatchToken = AppDispatcher.register(function(payload) {
  switch(payload.type) {
    case "GET_BRANCH_LIST":
      _branches = payload.branches;
      BranchStore.emitChange();
      break;
    default:
      // Do nothing.
  }
});

module.exports = BranchStore;
