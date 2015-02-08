var app = app || {};

(function(){
  var GitbroApp = React.createClass({
    render: function() {
      return <h1>Gitbro</h1>;
    }
  });

  app.render = function(el) {
    React.render(<GitbroApp />, el);
  };
})();
