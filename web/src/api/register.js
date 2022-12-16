import request from '@/utils/request'

// 邮箱注册
export function emailRegister (data) {
  return request({
    url: `/auth/register`,
    method: 'POST',
    data
  })
}
