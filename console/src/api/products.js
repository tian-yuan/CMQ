import request from '@/utils/request'

export function fetchProductList(params) {
  params['Action'] = 'QueryProductList'
  params['Version'] = '2019-04-10'
  return request({
    url: '/iothub',
    method: 'get',
    params
  })
}
