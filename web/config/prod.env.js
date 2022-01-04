'use strict'
let params = process.argv[2]
let baseUrl = ''
let paymentGatewayUrl = ''
switch (params) {
    case 'dev':
      baseUrl = '"http://192.168.88.216:5002/"'
      paymentGatewayUrl = '"http://192.168.88.217:8889/"'
      break
    default:
      baseUrl = '"https://calibration-api.filswan.com/"'
      paymentGatewayUrl = '"https://calibration-mcp-api.filswan.com/"'
}

module.exports = {
  NODE_ENV: '"production"',
  BASE_API: baseUrl,
  BASE_PAYMENT_GATEWAY_API: paymentGatewayUrl
}
