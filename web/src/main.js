// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import Router from 'vue-router'
import Vuex from 'vuex'
import store from './store'
import i18n from './lang'

import ElementUI from 'element-ui'
import locale from 'element-ui/lib/locale/lang/en'
import localeCN from 'element-ui/lib/locale/lang/zh-CN'
import 'element-ui/lib/theme-chalk/index.css'
import './assets/css/icon.css'
import './assets/css/main.css'

import Meta from 'vue-meta'
import * as statusColor from './utils/status_color'
import commonFun from './utils/common'
import * as metaLogin from './utils/login'
import Web3 from 'web3'
import md5 from 'js-md5'
import SparkMD5 from 'spark-md5'
import uploader from 'vue-simple-uploader'

Vue.use(Router)
let langNew = store.getters.languageMcs === 'en' ? {
  locale
} : {
  localeCN
}
Vue.use(ElementUI, langNew)
Vue.use(Vuex)
Vue.use(Meta)
Vue.use(uploader)
Vue.prototype.$statusColor = statusColor
Vue.prototype.$commonFun = commonFun
Vue.prototype.$metaLogin = metaLogin
Vue.prototype.$web3Init = commonFun.web3Init
Vue.prototype.$md5 = md5
Vue.prototype.$SparkMD5 = SparkMD5

Vue.config.productionTip = false
let netData = Number(sessionStorage.getItem('networkID')) || 0
// netData === 97 ? process.env.BASE_PAYMENT_GATEWAY_BSC_API :
// netData === 97 ? process.env.BASE_BSC_ADDRESS :
Vue.prototype.baseAPIURL = netData === 80001 ? process.env.BASE_PAYMENT_GATEWAY_API : process.env.BASE_PAYMENT_GATEWAY_POLYGON_API
Vue.prototype.baseAddressURL = netData === 80001 ? process.env.BASE_MUMBAI_ADDRESS : process.env.BASE_POLYGON_ADDRESS
Vue.prototype.Web3 = Web3
Vue.prototype.baseNetwork = process.env.BASE_ENV === true
console.log('update time: 2022-1-9', 'env:', process.env.BASE_ENV === true ? 'Main' : 'Cali', process.env.BASE_ENV)
router.beforeEach((to, from, next) => {
  if (to.matched.some(record => record.meta.requiresAuth)) {
    if (!sessionStorage.getItem('metaAddress')) {
      next({
        path: '/home',
        query: {
          redirect: to.fullPath
        }
      })
    } else {
      next()
    }
  } else {
    if (to.meta.metaInfo) {
      // this.$store.dispatch('setMetaInfo', to.meta.metaInfo)
      // console.log(to.meta.metaInfo);
      store.commit('SET_METAINFO', to.meta.metaInfo)
    }
    next()
  }

  window.scrollTo(0, 0)
})

/* eslint-disable no-new */
new Vue({
  el: '#app',
  components: {
    App
  },
  template: '<App/>',
  router,
  store,
  i18n,
  metaInfo () {
    return {
      title: this.$store.getters.metaInfo.title,
      meta: [{
        name: 'description',
        content: this.$store.getters.metaInfo.description
      }]
    }
  },
  data: function () {
    return {
      LOCK_TIME: '6',
      PAY_GAS_LIMIT: '9999999',
      PAY_WITH_MULTIPLY_FACTOR: '1.5',
      RECIPIENT: '',
      SWAN_PAYMENT_CONTRACT_ADDRESS: '',
      USDC_ADDRESS: '',
      dao_threshold: '',
      filecoin_price: ''
    }
  }
})
