import request from '@/utils/request'

export const getStatsStorage = () => {
    return request({
      url: '/stats/storage',
      method:'get',
    })
  }