import request from '@/utils/request'
// import Qs from 'qs'

// 重新发送邮件
export function sendActivateLink (data) {
  return request({
    url: '/send_activate_link',
    method: 'POST',
    data
  })
}

// 检查
export function checkCode (data) {
  return request({
    url: `/activate_user?${Qs.stringify(data)}`,
    method: 'get'
  })
}

// 网易云易盾二次验证
export function secondVerify (data) {
  return request({
    url: `/thirdApi/wyy/secondVerify`,
    method: 'POST',
    data
  })
}
