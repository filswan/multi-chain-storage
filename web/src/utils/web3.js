//web3.js
import store from '../store'
function Init(callback){
  if (typeof window.ethereum === "undefined") {
    window.open('https://metamask.io/download.html')
    alert("Consider installing MetaMask!");
  } else {
    const ethereum = window.ethereum;
    ethereum
    .request({ method: 'eth_requestAccounts' })
    .then((accounts) => {
      // var Web3 = web3js.getWeb3();
      if (window.ethereum) {
        web3 = new Web3(ethereum);
        web3.setProvider(ethereum);
      }
      // Legacy dapp browsers...
      else if (window.web3) {
        // Use Mist/MetaMask's provider.
        web3 = window.web3;
        console.log("Injected web3 detected.");
      }
      // Fallback to localhost; use dev console port by default...
      else {
        var currentProvider = web3.currentProvider;
        // var Web3 = web3js.getWeb3();
        web3 = new Web3(currentProvider);
        web3.setProvider(currentProvider);
        console.log("No web3 instance injected, using Local web3.");
      }
      if(!accounts){
        return false
      }
      web3.eth.getAccounts().then(async webAccounts => {
        store.dispatch('setMetaAddress', webAccounts[0])
        const chainId = await ethereum.request({ method: 'eth_chainId' })
        store.dispatch('setMetaNetworkId', parseInt(chainId, 16))
        callback(webAccounts[0])
      })
      .catch(async (error) => {
        store.dispatch('setMetaAddress', accounts[0])
        const chainId = await ethereum.request({ method: 'eth_chainId' })
        store.dispatch('setMetaNetworkId', parseInt(chainId, 16))
        callback(accounts[0])
      })
    })
    .catch((error) => {
      if (error === "User rejected provider access") {
      } else {
        alert("Please unlock MetaMask and switch to the correct network.");
        return false
      }
      console.error(
        `Error fetching accounts: ${error.message}.
        Code: ${error.code}. Data: ${error.data}`
      );
    });
  }
}

export default {
  Init
}