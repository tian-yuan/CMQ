const path = require('path');
const webpack = require('webpack');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const AddAssetHtmlPlugin = require('add-asset-html-webpack-plugin');

module.exports = {
    version: '>=0.7.4',
    type: 'app',
    staticPath: './static',
    docs: false,
    extractCSS: true,
    sourceMap: false,
    libraryPath: './src/components',
    // forceShaking: 'proto-ui.vusion',
    webpack: {
        entry: {
            // babel-polyfill 与 whatwg-fetch 为了兼容低版本浏览器
            // 而且在这里必须添加，相当于一个 import，否则 dll 不知道要引入此包
            bundle: ['babel-polyfill', 'whatwg-fetch', './src/views/index.js'],
        },
        resolve: {
            alias: {
                vue$: path.resolve(__dirname, 'node_modules/vue/dist/vue.esm.js'),
                'vue-router$': path.resolve(__dirname, 'node_modules/vue-router/dist/vue-router.esm.js'),
                '@': path.resolve(__dirname, 'src'),
            },
        },
        plugins: [
            new HtmlWebpackPlugin({
                filename: 'index.html',
                hash: true,
                chunks: ['bundle'],
                template: './src/views/index.html',
            }),
            // 关联生成的 dll 信息文件
            new webpack.DllReferencePlugin({
                manifest: require('./dll/vendor.manifest.json'),
            }),
            // 将 vendor.js 带上 hash 并注入到 html 中
            new AddAssetHtmlPlugin({
                filepath: path.resolve(__dirname, 'dll/vendor.js'),
                hash: true,
                includeSourcemap: false,
            }),
            // 将子 chunk 的公共代码打包进父 chunk 中
            new webpack.optimize.CommonsChunkPlugin({
                children: true,
                minChunks: 3,
            }),
        ],
    },
};
