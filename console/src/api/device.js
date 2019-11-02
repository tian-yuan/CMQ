import request from '@/utils/request'

export function getDeviceInfo(params) {
  params['Action'] = 'QueryDeviceInfo'
  params['Version'] = '2019-04-10'
  return request({
    url: '/npns',
    method: 'get',
    params
  })
}

export function getPushListByDeviceId(params) {
  params['Action'] = 'QueryPushListByDeviceId'
  params['Version'] = '2019-04-10'
  return request({
    url: '/npns',
    method: 'get',
    params
  })
}
