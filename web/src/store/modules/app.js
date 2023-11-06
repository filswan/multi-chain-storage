// import Cookies from 'js-cookie'

const app = {
  state: {
    sidebar: {
      opened: true
    },
    languageMcs: localStorage.getItem('languageMcs') || 'en',
    routerMenu: localStorage.getItem('routerMenu') || 0,
    headertitle: localStorage.getItem('headertitle') || 'DASHBOARD',
    assetNow: localStorage.getItem('assetNow') || 1,
    metaInfo: {},
    avater: localStorage.getItem('oaxLoginAvater') || require('../../assets/images/user_big.png'),
    htmlGo: {},
    // collapseL:  localStorage.getItem('collapseL') || false,
    collapseL: false,
    metaAddress: sessionStorage.getItem('metaAddress') || '',
    networkID: sessionStorage.getItem('networkID') || 0,
    metaNetworkInfo: sessionStorage.getItem('metaNetworkInfo') || {},
    reverse: localStorage.getItem('reverse') || 0,
    free_usage: sessionStorage.getItem('free_usage') || 0,
    free_quota_per_month: sessionStorage.getItem('free_quota_per_month') || 0,
    free_bucket: sessionStorage.getItem('free_bucket') || 0,
    free_bucketAll: sessionStorage.getItem('free_bucketAll') || 0,
    mcsjwtToken: sessionStorage.getItem('mcs_dev_jwtToken') || '',
    mcsEmail: sessionStorage.getItem('mcs_bucket_email') || JSON.stringify({}),
    minBalance: sessionStorage.getItem('min_balance') || false
  },
  mutations: {
    TOGGLE_SIDEBAR: state => {
      if (state.sidebar.opened) {
        localStorage.setItem('sidebarStatus', 1)
      } else {
        localStorage.setItem('sidebarStatus', 0)
      }
      state.sidebar.opened = !state.sidebar.opened
    },
    SET_LANGUAGE: (state, languageMcs) => {
      state.languageMcs = languageMcs
      localStorage.setItem('languageMcs', languageMcs)
    },
    SET_ROUTERMENU: (state, routerMenu) => {
      // console.log(state, routerMenu);
      state.routerMenu = routerMenu
      localStorage.setItem('routerMenu', routerMenu)
    },
    SET_HEADERTITLE: (state, headertitle) => {
      // console.log(state, routerMenu);
      state.headertitle = headertitle
      localStorage.setItem('headertitle', headertitle)
    },
    SET_ASSETNOW: (state, assetNow) => {
      state.assetNow = assetNow
      localStorage.setItem('assetNow', assetNow)
    },
    SET_METAINFO: (state, metaInfo) => {
      state.metaInfo = metaInfo
      localStorage.setItem('metaInfo', metaInfo)
    },
    SET_AVATER: (state, avater) => {
      state.avater = avater
    },
    IS_HTML_GO: (state, avater) => {
      state.htmlGo = avater
    },
    SET_COLLAPSE: (state, collapseL) => {
      state.collapseL = collapseL
      localStorage.setItem('collapseL', collapseL)
    },
    SET_METAADDRESS: (state, metaAddress) => {
      state.metaAddress = metaAddress
      sessionStorage.setItem('metaAddress', metaAddress)
    },
    SET_METANETWORKID: (state, networkID) => {
      state.networkID = networkID
      sessionStorage.setItem('networkID', Number(networkID))
    },
    SET_METANETWORKINFO: (state, metaNetworkInfo) => {
      state.metaNetworkInfo = metaNetworkInfo
      sessionStorage.setItem('metaNetworkInfo', metaNetworkInfo)
    },
    SET_REVERSE: (state, reverse) => {
      state.reverse = reverse
      localStorage.setItem('reverse', reverse)
    },
    SET_FREEUSAGE: (state, freeUsage) => {
      state.free_usage = freeUsage
      sessionStorage.setItem('free_usage', freeUsage)
    },
    SET_FREEQUOTE: (state, freeQuotaPerMonth) => {
      state.free_quota_per_month = freeQuotaPerMonth
      sessionStorage.setItem('free_quota_per_month', freeQuotaPerMonth)
    },
    SET_FREEBUCKET: (state, freeBucket) => {
      state.free_bucket = freeBucket
      sessionStorage.setItem('free_bucket', freeBucket)
    },
    SET_FREEBUCKETALL: (state, freeBucketAll) => {
      state.free_bucketAll = freeBucketAll
      sessionStorage.setItem('free_bucketAll', freeBucketAll)
    },
    SET_MCSJWTTOKEN: (state, mcsjwtToken) => {
      state.mcsjwtToken = mcsjwtToken
      sessionStorage.setItem('mcs_dev_jwtToken', mcsjwtToken)
    },
    SET_MCSEMAIL: (state, mcsEmail) => {
      state.mcsEmail = mcsEmail
      sessionStorage.setItem('mcs_bucket_email', mcsEmail)
    },
    SET_MINBALANCE: (state, minBalance) => {
      state.minBalance = minBalance
      sessionStorage.setItem('min_balance', minBalance)
    }
  },
  actions: {
    toggleSideBar ({
      commit
    }) {
      commit('TOGGLE_SIDEBAR')
    },
    setLanguage ({
      commit
    }, languageMcs) {
      commit('SET_LANGUAGE', languageMcs)
    },
    setRouterMenu ({
      commit
    }, routerMenu) {
      commit('SET_ROUTERMENU', routerMenu)
    },
    setHeadertitle ({
      commit
    }, headertitle) {
      commit('SET_HEADERTITLE', headertitle)
    },
    setAssetNow ({
      commit
    }, assetNow) {
      commit('SET_ASSETNOW', assetNow)
    },
    setMetaInfo ({
      commit
    }, metaInfo) {
      commit('SET_METAINFO', metaInfo)
    },
    setAvater ({
      commit
    }, avater) {
      commit('SET_AVATER', avater)
    },
    isHtmlGo ({
      commit
    }, avater) {
      commit('IS_HTML_GO', avater)
    },
    setCollapse ({
      commit
    }, collapseL) {
      commit('SET_COLLAPSE', collapseL)
    },
    setMetaAddress ({
      commit
    }, metaAddress) {
      commit('SET_METAADDRESS', metaAddress)
    },
    setMetaNetworkId ({
      commit
    }, networkID) {
      commit('SET_METANETWORKID', networkID)
    },
    setMetaNetworkInfo ({
      commit
    }, metaNetworkInfo) {
      commit('SET_METANETWORKINFO', metaNetworkInfo)
    },
    setReverse ({
      commit
    }, reverse) {
      commit('SET_REVERSE', reverse)
    },
    setFreeUsage ({
      commit
    }, freeUsage) {
      commit('SET_FREEUSAGE', freeUsage)
    },
    setFreeQuote ({
      commit
    }, freeQuotaPerMonth) {
      commit('SET_FREEQUOTE', freeQuotaPerMonth)
    },
    setFreeBucket ({
      commit
    }, freeBucket) {
      commit('SET_FREEBUCKET', freeBucket)
    },
    setFreeBucketAll ({
      commit
    }, freeBucketAll) {
      commit('SET_FREEBUCKETALL', freeBucketAll)
    },
    setMCSjwtToken ({
      commit
    }, mcsjwtToken) {
      commit('SET_MCSJWTTOKEN', mcsjwtToken)
    },
    setMCSEmail ({
      commit
    }, mcsEmail) {
      commit('SET_MCSEMAIL', mcsEmail)
    },
    setMinBalance ({
      commit
    }, minBalance) {
      commit('SET_MINBALANCE', minBalance)
    }
  }
}

export default app
