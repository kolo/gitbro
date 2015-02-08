module.exports = function(grunt) {
  // Load grunt tasks automatically.
  require('load-grunt-tasks')(grunt);

  var path = require('path');
  var extend = require('util')._extend;

  var env = extend({webroot: path.resolve('.')}, process.env);

  grunt.initConfig({
    env: env,

    concurrent: {
      server: {
        tasks: ['shell:server', 'watch'],
        options: {
          logConcurrentOutput: true
        }
      }
    },

    react: {
      components: {
        files: [{
          expand: true,
          cwd: 'webroot/scripts/components',
          src: ['{,*/}*.react.js'],
          dest: '.tmp/scripts/components',
          ext: '.js'
        }]
      }
    },

    shell: {
      gobuild: {
        command: 'go build gitbro',
        opions: {
          stderr: false
        }
      },

      server: {
        command: 'gin -b gitbro'
      }
    },

    watch: {
      react: {
        files: ['webroot/scripts/components/{,*/}*.react.js'],
        tasks: ['react']
      }
    }
  });

  grunt.registerTask('serve', [
    'react',
    'concurrent:server'
  ])
};
