import request from '@/utils/request'
import qs from 'Qs'

export function fetchDeviceList(params) {
  params['Action'] = 'QueryDeviceList'
  params['Version'] = '2019-04-10'
  return request({
    url: '/iothub',
    method: 'get',
    params
  })
}

export function registerDevices(params, body) {
  params['Action'] = 'RegisterDevices'
  params['Version'] = '2019-04-10'
  return request({
    url: '/iothub',
    method: 'post',
    params,
    data: qs.stringify(body)
  })
}

export function deleteDevice(params) {
  params['Action'] = 'DeleteDevice'
  params['Version'] = '2019-04-10'
  return request({
    url: '/iothub',
    method: 'get',
    params
  })
}
