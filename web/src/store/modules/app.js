// import Cookies from 'js-cookie'

const app = {
    state: {
        sidebar: {
            opened: true
        },
        // languageMcp: 'en',
        languageMcp: localStorage.getItem('languageMcp') || 'en',
        routerMenu: localStorage.getItem('routerMenu') || 0,
        headertitle: localStorage.getItem('headertitle') || 'DASHBOARD',
        assetNow: localStorage.getItem('assetNow') || 1,
        metaInfo: {},
        avater: localStorage.getItem('oaxLoginAvater') || require('../../assets/images/user_big.png'),
        htmlGo:{},
        collapseL:  localStorage.getItem('collapseL') || false,
        metaAddress: sessionStorage.getItem('metaAddress') || '',
        networkID: sessionStorage.getItem('networkID') || 0,
        metaNetworkInfo: sessionStorage.getItem('metaNetworkInfo') || {}
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
        SET_LANGUAGE: (state, languageMcp) => {
            state.languageMcp = languageMcp
            localStorage.setItem('languageMcp', languageMcp)
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
    },
    actions: {
        toggleSideBar({commit}) {
            commit('TOGGLE_SIDEBAR')
        },
        setLanguage({commit}, languageMcp) {
            commit('SET_LANGUAGE', languageMcp)
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
        }
    }
}

export default app
