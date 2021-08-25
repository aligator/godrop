/* eslint-disable @typescript-eslint/no-var-requires */
/* eslint-disable no-undef */

const { merge } = require('webpack-merge');
const common = require('./webpack.common.js');

module.exports = merge(common, {
  mode: 'development',
  devtool: 'source-map',
  devServer: {
    contentBase: './ui/dist'
  }
});