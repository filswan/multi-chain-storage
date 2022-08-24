import Vue from 'vue'
import Router from 'vue-router'
Vue.use(Router)
//路由懒加载
const home = () => import("@/components/Home");
const my_files = () => import("@/views/uploadFiles/index");
const my_files_index = () => import("@/views/uploadFiles/dashboard/index");
const my_files_filename = () => import("@/views/uploadFiles/fileName/index");
const my_files_detail = () => import("@/views/uploadFiles/detail/index");
const upload_file = () => import("@/components/uploadFiles");
const Search_file = () => import("@/views/searchFile/index");
const settings = () => import("@/views/settings/index");
const billing = () => import("@/views/settings/billing");
const stats = () => import("@/views/stats/index");


const login = () => import("@/views/login/index");
const metamask_login = () => import("@/views/login/metamaskLogin");
const register = () => import("@/views/register/index");
const activation_success = () => import("@/views/activationSuccess/index.vue");    
const forget = () => import("@/views/forgetPassword/index.vue");
const mail_forget = () => import("@/views/mailForget/index.vue");
const mail_reset_password = () => import("@/views/mailResetPassword/index.vue");
const mail_forget_success = () => import("@/views/mailForgetSuccess/index.vue");
const account_activation = () => import("@/views/accountActivation/index.vue");
const supplierAllBack = () => import("@/components/supplierAllBack.vue");



//配置路由
export default new Router({
	// mode: 'history', // 后端支持可开
	mode: 'hash',
	routes: [
        {
            path: '/',
            redirect: '/my_files'
        },
        {
            path: '/',
            component: home,
            // meta: { title: '自述文件' },
            children: [
                // {
                //     path: '/upload_file',
                //     name: 'upload_file',
                //     component: upload_file,
                //     meta: {
                //         metaInfo: {
                //             title: 'Upload File',
                //             description: "Swan is a marketplace for Filecoin miners, clients post/bidding deals the online."
                //         }
                //     }
                // },
                {
                    path: '/my_files',
                    name: 'my_files',
                    component: my_files_index,
                    beforeEnter: (to, from, next) => {
                        if (!sessionStorage.getItem('metaAddress')) {
                            next({
                                path: '/metamask_login',
                                query: { redirect: to.fullPath }
                            })
                        } else {
                            next()
                        }
                    },
                    meta: {
                        metaInfo: {
                            title: 'My Files',
                            description: "Swan is a marketplace for Filecoin miners, clients post/bidding deals the online."
                        }
                    }
                },
                {
                    path: '/my_files/detail/:id',
                    name: 'my_files_filename',
                    component: my_files_filename,
                    beforeEnter: (to, from, next) => {
                        if (!sessionStorage.getItem('metaAddress')) {
                            next({
                                path: '/metamask_login',
                                query: { redirect: to.fullPath }
                            })
                        } else {
                            next()
                        }
                    },
                    meta: {
                        metaInfo: {
                            title: 'My Files',
                            description: "Swan is a marketplace for Filecoin miners, clients post/bidding deals the online."
                        }
                    }
                },
                {
                    path: '/my_files/detail/:id/:deal_id/:source_file_upload_id/:isFree',
                    name: 'my_files_detail',
                    component: my_files_detail,
                    beforeEnter: (to, from, next) => {
                        if (!sessionStorage.getItem('metaAddress')) {
                            next({
                                path: '/metamask_login',
                                query: { redirect: to.fullPath }
                            })
                        } else {
                            next()
                        }
                    },
                    meta: {
                        metaInfo: {
                            title: 'My Files',
                            description: "Swan is a marketplace for Filecoin miners, clients post/bidding deals the online."
                        }
                    }
                },
                {
                    path: '/Search_file',
                    name: 'Search_file',
                    component: Search_file,
                    beforeEnter: (to, from, next) => {
                        if (!sessionStorage.getItem('metaAddress')) {
                            next({
                                path: '/metamask_login',
                                query: { redirect: to.fullPath }
                            })
                        } else {
                            next()
                        }
                    },
                    meta: {
                        metaInfo: {
                            title: 'Search File',
                            description: "Swan is a marketplace for Filecoin miners, clients post/bidding deals the online."
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
                                path: '/metamask_login',
                                query: { redirect: to.fullPath }
                            })
                        } else {
                            next()
                        }
                    },
                    meta: {
                        metaInfo: {
                            title: 'Billing',
                            description: "Swan is a marketplace for Filecoin miners, clients post/bidding deals the online."
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
                                path: '/metamask_login',
                                query: { redirect: to.fullPath }
                            })
                        } else {
                            next()
                        }
                    },
                    meta: {
                        metaInfo: {
                            title: 'Settings',
                            description: "Swan is a marketplace for Filecoin miners, clients post/bidding deals the online."
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
                            description: "Swan is a marketplace for Filecoin miners, clients post/bidding deals the online."
                        }
                    }
                },
                {
                    path: '/supplierAllBack',
                    name: 'supplierAllBack',
                    component: supplierAllBack,
                    beforeEnter: (to, from, next) => {
                        if (!sessionStorage.getItem('metaAddress')) {
                            next({
                                path: '/metamask_login',
                                query: { redirect: to.fullPath }
                            })
                        } else {
                            next()
                        }
                    },
                    meta: {
                        metaInfo: {
                            description: "Swan is a marketplace for Filecoin miners, clients post/bidding deals the online."
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
                            description: "Swan is a marketplace for Filecoin miners, clients post/bidding deals the online."
                        }
                    }
                },
                {
                    path: '/metamask_login',
                    name: 'metamask_login',
                    component: metamask_login,
                    meta: {
                        metaInfo: {
                            title: 'Login',
                            description: "Swan is a marketplace for Filecoin miners, clients post/bidding deals the online."
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
                            description: "Swan is a marketplace for Filecoin miners, clients post/bidding deals the online."
                        }
                    }
                },
                {
                    path: '/activate_user',
                    name: 'activation_success',
                    component: activation_success,
                },
                {
                    path: '/forget',
                    name: 'forget',
                    component: forget,
                    meta: {
                        metaInfo: {
                            title: 'Forget',
                            description: "Swan is a marketplace for Filecoin miners, clients post/bidding deals the online."
                        }
                    }
                },
                {
                    path: '/mail_forget',
                    name: 'mail_forget',
                    component: mail_forget,
                    meta: {
                        metaInfo: {
                            title: 'MailForget',
                            description: "Swan is a marketplace for Filecoin miners, clients post/bidding deals the online."
                        }
                    }
                },
                {
                    path: '/mail_reset_password',
                    name: 'mail_reset_password',
                    component: mail_reset_password,
                    meta: {
                        metaInfo: {
                            title: 'MailResetPassword',
                            description: "Swan is a marketplace for Filecoin miners, clients post/bidding deals the online."
                        }
                    }
                },
                {
                    path: '/mail_forget_success',
                    name: 'mail_forget_success',
                    component: mail_forget_success,
                    meta: {
                        metaInfo: {
                            title: 'MailForgetSuccess',
                            description: "Swan is a marketplace for Filecoin miners, clients post/bidding deals the online."
                        }
                    }
                },
                {
                    path: '/account_activation',
                    name: 'account_activation',
                    component: account_activation,
                    meta: {
                        metaInfo: {
                            title: 'Account Activation',
                            description: "Swan is a marketplace for Filecoin miners, clients post/bidding deals the online."
                        }
                    }
                },
            ]
        },
        {
            path: '*',
            redirect: '/'
        }
	]
})
const originalPush = Router.prototype.push
	Router.prototype.push = function push(location) {
    return originalPush.call(this, location).catch(err => err)
}
