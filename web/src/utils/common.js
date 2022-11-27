import store from '../store'
import axios from 'axios'

async function sendRequest (apilink, type, jsonObject, config) {
  // axios.defaults.timeout = 60000
  axios.defaults.headers.common['Authorization'] = 'Bearer ' + store.getters.mcsjwtToken
  try {
    let response
    switch (type) {
      case 'post':
        response = await axios.post(apilink, jsonObject, config)
        return response.data
      case 'put':
        response = await axios.put(apilink, jsonObject)
        return response.data
      case 'get':
        response = await axios.get(apilink)
        return response.data
      case 'delete':
        response = await axios.delete(apilink, {
          data: jsonObject
        })
        return response.data
    }
  } catch (err) {
    console.error(err)
  }
}

async function timeout (delay) {
  return new Promise((resolve) => setTimeout(resolve, delay))
}

export default {
  sendRequest,
  timeout
}
