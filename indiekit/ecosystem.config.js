const env = require('dotenv').config().parsed || {};

module.exports = {
  apps : [
      {
        name: "indiekit",
        script: "npx indiekit serve --port 8088",
        watch: false,
        env,
      }
  ]
};
