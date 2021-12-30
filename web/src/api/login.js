import  request  from  '../utils/request';
import  QS  from  'qs';
// 用户登录
export function login(data) {
  return request({
    url: `/auth/login`,
    method: 'POST',
    data
  })
}


// 注销
export function logout(data) {
  return request({
    url: `/auth/logout`,
    method: 'POST',
    data:QS.stringify(data)
  })
}

// 验证码
export function sendLoginCode(data) {
  return request({
    url: `/sendLoginCode`,
    method: 'POST',
    data
  })
}