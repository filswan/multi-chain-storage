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
      var currentProvider = web3.currentProvider;
      // var Web3 = web3js.getWeb3();
      web3 = new Web3(currentProvider);
      web3.setProvider(currentProvider);
      if(!accounts){
        return false
      }
      web3.eth.getAccounts().then(webAccounts => {
        store.dispatch('setMetaAddress', webAccounts[0])
        callback(webAccounts[0])
      })
      .catch((error) => {
        store.dispatch('setMetaAddress', accounts[0])
        callback(accounts[0])
      })
    })
    .catch((error) => {
      if (error === "User rejected provider access") {
      } else {
        alert("There was an issue signing you in. Please lock and switch to the correct network.");
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