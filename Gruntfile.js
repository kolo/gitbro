module.exports = function(grunt) {
  // Load grunt tasks automatically.
  require('load-grunt-tasks')(grunt);

  var path = require('path');
  var extend = require('util')._extend;

  var env = extend({
    webroot: path.resolve('.'),
    dir: path.resolve('.')
  }, process.env);

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

    shell: {
      gobuild: {
        command: 'go build gitbro/server',
        opions: {
          stderr: false
        }
      },

      server: {
        command: 'gin -b gitbro',
        options: {
          execOptions: {
            cwd: 'server',
            env: env
          }
        }
      }
    },

    watch: {
    }
  });

  grunt.registerTask('serve', [
    'concurrent:server'
  ])
};
