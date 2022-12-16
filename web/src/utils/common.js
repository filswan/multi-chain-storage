import store from '../store'
import axios from 'axios'

export async function sendRequest (apilink, type, jsonObject, config) {
  // axios.defaults.timeout = 60000
  axios.defaults.headers.common['Authorization'] = 'Bearer ' + store.getters.mcsjwtToken
  try {
    let response
    switch (type) {
      case 'post':
        response = await axios.post(apilink, jsonObject, config)
        return response.data
      case 'put':
        response = await axios.put(apilink, jsonObject)
        return response.data
      case 'get':
        response = await axios.get(apilink)
        return response.data
      case 'delete':
        response = await axios.delete(apilink, {
          data: jsonObject
        })
        return response.data
    }
  } catch (err) {
    console.error(err)
  }
}

export async function timeout (delay) {
  return new Promise((resolve) => setTimeout(resolve, delay))
}

const Web3 = require('web3')
const ethereum = window.ethereum
let web3
if (window.ethereum) {
  web3 = new Web3(ethereum)
  web3.setProvider(ethereum)
} else if (window.web3) {
  web3 = window.web3
  console.log('Injected web3 detected.')
} else if (typeof web3 !== 'undefined') {
  var currentProvider = web3.currentProvider
  web3 = new Web3(currentProvider)
  web3.setProvider(currentProvider)
  console.log('No web3 instance injected, using Local web3.')
}
const web3Init = web3

export async function Init (callback) {
  if (typeof window.ethereum === 'undefined') {
    if (/Mobi|Android|iPhone/i.test(navigator.userAgent)) {
      alert('Please use MetaMask\'s built-in browser if you are using a cell phone')
      return
    }
    alert('Consider installing MetaMask! ')
    window.open('https://metamask.io/download.html')
  } else {
    ethereum
      .request({
        method: 'eth_requestAccounts'
      })
      .then((accounts) => {
        if (!accounts) {
          return false
        }
        web3Init.eth.getAccounts().then(async webAccounts => {
          store.dispatch('setMetaAddress', webAccounts[0])
          const chainId = await ethereum.request({
            method: 'eth_chainId'
          })
          store.dispatch('setMetaNetworkId', parseInt(chainId, 16))
          callback(webAccounts[0])
        })
          .catch(async () => {
            store.dispatch('setMetaAddress', accounts[0])
            const chainId = await ethereum.request({
              method: 'eth_chainId'
            })
            store.dispatch('setMetaNetworkId', parseInt(chainId, 16))
            callback(accounts[0])
          })
      })
      .catch((error) => {
        if (error === 'User rejected provider access') {} else {
          alert('Please unlock MetaMask and switch to the correct network.')
          return false
        }
        console.error(
          `Error fetching accounts: ${error.message}.
        Code: ${error.code}. Data: ${error.data}`
        )
      })
  }
}

export default {
  Init,
  sendRequest,
  timeout,
  web3Init
}
