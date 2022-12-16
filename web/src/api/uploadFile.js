import request from '@/utils/request'
import QS from 'qs'

// task list
export const getTasksList = (query) => {
  return request({
    url: '/tasks?' + QS.stringify(query),
    method: 'get',
    data: query
  })
}

export const getPaymentDeals = (query) => {
  return request({
    url: '/paymentgateway/deals?' + QS.stringify(query),
    method: 'get',
    data: query
  })
}

// task details
export const getTasksListDetails = (data) => {
  return request({
    url: `/tasks/${data}`,
    method: 'get'
  })
}

// 上传文件
export function uploadFile (data) {
  return request({
    url: `/ipfs/public/upload`,
    method: 'POST',
    data
  })
}

// search file
export const getTasksSearch = (data) => {
  return request({
    url: `/tasks/public/status/${data}`,
    method: 'get'
  })
}
