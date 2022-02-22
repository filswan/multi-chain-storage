// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import Router from 'vue-router'
import Vuex from 'vuex'
import store from './store'
import i18n from './lang' // 语言包

import ElementUI from 'element-ui';
import locale from 'element-ui/lib/locale/lang/en'
import localeCN from 'element-ui/lib/locale/lang/zh-CN'
import 'element-ui/lib/theme-chalk/index.css'; // 默认主题
// import './assets/css/theme-green/index.css'; // 浅绿色主题
import './assets/css/icon.css';
import './assets/css/main.css';

import 'element-ui/lib/theme-chalk/index.css';

Vue.use(Router);
let langNew = store.getters.languageMcs == 'en'?{ locale }:{ localeCN }
Vue.use(ElementUI, langNew)
// Vue.use(ElementUI);
Vue.use(Vuex)
import Meta from 'vue-meta';
// 使用vue-meta
Vue.use(Meta);
import * as status_color from './utils/status_color'
Vue.prototype.$status_color = status_color;

Vue.config.productionTip = false
Vue.prototype.data_api=process.env.NODE_ENV === 'production' ? process.env.BASE_API : '/api/';

import './utils/web3.min.js';
import Web3 from 'web3'
Vue.prototype.Web3 = Web3

router.beforeEach((to, from, next) => {
    if (to.matched.some(record => record.meta.requiresAuth)) {
        if (!sessionStorage.getItem('metaAddress')) {
            next({
                path: '/metamask_login',
                query: { redirect: to.fullPath }
            })
        } else {
            next()
        }
    } else {
      if (to.meta.metaInfo){
        // this.$store.dispatch('setMetaInfo', to.meta.metaInfo) 
        // console.log(to.meta.metaInfo);
        store.commit('SET_METAINFO', to.meta.metaInfo)
      }
      next() // 确保一定要调用 next()
    }
    
    window.scrollTo(0,0);
  })


/* eslint-disable no-new */
new Vue({
  el: '#app',
  components: { App },
  template: '<App/>',
  router,
  store,
  i18n,  
  metaInfo(){
    return {
        title: this.$store.getters.metaInfo.title,
        meta: [
            {
                name: "description",
                content: this.$store.getters.metaInfo.description
            }
        ]
    }
  },
  data: function(){
      return {
        LOCK_TIME: "6",
        PAY_GAS_LIMIT: "9999999",
        PAY_WITH_MULTIPLY_FACTOR: "1.5",
        RECIPIENT: "",
        SWAN_PAYMENT_CONTRACT_ADDRESS: "",
        USDC_ADDRESS: ""
      }
  }
})
