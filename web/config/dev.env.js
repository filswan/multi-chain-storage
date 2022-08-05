'use strict'
const merge = require('webpack-merge')
const prodEnv = require('./calibration.env')

module.exports = merge(prodEnv, {
  NODE_ENV: '"development"',
  BASE_API: '"https://calibration-api.filswan.com/"',
  BASE_PAYMENT_GATEWAY_API: '"https://calibration-mcs-api.filswan.com/"',
  BASE_MAINNET_ADDRESS: '"https://filscan.io/tipset/dsn-detail"',
  BASE_CALIBRATION_ADDRESS: '"https://calibration.filscan.io/tipset/dsn-detail"'
})
