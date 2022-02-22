import { login, logout } from '@/api/login'
import { Message } from 'element-ui'

const user = {
  state: {
    enable: sessionStorage.getItem('login') || 0,
    name: sessionStorage.oaxLoginName || '',
    showName: sessionStorage.oaxShowName || '',
    email: localStorage.getItem('mcsLoginEmail') || '',
    phone: sessionStorage.oaxLoginPhone || '',
    userId: sessionStorage.oaxLoginUserId || '',
    accessToken: localStorage.getItem('mcsLoginAccessToken') || '',
    checkStatus: sessionStorage.oaxLoginCheckStatus || '',
    level: sessionStorage.oaxLoginLevel || '',
    registerType: sessionStorage.oaxLoginRegisterTypee || '',
    source: sessionStorage.oaxLoginSource || '',
    menus: [],
    permissions: [],
    linkPageName: '',
    linkTime: sessionStorage.linkTime
  },

  mutations: {
    SET_ENABLE: (state, enable) => {
      state.enable = enable
    },
    SET_NAME: (state, name) => {
      state.name = name
    },
    SET_SHOWNAME: (state, showName) => {
      state.showName = showName
    },
    SET_EMAIL: (state, email) => {
      state.email = email
    },
    SET_PHONE: (state, phone) => {
      state.phone = phone
    },
    SET_USERID: (state, userId) => {
      state.userId = userId
    },
    SET_ACCESSTOKEN: (state, accessToken) => {
      state.accessToken = accessToken
    },
    SET_MENUS: (state, menus) => {
      state.menus = menus
    },
    SET_PERMISSIONS: (state, permissions) => {
      state.permissions = permissions
    },
    SET_LINKTIME: (state, linkTime) => {
      state.linkTime = linkTime
    },
    SET_CHECKSTATUS: (state, checkStatus) => {
      state.checkStatus = checkStatus
    },
    SET_LEVEL: (state, level) => {
      state.level = level
    },
    SET_REGISTERTYPE: (state, registerType) => {
      state.registerType = registerType
    },
    SET_SOURCE: (state, source) => {
      state.source = source
    }
  },

  actions: {
    // 登录
    Login({ commit }, userInfo) {
      var _this = this
      return new Promise((resolve, reject) => {
        login(userInfo)
          .then(response => {
            if (response.code === '10020') {
              Message({
                message: response.msg,
                type: 'error',
                duration: 5 * 1000
              })
              sessionStorage.oaxRegisterMail = userInfo.username
              _this.$router.push("/asset_management")
              return false
            }
            _this.loginLoad = false
            if (response.success === true) {
              sessionStorage.oaxLoginUserId = response.data.userId
              sessionStorage.mcsLoginAccessToken = response.data.accessToken
              sessionStorage.oaxLoginName = response.data.name
              sessionStorage.oaxShowName = response.data.showName
              sessionStorage.mcsLoginEmail = response.data.email
              sessionStorage.oaxLoginPhone = response.data.phone
              sessionStorage.oaxLoginCheckStatus = response.data.checkStatus
              sessionStorage.oaxLoginLevel = response.data.level
              sessionStorage.oaxLoginRegisterTypee = response.data.registerType
              sessionStorage.oaxLoginSource = response.data.source
              const data = response.data
              commit('SET_NAME', data.name)
              commit('SET_SHOWNAME', data.showName)
              commit('SET_EMAIL', data.email)
              commit('SET_PHONE', data.phone)
              commit('SET_USERID', data.userId)
              commit('SET_ACCESSTOKEN', data.accessToken)
              commit('SET_CHECKSTATUS', data.checkStatus)
              commit('SET_LEVEL', data.level)
              commit('SET_REGISTERTYPE', data.registerType)
              commit('SET_SOURCE', data.source)
              newFunction(data)
              resolve()
            } else {
              Message({
                message: response.msg,
                type: 'error',
                duration: 5 * 1000
              })
            }
          })
          .catch(error => {
            _this.loginLoad = false
            Message({
              message: '登录失败',
              type: 'error',
              duration: 5 * 1000
            })
            console.log(error)
            reject(error)
          })
      })
    },

    // 获取用户信息
    SetTime({ commit }, time) {
      return new Promise((resolve, reject) => {
        commit('SET_LINKTIME', time)
      })
    },
    // 登出
    LogOut({ commit, state }) {
      // var _this = this
      return new Promise((resolve, reject) => {
        logout(state.enable)
          .then(() => {
            sessionStorage.removeItem('oaxLoginUserId')
            sessionStorage.removeItem('mcsLoginAccessToken')
            sessionStorage.removeItem('oaxLoginName')
            sessionStorage.removeItem('oaxShowName')
            sessionStorage.removeItem('mcsLoginEmail')
            sessionStorage.removeItem('oaxLoginPhone')
            sessionStorage.removeItem('oaxLoginCheckStatus')
            sessionStorage.removeItem('oaxLoginLevel')
            sessionStorage.removeItem('oaxLoginRegisterTypee')
            sessionStorage.removeItem('oaxLoginSource')
            sessionStorage.removeItem('oaxLoginpassword')
            commit('SET_USERID', '')
            commit('SET_SHOWNAME', '')
            commit('SET_EMAIL', '')
            commit('SET_PHONE', '')
            commit('SET_CHECKSTATUS', '')
            commit('SET_LEVEL', '')
            commit('SET_REGISTERTYPE', '')
            commit('SET_SOURCE', '')
            commit('SET_ACCESSTOKEN', '')
            commit('SET_PERMISSIONS', [])
            sessionStorage.removeItem('login')
            commit('SET_MENUS', [])
            resolve()
          })
          .catch(error => {
            reject(error)
          })
      })
    },

    // 前端 登出
    FedLogOut({ commit }) {
      // var _this = this
      return new Promise(resolve => {
        sessionStorage.removeItem('oaxLoginUserId')
        localStorage.removeItem('mcsLoginAccessToken')
        sessionStorage.removeItem('oaxLoginName')
        sessionStorage.removeItem('oaxShowName')
        localStorage.removeItem('mcsLoginEmail')
        localStorage.removeItem('oaxLoginAvater')
        sessionStorage.removeItem('oaxLoginPhone')
        sessionStorage.removeItem('oaxLoginCheckStatus')
        sessionStorage.removeItem('oaxLoginLevel')
        sessionStorage.removeItem('oaxLoginRegisterTypee')
        sessionStorage.removeItem('oaxLoginSource')
        commit('SET_USERID', '')
        commit('SET_SHOWNAME', '')
        commit('SET_EMAIL', '')
        commit('SET_PHONE', '')
        commit('SET_CHECKSTATUS', '')
        commit('SET_LEVEL', '')
        commit('SET_REGISTERTYPE', '')
        commit('SET_SOURCE', '')
        commit('SET_ACCESSTOKEN', '')
        commit('SET_PERMISSIONS', [])
        sessionStorage.removeItem('login')
        commit('SET_MENUS', [])
        resolve()
      })
    }
  }
}

export default user
function newFunction(data) {
  console.log(data.name)
}

