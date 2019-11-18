// module.exports = {
//   publicPath: '/statics'
// }

const { resolve } = require("path")
const CopyWebpackPlugin = require('copy-webpack-plugin');

const env = process.env.NODE_ENV;

let config = {
  configureWebpack: {
    resolve: {
      alias: {
        "@/": resolve(__dirname, 'src/')
      }
    },
    plugins: [
      new CopyWebpackPlugin([
        { 
          from: resolve("./src", "statics"),
          to: 'statics',
          ignore: [
            ".gitkeep"
          ]
        }
      ])  
    ]

  }
}

if (env && env.toLowerCase() !== "development") {
  config.publicPath = "/statics/client";
}

module.exports = config;