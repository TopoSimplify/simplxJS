var compressor = require('node-minify');

// Using Google Closure Compiler
compressor.minify({
  compressor: 'gcc',
  input: 'gojs.js',
  output: 'gojs-min.js',
  callback: function (err, min) {}
});