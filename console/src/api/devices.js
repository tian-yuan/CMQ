import request from '@/utils/request'

export function fetchDeviceList(params) {
  params['Action'] = 'QueryDeviceList'
  params['Version'] = '2019-04-10'
  return request({
    url: '/iothub',
    method: 'get',
    params
  })
}
