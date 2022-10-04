import store from '../store'
import router from '../router'
import axios from 'axios'
import { Message } from 'element-ui';
const ethereum = window.ethereum;
async function login() {
    if (!store.getters.metaAddress || store.getters.metaAddress == 'undefined'){
      const accounts = await ethereum.request({ method: 'eth_requestAccounts' })
      store.dispatch('setMetaAddress', accounts[0])
    }
    const [status_, nonce] = await get_nonce()
    if (status_ != 'success') {
      Message.error(status_?status_:'Fail')
      signOutFun()
      return
    }

    const signature = await sign(nonce)
    if(!signature) return false
    const token = await perform_signin(signature, nonce)
    return token ? true : false
}

async function sendPostRequest(apilink, jsonObject) {
    try {
        const response = await axios.post(apilink, jsonObject)
        return response.data
    } catch (err) {
        console.error(err)
    }
}

async function get_nonce() {
    const reqOpts = {
        "public_key_address": store.getters.metaAddress
    }
    const baseAPIURL = store.getters.networkID == 97?process.env.BASE_PAYMENT_GATEWAY_BSC_API:process.env.BASE_PAYMENT_GATEWAY_API
    const response = await sendPostRequest(`${baseAPIURL}api/v1/user/register`, reqOpts)
    if (response.status == 'success') {
        const nonce = response.data.nonce
        return ['success', nonce]
    }
    return [response.message?response.message:'', ""]
}

async function sign(nonce) {
    store.dispatch('setMCSjwtToken', '')
    const buff = Buffer.from(nonce, "utf-8");
    let signature = null
    await ethereum.request({
        method: "personal_sign",
        params: [ buff.toString("hex"), store.getters.metaAddress],
    }).then(sig => {
        signature =  sig
    }).catch(err => {
        console.log(err)
        signOutFun()
        signature = ''
    })
    return signature
}

async function perform_signin(sig, nonce) {
    const reqOpts = {
        public_key_address: store.getters.metaAddress, 
        nonce: nonce,
        signature: sig,
        network: store.getters.networkID == 97 ? "bsc.testnet" : "polygon.mumbai"
    }
    const baseAPIURL = store.getters.networkID == 97?process.env.BASE_PAYMENT_GATEWAY_BSC_API:process.env.BASE_PAYMENT_GATEWAY_API
    const response = await sendPostRequest(`${baseAPIURL}api/v1/user/login_by_metamask_signature`, reqOpts)
    if (response.status == 'success') {
        const data = response.data.jwt_token
        store.dispatch("setMCSjwtToken", data);
        return data
    }
    Message.error(response.message ? response.message : 'Fail')
    signOutFun()
    return null
}

function signOutFun() {
    store.dispatch('setMetaAddress', '')
    store.dispatch('setMCSjwtToken', '')
    store.dispatch('setMetaNetworkId', 0)
    const network = {
        name: '',
        unit: '',
        center_fail: false
    }
    store.dispatch('setMetaNetworkInfo', JSON.stringify(network))
    router.push("/supplierAllBack");
}

export default {
    login,
    sendPostRequest,
    get_nonce,
    sign,
    perform_signin,
    signOutFun
}