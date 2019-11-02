import request from '@/utils/request'

export function fetchMessages(params) {
  params['Action'] = 'QueryPushMessageList'
  params['Version'] = '2019-04-10'
  return request({
    url: '/npns',
    method: 'get',
    params
  })
}
