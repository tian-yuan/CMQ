import { param2Obj } from './utils'
import openapi from './openapi'

export default {
  openapi: res => {
    const params = param2Obj(res.url)
    console.log('params : ' + params.toString())
    if (params.Action === 'CreateProduct') {
      return openapi.CreateProduct(params)
    }
    if (params.Action === 'QueryProduct') {
      return openapi.QueryProduct(params)
    }
    if (params.Action === 'QueryProductList') {
      return openapi.QueryProductList(params)
    }
    if (params.Action === 'RegisterDevices') {
      return openapi.RegisterDevices(params)
    }
    if (params.Action === 'QueryDeviceList') {
      return openapi.QueryDeviceList(params)
    }
  }
}

