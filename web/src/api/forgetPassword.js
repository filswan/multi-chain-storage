import request from '@/utils/request'
import QS from 'qs';

// 发送邮箱找回邮件
export function sendForgetPasswordUrl(data) {
  return request({
    url: `/user/forget_password`,
    method: 'POST',
    data
  })
}

// 忘记密码，新密码设置
export function forgetPasswordEmailUserUrl(data) {
  return request({
    url: `/user/forget_password_email_url`,
    method: 'POST',
    data
  })
}

// 邮箱找回
export function updateLoginPassword(data) {
  return request({
    url: `/user/update_login_password`,
    method: 'POST',
    data
  })
}
