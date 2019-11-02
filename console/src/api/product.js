import request from '@/utils/request'

export function getProductInfo(params) {
  params['Action'] = 'QueryProduct'
  params['Version'] = '2019-04-10'
  return request({
    url: '/iothub',
    method: 'get',
    params
  })
}
