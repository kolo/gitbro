var App = require('./components/app.react');

(function() {

  document.addEventListener('DOMContentLoaded', function() {
    App.render(document.getElementById('content'));
  });
})();
