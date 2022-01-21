let mix = require('laravel-mix');
var path = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const HistoryFallback = require('connect-history-api-fallback');

mix.webpackConfig({ // Hack for https://github.com/JeffreyWay/laravel-mix/issues/1717
  output: {
    publicPath: '',
  },
});

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

mix.setPublicPath(path.normalize('./public/'));

mix.webpackConfig({
  output: {
    chunkFilename: '[name].js',
  },
  plugins: [
    new HtmlWebpackPlugin({
      template: 'src/index.html',
    }),
  ],
});

mix.js('src/js/app.js', 'js').vue({version: 2});

mix.autoload({'lodash': ['_']})
    .extract();

mix.sass('src/sass/app.scss', 'css/app.css');


if (!mix.inProduction()) {
  mix.browserSync({
    proxy: false,
    port: '3000',
    server: {
      baseDir: 'public',
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
