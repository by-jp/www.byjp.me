const env = require('dotenv').config().parsed || {};

module.exports = {
  apps : [
      {
        name: "indiekit",
        script: "npm run serve serve --port 8088",
        watch: false,
        env,
      }
  ]
};
