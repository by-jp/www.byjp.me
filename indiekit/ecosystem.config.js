const env = require('dotenv').config().parsed || {};

module.exports = {
  apps : [
      {
        name: "indiekit",
        script: "npm run start",
        watch: false,
        env,
      }
  ]
};
