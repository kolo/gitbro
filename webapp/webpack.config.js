module.exports = {
  context: __dirname,
  entry: './scripts/main.js',
  output: {
    path: __dirname + '/scripts',
    filename: 'webapp.bundle.js'
  },
  module: {
    loaders: [
      { test: /\.react\.js$/, loaders: ['jsx?harmony'] }
    ]
  }
};
