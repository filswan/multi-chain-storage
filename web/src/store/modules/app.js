// import Cookies from 'js-cookie'

const app = {
    state: {
        sidebar: {
            opened: true
        },
        // languageMcs: 'en',
        languageMcs: localStorage.getItem('languageMcs') || 'en',
        routerMenu: localStorage.getItem('routerMenu') || 0,
        headertitle: localStorage.getItem('headertitle') || 'DASHBOARD',
        assetNow: localStorage.getItem('assetNow') || 1,
        metaInfo: {},
        avater: localStorage.getItem('oaxLoginAvater') || require('../../assets/images/user_big.png'),
        htmlGo:{},
        // collapseL:  localStorage.getItem('collapseL') || false,
        collapseL:  false,
        metaAddress: sessionStorage.getItem('metaAddress') || '',
        networkID: sessionStorage.getItem('networkID') || 0,
        metaNetworkInfo: sessionStorage.getItem('metaNetworkInfo') || {},
        reverse: localStorage.getItem('reverse') || 0,
        free_usage: sessionStorage.getItem('free_usage') || 0,
        free_quota_per_month: sessionStorage.getItem('free_quota_per_month') || 0,
        mcsjwtToken: sessionStorage.getItem('mcs_jwtToken') || ''
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
            sessionStorage.setItem('networkID', networkID)
        },
        SET_METANETWORKINFO: (state, metaNetworkInfo) => {
            state.metaNetworkInfo = metaNetworkInfo
            sessionStorage.setItem('metaNetworkInfo', metaNetworkInfo)
        },
        SET_REVERSE: (state, reverse) => {
            state.reverse = reverse
            localStorage.setItem('reverse', reverse)
        },
        SET_FREEUSAGE: (state, free_usage) => {
            state.free_usage = free_usage
            sessionStorage.setItem('free_usage', free_usage)
        },
        SET_FREEQUOTE: (state, free_quota_per_month) => {
            state.free_quota_per_month = free_quota_per_month
            sessionStorage.setItem('free_quota_per_month', free_quota_per_month)
        },
        SET_MCSJWTTOKEN: (state, mcsjwtToken) => {
            state.mcsjwtToken = mcsjwtToken
            sessionStorage.setItem('mcs_jwtToken', mcsjwtToken)
        }
    },
    actions: {
        toggleSideBar({commit}) {
            commit('TOGGLE_SIDEBAR')
        },
        setLanguage({commit}, languageMcs) {
            commit('SET_LANGUAGE', languageMcs)
        },
        setRouterMenu({commit}, routerMenu) {
            commit('SET_ROUTERMENU', routerMenu)
        },
        setHeadertitle({commit}, headertitle) {
            commit('SET_HEADERTITLE', headertitle)
        },
        setAssetNow({commit}, assetNow) {
            commit('SET_ASSETNOW', assetNow)
        },
        setMetaInfo({commit}, metaInfo) {
            commit('SET_METAINFO', metaInfo)
        },
        setAvater({commit}, avater) {
            commit('SET_AVATER', avater)
        },
        isHtmlGo({commit}, avater) {
            commit('IS_HTML_GO', avater)
        },
        setCollapse({commit}, collapseL) {
            commit('SET_COLLAPSE', collapseL)
        },
        setMetaAddress({commit}, metaAddress) {
            commit('SET_METAADDRESS', metaAddress)
        },
        setMetaNetworkId({commit}, networkID) {
            commit('SET_METANETWORKID', networkID)
        },
        setMetaNetworkInfo({commit}, metaNetworkInfo) {
            commit('SET_METANETWORKINFO', metaNetworkInfo)
        },
        setReverse({commit}, reverse) {
            commit('SET_REVERSE', reverse)
        },
        setFreeUsage({commit}, free_usage) {
            commit('SET_FREEUSAGE', free_usage)
        },
        setFreeQuote({commit}, free_quota_per_month) {
            commit('SET_FREEQUOTE', free_quota_per_month)
        },
        setMCSjwtToken({commit}, mcsjwtToken) {
            commit('SET_MCSJWTTOKEN', mcsjwtToken)
        }
    }
}

export default app
