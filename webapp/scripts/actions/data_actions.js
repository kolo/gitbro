var request = require('superagent');

var AppDispatcher = require('../dispatcher');

module.exports = {
  loadInitialData: function() {
    request
      .get('branches')
      .accept('application/json')
      .end(function(res) {
        var payload = {
          type: "GET_BRANCH_LIST",
          branches: res.body.branches
        };

        AppDispatcher.dispatch(payload);
    });
  }
};
