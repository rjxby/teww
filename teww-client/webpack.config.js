const webpack = require('webpack');
const path = require('path');

const GLOBALS = {
  'process.env.NODE_ENV': JSON.stringify('production')
};

const pluginsDev = [ 
  new webpack.NoEmitOnErrorsPlugin() 
];

const pluginsProd = [     
    new webpack.optimize.OccurrenceOrderPlugin(),
    new webpack.DefinePlugin(GLOBALS),
    new webpack.optimize.DedupePlugin(),
    new webpack.optimize.UglifyJsPlugin() 
];

const config = {
  entry: './static/js/src/index.js',
  target: 'web',
  output: {
    path: path.resolve(__dirname, './static/js/dist'),
    filename: 'bundle.js'
  },
  module: {
    rules: [
      {test: /\.(js|jsx)$/, use: 'babel-loader'},
      {test: /(\.css)$/, use: [{ loader: 'style-loader' }, { loader: 'css-loader', options: { importLoaders: 1 }}]},
      {test: /\.(png|woff|woff2|eot|ttf|svg)$/, loader: 'url-loader?limit=100000'}
    ]
  },
  plugins: JSON.stringify('production') === 'production' ? pluginsProd : pluginsDev
};

module.exports = config;
