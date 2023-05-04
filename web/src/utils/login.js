import store from '../store'
import router from '../router'
// import axios from 'axios'
import common from './common'
import {
  Message
} from 'element-ui'
const ethereum = window.ethereum
let lastTime = 0
export async function login () {
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
  const email = await emailSign(token)
  console.log(email)
  return !!token
}

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
  }
}

export async function getNonce () {
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

export async function sign (nonce) {
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

export async function performSignin (sig, nonce) {
  const netId = Number(store.getters.networkID)
  // netId === 97 ? 'bsc.testnet' :
  const reqOpts = {
    public_key_address: store.getters.metaAddress,
    nonce: nonce,
    signature: sig,
    network: 'polygon.mainnet'
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
  let status
  const baseNet = process.env.BASE_ENV === true
  // switch (id) {
  //   case 80001:
  //     status = !baseNet
  //     break
  //     // case 97:
  //     //   status = !baseNet
  //     //   break
  //   case 137:
  //     status = !!baseNet
  //     break
  //   default:
  //     status = false
  //     break
  // }
  status = true
  return status
}

export async function urlBase (id) {
  let url = ''
  switch (id) {
    // case 80001:
    //   url = process.env.BASE_PAYMENT_GATEWAY_API
    //   break
    // case 97:
    //   url = process.env.BASE_PAYMENT_GATEWAY_BSC_API
    //   break
    case 137:
      url = process.env.BASE_PAYMENT_GATEWAY_POLYGON_API
      break
    default:
      url = process.env.BASE_PAYMENT_GATEWAY_POLYGON_API
      break
  }
  return url
}
