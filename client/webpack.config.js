const path = require("path");
const HtmlWebpackPlugin = require("html-webpack-plugin");
const TerserPlugin = require("terser-webpack-plugin");

module.exports = {
  entry: path.resolve(__dirname, "src/index.js"),
  output: {
    path: path.resolve(__dirname, "./public/"),
    filename: "[name].[contenthash].js"
  },
  resolve: {
    extensions: [".js", ".jsx"]
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
  plugins: [
    new HtmlWebpackPlugin({
      hash: true,
      template: "./public/index.html"
    })
  ]
};
