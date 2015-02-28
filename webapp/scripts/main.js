var DataActions = require('./actions/data_actions.js');
var App = require('./components/app.react');

(function() {
  document.addEventListener('DOMContentLoaded', function() {
    DataActions.loadInitialData();
    App.render(document.getElementById('content'));
  });
})();
