import store from '../store'
import router from '../router'
import axios from 'axios'
import {
  Message
} from 'element-ui'
const ethereum = window.ethereum
let lastTime = 0
async function login () {
  if (!store.getters.metaAddress || store.getters.metaAddress === undefined) {
    const accounts = await ethereum.request({
      method: 'eth_requestAccounts'
    })
    store.dispatch('setMetaAddress', accounts[0])
  }
  const time = await throttle()
  if (!time) return false
  const [status_, nonce] = await getNonce()
  if (status_ !== 'success') {
    Message.error(status_ || 'Fail')
    signOutFun()
    return
  }

  const signature = await sign(nonce)
  if (!signature) return false
  const token = await performSignin(signature, nonce)
  return !!token
}

async function throttle () {
  // Prevent multiple signatures
  let now = new Date().valueOf()
  if (lastTime > 0 && (now - lastTime) <= 2000) return false
  lastTime = now
  return true
}

async function sendPostRequest (apilink, jsonObject) {
  try {
    const response = await axios.post(apilink, jsonObject)
    return response.data
  } catch (err) {
    console.error(err)
    signOutFun()
  }
}

async function getNonce () {
  const reqOpts = {
    'public_key_address': store.getters.metaAddress
  }
  const netId = Number(store.getters.networkID)
  const baseAPIURL = await urlBase(netId)
  const response = await sendPostRequest(`${baseAPIURL}api/v1/user/register`, reqOpts)
  if (response.status === 'success') {
    const nonce = response.data.nonce
    return ['success', nonce]
  }
  return [response.message ? response.message : '', '']
}

async function sign (nonce) {
  store.dispatch('setMCSjwtToken', '')
  const buff = Buffer.from(nonce, 'utf-8')
  let signature = null
  await ethereum.request({
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

async function performSignin (sig, nonce) {
  const netId = Number(store.getters.networkID)
  const reqOpts = {
    public_key_address: store.getters.metaAddress,
    nonce: nonce,
    signature: sig,
    network: netId === 97 ? 'bsc.testnet' : netId === 80001 ? 'polygon.mumbai' : 'polygon.mainnet'
  }
  const baseAPIURL = await urlBase(netId)
  const response = await sendPostRequest(`${baseAPIURL}api/v1/user/login_by_metamask_signature`, reqOpts)
  if (response.status === 'success') {
    const data = response.data.jwt_token
    store.dispatch('setMCSjwtToken', data)
    return data
  }
  Message.error(response.message ? response.message : 'Fail')
  signOutFun()
  return null
}

function signOutFun () {
  store.dispatch('setMetaAddress', '')
  store.dispatch('setMCSjwtToken', '')
  store.dispatch('setMetaNetworkId', 0)
  const network = {
    name: '',
    unit: '',
    center_fail: false
  }
  store.dispatch('setMetaNetworkInfo', JSON.stringify(network))
  router.push('/supplierAllBack')
}

async function netStatus (id) {
  let status
  const baseNet = process.env.BASE_ENV === true
  switch (id) {
    case 80001:
      status = !baseNet
      break
    case 97:
      status = !baseNet
      break
    case 137:
      status = !!baseNet
      break
    default:
      status = false
      break
  }
  return status
}

async function urlBase (id) {
  let url = ''
  switch (id) {
    case 80001:
      url = process.env.BASE_PAYMENT_GATEWAY_API
      break
    case 97:
      url = process.env.BASE_PAYMENT_GATEWAY_BSC_API
      break
    case 137:
      url = process.env.BASE_PAYMENT_GATEWAY_POLYGON_API
      break
    default:
      url = process.env.BASE_PAYMENT_GATEWAY_POLYGON_API
      break
  }
  return url
}

export default {
  login,
  sendPostRequest,
  getNonce,
  sign,
  performSignin,
  signOutFun,
  netStatus,
  urlBase
}
