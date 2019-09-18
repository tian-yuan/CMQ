const path = require('path');
const webpack = require('webpack');

module.exports = {
    entry: {
        vendor: ['babel-polyfill', 'whatwg-fetch', 'lodash', 'vue', 'vue-i18n', 'vue-router'],
    },
    output: {
        filename: '[name].js',
        path: path.resolve(__dirname, './dll'),
        library: '[name]',
    },
    resolve: {
        alias: {
            vue$: path.resolve(__dirname, 'node_modules/vue/dist/vue.esm.js'),
            'vue-router$': path.resolve(__dirname, 'node_modules/vue-router/dist/vue-router.esm.js'),
            'vue-i18n$': path.resolve(__dirname, 'node_modules/vue-i18n/dist/vue-i18n.esm.js'),
        },
    },
    plugins: [
        new webpack.DefinePlugin({
            'process.env': {
                NODE_ENV: '"production"',
            },
        }),
        new webpack.DllPlugin({
            path: path.join(__dirname, './dll', '[name].manifest.json'),
            name: '[name]',
        }),
        new webpack.optimize.UglifyJsPlugin({
            compress: {
                warnings: false,
            },
        }),
    ],
};
