const path = require('path');
const webpack = require('webpack');
const UglifyJsPlugin = require('uglifyjs-webpack-plugin');
const UnminifiedWebpackPlugin = require('unminified-webpack-plugin');
const nodeExternals = require('webpack-node-externals');

const DEV = !!process.env.dev;
const serverConfig = {
    target: 'node',
    entry: './index.js',
    output: {
        path: path.resolve(__dirname, 'dist'),
        filename: DEV ? `neyun.node.dev.min.js` : `neyun.node.min.js`,
        library: 'NEYUN',
        libraryTarget: 'umd',
    },
    plugins: [
        new webpack.DefinePlugin({
            DEV,
            SERVER: true,
        }),
        new UglifyJsPlugin(),
        new UnminifiedWebpackPlugin(),
    ],
    module: {
        rules: [
            {
                test: /\.js$/,
                exclude: /(node_modules|bower_components)/,
                use: {
                    loader: 'babel-loader',
                    options: {
                        presets: ['babel-preset-env'],
                    },
                },
            },
        ],
    },
    externals: [nodeExternals()],
};

const clientConfig = {
    target: 'web',
    entry: './index.js',
    output: {
        path: path.resolve(__dirname, 'dist'),
        filename: DEV ? `neyun.dev.min.js` : `neyun.min.js`,
        library: 'NEYUN',
        libraryTarget: 'umd',
    },
    module: {
        rules: [
            {
                test: /\.js$/,
                exclude: /(node_modules|bower_components)/,
                use: [
                    {
                        loader: 'babel-loader',
                        options: {
                            presets: ['babel-preset-env'],
                        },
                    },
                ],
            },
        ],
    },
    plugins: [
        new webpack.DefinePlugin({
            DEV,
            SERVER: false,
        }),
        new UglifyJsPlugin(),
        new UnminifiedWebpackPlugin(),
    ],
};

module.exports = [serverConfig, clientConfig];
