'use strict'
const merge = require('webpack-merge')
const prodEnv = require('./prod.env')

module.exports = merge(prodEnv, {
  NODE_ENV: '"development"'
  // BASE_METASPACE: '"https://test-meta-xieyi.nbai.io/"',
  // BASE_PAYMENT_GATEWAY_API: '"https://test-mcs-xieyi.nbai.io/"'
})
