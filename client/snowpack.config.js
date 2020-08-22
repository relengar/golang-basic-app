const rollupPluginSvelte = require('rollup-plugin-svelte');

module.exports = {
  instalOptions: {
    rollup: {
      plugins: [rollupPluginSvelte(), "@snowpack/plugin-svelte"]
    }
  },
  extends: "@snowpack/app-scripts-svelte",
  scripts: {
    
  },
  plugins: [],
  proxy: {
    "*": "http://localhost:5000/api/recipes"
  }
}
