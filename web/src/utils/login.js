import store from '../store'
import router from '../router'
// import axios from 'axios'
import common from './common'
import {
  Message
} from 'element-ui'
let lastTime = 0

export async function login () {
  console.log('login')
  if (!store.getters.metaAddress || store.getters.metaAddress === undefined) {
    console.log('meta address undefined')
    const accounts = await common.providerInit.request({
      method: 'eth_requestAccounts'
    })
    console.log('meta address', accounts)
    store.dispatch('setMetaAddress', accounts[0])
  }
  const time = await throttle()
  if (!time) return false
  const [status_, nonce] = await getNonce()
  if (status_ !== 'success') {
    Message.error(status_ || 'Fail')
    signOutFun()
  }
  console.log('sign:')
  const signature = await sign(nonce)
  if (!signature) return false
  console.log('signature:', signature)
  const token = await performSignin(signature, nonce)
  console.log('token:', token)
  // const email = await emailSign(token)
  // console.log(email)
  return !!token
}

// export async function walletBalance () {
//   const Web3 = require('web3')
//   const $web3Ethereum = new Web3(process.env.BASE_ETHEREUM_RPC)
//   const $web3Polygon = new Web3(process.env.BASE_POLYGON_RPC)
//   let balance = false
//   try {
//     const ethBalance = await $web3Ethereum.eth.getBalance(store.getters.metaAddress).then()
//     const polygonBalance = await $web3Polygon.eth.getBalance(store.getters.metaAddress).then()
//     const ethBalanceWei = $web3Ethereum.utils.fromWei(ethBalance, 'ether')
//     const polygonBalanceWei = $web3Ethereum.utils.fromWei(polygonBalance, 'ether')
//
//     balance = ethBalanceWei >= '0.01' || polygonBalanceWei >= '10'
//   } catch (err) {
//     balance = false
//   }
//   store.dispatch('setMinBalance', balance)
//   return balance
// }

export async function throttle () {
  // Prevent multiple signatures
  let now = new Date().valueOf()
  if (lastTime > 0 && (now - lastTime) <= 2000) return false
  lastTime = now
  return true
}

export async function sendPostRequest (apilink, jsonObject) {
  try {
    const response = await axios.post(apilink, jsonObject)
    return response.data
  } catch (err) {
    console.error(err)
    signOutFun()
    if (err.response) {
      return {
        status: 'error',
        message: err.response.data.message || 'Unknown error',
        statusCode: err.response.status
      }
    } else {
      return {
        status: 'error',
        message: 'Network or other error',
        statusCode: 500 // 可以选择一个适当的状态码
      }
    }
  }
}
export async function getNonce () {
  const reqOpts = {
    'address': store.getters.metaAddress
  }
  const netId = Number(store.getters.networkID)
  const baseAPIURL = await urlBase(netId)
  const response = await sendPostRequest(`${baseAPIURL}api/v2/user/wallet_register`, reqOpts)
  if (response.status === 'success') {
    const nonce = response.data
    return ['success', nonce]
  }
  return [response.message ? response.message : '', '']
}

export async function sign (nonce) {
  store.dispatch('setMCSjwtToken', '')
  const buff = Buffer.from(nonce, 'utf-8')
  let signature = null
  console.log('sign method:', store.getters.metaAddress)
  console.log('sign method:', buff.toString('hex'))
  await common.providerInit.request({
    method: 'personal_sign',
    params: [buff.toString('hex'), store.getters.metaAddress]
  }).then(sig => {
    signature = sig
  }).catch(err => {
    console.log(err)
    signOutFun()
    signature = ''
  })
  return signature
}

export async function performSignin (sig, nonce) {
  const netId = Number(store.getters.networkID)
  // netId === 97 ? 'bsc.testnet' :
  const reqOpts = {
    address: store.getters.metaAddress,
    message: nonce,
    signature: sig
  }
  const baseAPIURL = await urlBase(netId)
  const response = await sendPostRequest(`${baseAPIURL}api/v2/user/wallet_login`, reqOpts)
  if (response.message === 'invalid param value:insufficient balance') {
    console.log('insufficient balance')
    Message({
      showClose: true,
      message: store.getters.languageMcs === 'en' ? 'You need to have minimal balance of 0.01 ETH or 10 MATIC to access MCS' : '您需要拥有0.01 ETH或10 MATIC的最低余额才能访问MCS',
      type: 'error'
    })
    signOutFun()
    return null
  }
  if (response.status === 'success') {
    const data = response.data
    store.dispatch('setMCSjwtToken', data)
    return data
  }
  Message.error(response.message ? response.message : 'Fail')
  signOutFun()
  return null
}

export async function emailSign (token, type) {
  const response = await common.sendRequest(`${process.env.BASE_PAYMENT_GATEWAY_API}api/v1/user/wallet`, 'get')
  if (response && response.status === 'success') {
    const data = response.data.wallet
    const dataEmail = data.email_popup_at || ''
    const dataShow = type ? false : !dataEmail
    data.apiStatus = token && !dataEmail ? await setPopupTime() : dataShow // Control pop-up display
    store.dispatch('setMCSEmail', JSON.stringify(data))
    return data
  }
  Message.error(response ? response.message : 'Failed to get mailbox')
  store.dispatch('setMCSEmail', JSON.stringify({
    apiStatus: false
  }))
  return null
}

export async function setPopupTime () {
  const response = await common.sendRequest(`${process.env.BASE_PAYMENT_GATEWAY_API}api/v1/user/wallet/set_popup_time`, 'put')
  if (response && response.status === 'success') return true
  return false
}

export async function Disconnect () {
  const response = await common.sendRequest(`${process.env.BASE_PAYMENT_GATEWAY_API}api/v1/user/delete_email`, 'put')
  if (response && response.status === 'success') return true
  return false
}

export function signOutFun () {
  store.dispatch('setMetaAddress', '')
  store.dispatch('setMCSjwtToken', '')
  store.dispatch('setMetaNetworkId', 0)
  store.dispatch('setMCSEmail', JSON.stringify({}))
  const network = {
    name: '',
    unit: '',
    center_fail: false
  }
  store.dispatch('setMetaNetworkInfo', JSON.stringify(network))
  router.push('/supplierAllBack')
}

export async function netStatus (id) {
  return true
}

export async function urlBase (id) {
  let url = ''
  switch (id) {
    case 80001:
      url = process.env.BASE_PAYMENT_GATEWAY_API
      break
      // case 97:
      //   url = process.env.BASE_PAYMENT_GATEWAY_BSC_API
      //   break
    case 137:
      url = process.env.BASE_PAYMENT_GATEWAY_POLYGON_API
      break
    default:
      url = process.env.BASE_ENV === true ? process.env.BASE_PAYMENT_GATEWAY_POLYGON_API : process.env.BASE_PAYMENT_GATEWAY_API
      break
  }
  return url
}
