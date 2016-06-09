var ExtractTextPlugin = require('extract-text-webpack-plugin');
module.exports = {
  module: {
    loaders: [{
      test: /\.css$/,
      loader: ExtractTextPlugin.extract('style', 'css?minimize')
    }, {
      test: /\.js$/,
      exclude: /node_modules/,
      loaders: ['babel']
    }]
  },
  extensions: ['', '.js', '.css'],
  entry: "./client/app.js",
  output: {
    path: `static/`,
    filename: "bundle.js",
    publicPath: '/static/'

  },
  devtool: 'source-map',
  plugins: [
        new ExtractTextPlugin('main.css')
    ]
};
