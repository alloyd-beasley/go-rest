const path = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const HtmlWebpackRootPlugin = require('html-webpack-root-plugin');

module.exports = {
  // change to .tsx if necessary
  entry: './src/index.tsx',
  output: {
    filename: 'bundle.js',
    path: path.resolve(__dirname, 'dist'),
  },

  module: {
    rules: [
      {
        test: /\.(t|j)sx?$/,
        use: ['ts-loader'],
        exclude: '/node_modules/'
      },
      {
        test: /\.css$/i,
        use: ['style-loader', 'css-loader'],
      },
    ]
  },
  devServer: {
    contentBase: "./dist",
  },
  devtool: "source-map",
  plugins: [new HtmlWebpackPlugin(), new HtmlWebpackRootPlugin('root')],
  resolve: {
    extensions: [".ts", ".tsx", ".js", ".jsx"]
  },
}