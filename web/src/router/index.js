import Vue from 'vue'
import Router from 'vue-router'
Vue.use(Router)
// 路由懒加载
const home = () =>
  import('@/components/Home')
const stats = () =>
  import('@/views/stats/index')
const space = () =>
  import('@/views/space/index')
const spaceDetail = () =>
  import('@/views/space/detail')
const login = () =>
  import('@/views/login/index')
const register = () =>
  import('@/views/register/index')
const activationSuccess = () =>
  import('@/views/activationSuccess/index.vue')
const forget = () =>
  import('@/views/forgetPassword/index.vue')
const supplierAllBack = () =>
  import('@/components/supplierAllBack.vue')

const homeEntrance = () =>
  import('@/views/home/index.vue')

// 配置路由
export default new Router({
  // mode: 'history', // 后端支持可开
  mode: 'hash',
  routes: [{
    path: '/',
    redirect: '/my_buckets'
  },
  {
    path: '/',
    component: home,
    children: [
      {
        path: '/my_buckets',
        name: 'Space',
        component: space,
        beforeEnter: (to, from, next) => {
          if (!sessionStorage.getItem('metaAddress')) {
            next({
              path: '/home'
            })
          } else {
            next()
          }
        },
        meta: {
          metaInfo: {
            title: 'Bucket Storage',
            description: 'Multi-Chain storage (MCS) is a smart-contract-based cross-chain storage gateway that is integrated with oracle technology. It accelerates the mass adoption of decentralized storage by bridging multiple blockchain networks.'
          }
        }
      },
      {
        path: '/my_buckets_detail',
        name: 'Space_detail',
        component: spaceDetail,
        beforeEnter: (to, from, next) => {
          if (!sessionStorage.getItem('metaAddress')) {
            next({
              path: '/home'
            })
          } else {
            next()
          }
        },
        meta: {
          metaInfo: {
            title: 'Bucket Storage',
            description: 'Multi-Chain storage (MCS) is a smart-contract-based cross-chain storage gateway that is integrated with oracle technology. It accelerates the mass adoption of decentralized storage by bridging multiple blockchain networks.'
          }
        }
      },
      {
        path: '/supplierAllBack',
        name: 'supplierAllBack',
        component: supplierAllBack,
        meta: {
          metaInfo: {
            description: 'Multi-Chain storage (MCS) is a smart-contract-based cross-chain storage gateway that is integrated with oracle technology. It accelerates the mass adoption of decentralized storage by bridging multiple blockchain networks.'
          }
        }
      },
      {
        path: '/login',
        name: 'login',
        component: login,
        // meta: { title: 'login' },
        meta: {
          metaInfo: {
            title: 'Login',
            description: 'Multi-Chain storage (MCS) is a smart-contract-based cross-chain storage gateway that is integrated with oracle technology. It accelerates the mass adoption of decentralized storage by bridging multiple blockchain networks.'
          }
        }
      },
      {
        path: '/register',
        name: 'register',
        component: register,
        // meta: { title: 'register' },
        meta: {
          metaInfo: {
            title: 'Register',
            description: 'Multi-Chain storage (MCS) is a smart-contract-based cross-chain storage gateway that is integrated with oracle technology. It accelerates the mass adoption of decentralized storage by bridging multiple blockchain networks.'
          }
        }
      },
      {
        path: '/activate_user',
        name: 'activation_success',
        component: activationSuccess
      },
      {
        path: '/forget',
        name: 'forget',
        component: forget,
        meta: {
          metaInfo: {
            title: 'Forget',
            description: 'Multi-Chain storage (MCS) is a smart-contract-based cross-chain storage gateway that is integrated with oracle technology. It accelerates the mass adoption of decentralized storage by bridging multiple blockchain networks.'
          }
        }
      }
    ]
  },
  {
    path: '/home',
    name: 'home_entrance',
    component: homeEntrance,
    meta: {
      metaInfo: {
        title: 'Home',
        description: 'Multi-Chain storage (MCS) is a smart-contract-based cross-chain storage gateway that is integrated with oracle technology. It accelerates the mass adoption of decentralized storage by bridging multiple blockchain networks.'
      }
    }
  },
  {
    path: '/stats',
    name: 'Stats',
    component: stats,
    meta: {
      metaInfo: {
        title: 'Stats',
        description: 'Multi-Chain storage (MCS) is a smart-contract-based cross-chain storage gateway that is integrated with oracle technology. It accelerates the mass adoption of decentralized storage by bridging multiple blockchain networks.'
      }
    }
  },
  {
    path: '*',
    redirect: '/'
  }
  ]
})
const originalPush = Router.prototype.push
Router.prototype.push = function push (location) {
  return originalPush.call(this, location).catch(err => err)
}
