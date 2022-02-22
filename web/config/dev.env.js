'use strict'
const merge = require('webpack-merge')
const prodEnv = require('./staging.env')

module.exports = merge(prodEnv, {
  NODE_ENV: '"development"',
  BASE_API: '"https://calibration-api.filswan.com/"',
  BASE_PAYMENT_GATEWAY_API: '"https://calibration-mcp-api.filswan.com/"'
})
