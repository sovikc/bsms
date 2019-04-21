const path = require("path");
const webpack = require("webpack");
const HtmlWebpackPlugin = require("html-webpack-plugin");
const TerserPlugin = require("terser-webpack-plugin");

module.exports = {
  entry: path.resolve(__dirname, "src/index.js"),
  output: {},
  resolve: {
    extensions: [".js", ".jsx"]
  },
  devServer: {
    writeToDisk: true,
    proxy: {
      "/": "http://localhost:8000/"
    }
  },
  module: {
    rules: [
      {
        test: /\.(js)$/,
        exclude: /node_modules/,
        use: [
          {
            loader: "babel-loader",
            options: {
              cacheDirectory: true,
              babelrc: false,
              presets: [
                [
                  "@babel/env",
                  {
                    targets: {
                      browsers: ["last 2 versions"]
                    },
                    modules: false,
                    loose: true
                  }
                ],
                "@babel/react"
              ],
              plugins: ["@babel/plugin-proposal-class-properties"]
            }
          }
        ]
      },
      {
        test: /\.html$/,
        use: [
          {
            loader: "html-loader",
            options: {
              interpolate: true,
              minimize: true,
              removeComments: true,
              collapseWhitespace: true
            }
          }
        ]
      },
      {
        test: /\.js$/,
        use: ["source-map-loader"],
        enforce: "pre"
      }
    ]
  },
  resolve: {
    modules: ["src", "node_modules"]
  },
  optimization: {
    minimizer: [
      new TerserPlugin({
        chunkFilter: chunk => {
          if (chunk.name === "vendor") {
            return false;
          }

          return true;
        },
        terserOptions: {
          ecma: 5,
          sourceMap: true,
          toplevel: true,
          warnings: false,
          parse: {},
          compress: {},
          mangle: {
            module: true,
            eval: true,
            toplevel: true
          },
          module: true,
          output: null,
          nameCache: null,
          ie8: false,
          keep_classnames: false,
          keep_fnames: false,
          safari10: false
        }
      })
    ],
    splitChunks: {
      chunks: "all",
      cacheGroups: {
        commons: {
          test: /[\\/]node_modules[\\/]/,
          name: "vendor",
          chunks: "all"
        }
      }
    }
  },
  devtool: false,
  plugins: [
    new HtmlWebpackPlugin({
      hash: true,
      template: "./public/index.html"
    }),
    new webpack.SourceMapDevToolPlugin({
      filename: "[name].js.map",
      exclude: ["vendor.js"]
    })
  ]
};