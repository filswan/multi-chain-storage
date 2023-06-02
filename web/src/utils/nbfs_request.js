// import axios from 'axios' // 引入axios  npm install axios
// import Qs from 'qs'
import store from '../store'
// import router from '../router'
import {
  Message
} from 'element-ui'

const service = axios.create({
  baseURL: process.env.NODE_ENV === 'production' ? process.env.NBFS_API : '/',
  timeout: 15000
})

// 请求
service.interceptors.request.use(function (config) {
  if (sessionStorage.oaxLoginUserId) {
    config.baseURL = process.env.BASE_API
  }
  if (config.method === 'post') {
    // config.data = Qs.stringify(config.data)
    // config.content-type = 'application/x-www-form-urlencoded'npm
  }
  config.headers.lang = store.getters.languageMcs
  // config.headers.userId = 74
  // config.headers.accessToken = "1723113221611601ad4845f48209bfe8d2cacb54c43"
  // config.headers['Authorization']  =  store.getters.accessToken
  config.headers['Authorization'] = 'Bearer' + store.getters.mcsjwtToken
  // config.headers['Access-Control-Allow-Origin'] = '*'
  sessionStorage.time = 70
  return config
}, function (error) {
  // Do something with request error
  console.log(error) // for debug
  return Promise.reject(error)
})
// 响应
service.interceptors.response.use(response => {
  const res = response
  if (!res.data.success) {
    return response.data
  } else {
    return response.data
  }
}, function (error) {
  // 失败处理
  console.log('responseError:' + error) // for debug
  Message({
    message: 'Error',
    type: 'error',
    duration: 5 * 1000
  })
  return Promise.reject(error)
})

export default service
