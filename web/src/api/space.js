import request from '@/utils/space_request'

export const getBucketsList = (query) => {
  return request({
    url: '/api/v3/directory',
    method: 'get'
  })
}
