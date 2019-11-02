import request from '@/utils/request'

export function getMessageInfo(params) {
  params['Action'] = 'QueryMessageInfo'
  params['Version'] = '2019-04-10'
  return request({
    url: '/npns',
    method: 'get',
    params
  })
}
