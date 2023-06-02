// import Vue from 'vue'
// import Router from 'vue-router'
Vue.use(VueRouter)
// 路由懒加载
const home = () =>
  import('@/components/Home')
// const my_files = () => import('@/views/uploadFiles/index')
const myFilesIndex = () =>
  import('@/views/uploadFiles/dashboard/index')
const myFilesFilename = () =>
  import('@/views/uploadFiles/fileName/index')
const myFilesDetail = () =>
  import('@/views/uploadFiles/detail/index')
// const upload_file = () => import('@/components/uploadFiles')
const settings = () =>
  import('@/views/settings/index')
const billing = () =>
  import('@/views/settings/billing')
const space = () =>
  import('@/views/space/index')
const spaceDetail = () =>
  import('@/views/space/detail')
const apiKey = () =>
  import('@/views/apiKey/index')

const login = () =>
  import('@/views/login/index')
// const metamaskLogin = () =>
//   import('@/views/login/metamaskLogin')
const register = () =>
  import('@/views/register/index')
const activationSuccess = () =>
  import('@/views/activationSuccess/index.vue')
const forget = () =>
  import('@/views/forgetPassword/index.vue')
const mailForget = () =>
  import('@/views/mailForget/index.vue')
const mailResetPassword = () =>
  import('@/views/mailResetPassword/index.vue')
const mailForgetSuccess = () =>
  import('@/views/mailForgetSuccess/index.vue')
const accountActivation = () =>
  import('@/views/accountActivation/index.vue')
const supplierAllBack = () =>
  import('@/components/supplierAllBack.vue')

const homeEntrance = () =>
  import('@/views/home/index.vue')

// 配置路由
export default new VueRouter({
  mode: 'history', // 后端支持可开
  base: '/',
  // mode: 'hash',
  routes: [{
    path: '/',
    redirect: '/my_buckets'
  },
    {
      path: '/',
      component: home,
      children: [{
        path: '/my_files',
        name: 'my_files',
        component: myFilesIndex,
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
            title: 'Onchain Storage',
            description: 'Multi-Chain storage (MCS) is a smart-contract-based cross-chain storage gateway that is integrated with oracle technology. It accelerates the mass adoption of decentralized storage by bridging multiple blockchain networks.'
          }
        }
      },
        {
          path: '/my_files/detail/:id',
          name: 'my_files_filename',
          component: myFilesFilename,
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
              title: 'Onchain Storage',
              description: 'Multi-Chain storage (MCS) is a smart-contract-based cross-chain storage gateway that is integrated with oracle technology. It accelerates the mass adoption of decentralized storage by bridging multiple blockchain networks.'
            }
          }
        },
        {
          path: '/my_files/detail/:id/:deal_id/:source_file_upload_id/:isFree',
          name: 'my_files_detail',
          component: myFilesDetail,
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
              title: 'Onchain Storage',
              description: 'Multi-Chain storage (MCS) is a smart-contract-based cross-chain storage gateway that is integrated with oracle technology. It accelerates the mass adoption of decentralized storage by bridging multiple blockchain networks.'
            }
          }
        },
        {
          path: '/billing',
          name: 'billing',
          component: billing,
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
              title: 'Billing',
              description: 'Multi-Chain storage (MCS) is a smart-contract-based cross-chain storage gateway that is integrated with oracle technology. It accelerates the mass adoption of decentralized storage by bridging multiple blockchain networks.'
            }
          }
        },
        {
          path: '/settings',
          name: 'settings',
          component: settings,
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
              title: 'Settings',
              description: 'Multi-Chain storage (MCS) is a smart-contract-based cross-chain storage gateway that is integrated with oracle technology. It accelerates the mass adoption of decentralized storage by bridging multiple blockchain networks.'
            }
          }
        },
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
          path: '/my_account',
          name: 'ApiKey',
          component: apiKey,
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
              title: 'My Account',
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
        },
        {
          path: '/mail_forget',
          name: 'mail_forget',
          component: mailForget,
          meta: {
            metaInfo: {
              title: 'MailForget',
              description: 'Multi-Chain storage (MCS) is a smart-contract-based cross-chain storage gateway that is integrated with oracle technology. It accelerates the mass adoption of decentralized storage by bridging multiple blockchain networks.'
            }
          }
        },
        {
          path: '/mail_reset_password',
          name: 'mail_reset_password',
          component: mailResetPassword,
          meta: {
            metaInfo: {
              title: 'MailResetPassword',
              description: 'Multi-Chain storage (MCS) is a smart-contract-based cross-chain storage gateway that is integrated with oracle technology. It accelerates the mass adoption of decentralized storage by bridging multiple blockchain networks.'
            }
          }
        },
        {
          path: '/mail_forget_success',
          name: 'mail_forget_success',
          component: mailForgetSuccess,
          meta: {
            metaInfo: {
              title: 'MailForgetSuccess',
              description: 'Multi-Chain storage (MCS) is a smart-contract-based cross-chain storage gateway that is integrated with oracle technology. It accelerates the mass adoption of decentralized storage by bridging multiple blockchain networks.'
            }
          }
        },
        {
          path: '/account_activation',
          name: 'account_activation',
          component: accountActivation,
          meta: {
            metaInfo: {
              title: 'Account Activation',
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
      path: '*',
      redirect: '/'
    }
  ]
})
const originalPush = VueRouter.prototype.push
VueRouter.prototype.push = function push (location) {
  return originalPush.call(this, location).catch(err => err)
}
