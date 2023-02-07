'use strict'
const merge = require('webpack-merge')
const prodEnv = require('./calibration.env')

module.exports = merge(prodEnv, {
  NODE_ENV: '"development"',
  BASE_METASPACE: '"http://172.20.10.5:8889/"',
  BASE_PAYMENT_GATEWAY_API: '"http://172.20.10.5:8889/"'
})
