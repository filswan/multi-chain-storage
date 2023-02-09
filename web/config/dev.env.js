'use strict'
const merge = require('webpack-merge')
const prodEnv = require('./calibration.env')

module.exports = merge(prodEnv, {
  NODE_ENV: '"development"',
  BASE_METASPACE: '"https://calibration-fevm-api.filswan.com/"',
  BASE_PAYMENT_GATEWAY_API: '"https://calibration-fevm-api.filswan.com/"'
})
