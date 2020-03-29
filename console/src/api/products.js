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

export function deleteProduct(params) {
  params['Action'] = 'DeleteProduct'
  params['Version'] = '2019-04-10'
  return request({
    url: '/iothub',
    method: 'get',
    params
  })
}

export function modifyProduct(params) {
  params['Action'] = 'UpdateProduct'
  params['Version'] = '2019-04-10'
  return request({
    url: '/iothub',
    method: 'get',
    params
  })
}
