// import axios from 'axios' // axios  npm install axios
// import Qs from 'qs'
import store from '../store'
import router from '../router'
import {
  Message
} from 'element-ui'

const service = axios.create({
  baseURL: process.env.NODE_ENV === 'production' ? process.env.BASE_API : '/api',
  timeout: 15000
})
service.interceptors.request.use(function (config) {
  if (sessionStorage.oaxLoginUserId) {
    config.baseURL = process.env.BASE_API
  }
  if (config.method === 'post') {
    // config.data = Qs.stringify(config.data)
    // config.content-type = 'application/x-www-form-urlencoded'
  }
  config.headers.lang = store.getters.languageMcs
  config.headers['Authorization'] = 'Bearer ' + store.getters.mcsjwtToken
  sessionStorage.time = 70
  return config
}, function (error) {
  // Do something with request error
  console.log(error) // for debug
  return Promise.reject(error)
})
service.interceptors.response.use(response => {
  const res = response
  if (!res.data.success) {
    // -1:用户未登录;
    if (res.data.code === '-1') {
      store.dispatch('FedLogOut').then(() => {
        // location.reload() // 为了重新实例化vue-router对象 避免bug
        //   location.href = '/login'
        router.push('/login')
      })
    }
    return response.data
  } else {
    return response.data
  }
}, function (error) {
  // 失败处理
  console.log('responseError:' + error + ',' + error.response.status) // for debug
  switch (error.response.status) {
    case 401:
      store.dispatch('FedLogOut').then(() => {
        router.push('/supplierAllBack')
      })
      break
    case 403:
      error.message = store.getters.languageMcs === 'en' ? 'please try again after 10 minutes' : '请在10分钟后再试'
      break
    case 500:
      error.message = store.getters.languageMcs === 'en' ? 'Server side error' : '服务器端出错'
      break
    case 502:
      error.message = store.getters.languageMcs === 'en' ? 'Network Error' : '网络错误'
      break
    case 504:
      error.message = store.getters.languageMcs === 'en' ? 'Network Timeout' : '网络超时'
      break
    default:
      error.message = store.getters.languageMcs === 'en' ? 'Error' : '网络连接错误'
  }
  Message({
    message: error.response.message ? error.response.message : error.message,
    type: 'error',
    duration: 5 * 1000
  })
  return Promise.reject(error)
})

export default service
