export default {
  QueryMessageInfo: msgId => {
    var data = {
      MessageInfo: '{"name": "bobo"}'
    }
    return data;
  },
  CreateProduct: params => {
    console.log('params : ', params)
    var ProductName = params.ProductName
    var ProductDesc = params.Description
    var data = {
      Code: '200',
      Message: 'Create Product Success',
      ProductInfo: {
        ProductKey: 'test'
      }
    }
    return data;
  },
  QueryProduct: params => {
    var data = {}
    return data
  },
  QueryProductList: params => {
    var data = {}
    return data
  },
  RegisterDevices: params => {
    var data = {}
    return data
  },
  QueryDeviceList: params => {
    var data = {}
    return data
  }
}
