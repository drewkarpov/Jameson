let mix = require('laravel-mix');
var path = require('path');
const replace = require('replace-in-file');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const HistoryFallback = require('connect-history-api-fallback');

let publicDir = '../public/';

mix.options({
    terser: {
        extractComments: (astNode, comment) => false,
        terserOptions: {
            format: {
                comments: false,
            }
        }
    }
});

mix.setPublicPath(path.normalize(publicDir));

mix.webpackConfig({
    output: {
        publicPath: './',
        filename: '[name].js',
    },
    plugins: [
        new HtmlWebpackPlugin({
            template: 'src/index.html',
        }),
    ],
});

mix.js('src/js/app.js', 'js').vue({version: 2})
    .then(() => replace.sync({ // https://github.com/laravel-mix/laravel-mix/issues/1717
        files: path.normalize(`${publicDir}/index.html`),
        from: /.\/\//gu,
        to: './',
    }));

mix.autoload({'lodash': ['_']})
    .extract();

mix.sass('src/sass/app.scss', 'css/app.css');


if (!mix.inProduction()) {
    mix.browserSync({
        proxy: false,
        port: '3000',
        server: {
            baseDir: '../public',
            middleware: [
                HistoryFallback()
            ]
        }
    });
}

// Some useful commands ↓
// mix.js('src/file.js', 'dist/file.js').extract(['vue']);
// mix.copy(from, to);
// mix.copyDirectory(fromDir, toDir);
// mix.sourceMaps(); // Enable sourcemaps
