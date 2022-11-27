import request from '@/utils/space_request'
import QS from 'qs';

export const getBucketsList = (query) => {
    return request({
        url: '/api/v3/directory',
        method: 'get'
    })
}
