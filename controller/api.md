### IOT CONTROLLER API

## 注意事项

- 请求为GET或POST请求
- 服务名称为iothub，所有接口path保持一致，不根据path来区分资源及操作
- 使用参数Action来区分具体操作，该参数放到url query string部分。Action名称格式为首字母大写，驼峰格式
- 使用参数Version来区分接口版本号，该参数放到url query string部分。Version统一为2018-11-01
- 响应字段Code使用英文，Message也使用英文，详细说明见附录-状态码
- 返回内容遵循最小权限原则，仅返回客户端（尤其是G0网关）需要的内容
- 所有业务参数和响应参数名称格式为：首字母大写、驼峰格式

## 目录

- [1. 概述](#1-概述)
- [2. 修订记录](#2-修订记录)
- [3. 产品管理API](#3-产品管理api)
- [3.1 创建产品](#31-创建产品)
- [3.2 查询产品](#32-查询产品)
- [3.3 查询产品列表](#33-查询产品列表)
- [3.4 更新产品](#34-更新产品)
- [3.5 删除产品](#35-删除产品)
- [3.6 查询产品配额](#36-查询产品配额)
- [3.7 修改产品配额-G1](#37-修改产品配额-g1)
- [4. 设备管理API](#4-设备管理api)
- [4.1 注册设备](#41-注册设备)
- [4.2 查询申请处理状态](#42-查询申请处理状态)
- [4.3 查询注册成功的设备列表](#43-查询注册成功的设备列表)
- [4.4 查询设备](#44-查询设备)
- [4.5 查询设备列表](#45-查询设备列表)
- [4.6 删除设备](#46-删除设备)
- [4.7 查询设备配额](#47-查询设备配额)
- [4.8 修改设备配额-G1](#48-修改设备配额-g1)
- [5. TOPIC类管理API](#5-topic类管理api)
- [5.1 创建自定义TOPIC类](#51-创建自定义topic类)
- [5.2 查询自定义TOPIC类](#52-查询自定义topic类)
- [5.3 查询自定义TOPIC类列表](#53-查询自定义topic类列表)
- [5.4 更新自定义TOPIC类](#54-更新自定义topic类)
- [5.5 删除自定义TOPIC类](#55-删除自定义topic类)
- [5.6 查询自定义TOPIC类配额](#56-查询自定义topic类配额)
- [5.7 修改自定义TOPIC类配额-G1](#57-修改自定义topic类配额-g1)
- [5.8 创建系统TOPIC类-G1](#58-创建系统topic类-g1)
- [5.9 查询系统TOPIC类](#59-查询系统topic类)
- [5.10 查询系统TOPIC类列表](#510-查询系统topic类列表)
- [5.11 更新系统TOPIC类-G1](#511-更新系统topic类-g1)
- [5.12 删除系统TOPIC类-G1](#512-删除系统topic类-g1)
- [6. 消息通信API](#6-消息通信api)
- [6.1 发布消息](#61-发布消息)
- [6.2 查询消息内容](#62-查询消息内容)
- [6.3 查询消息长度配额](#63-查询消息长度配额)
- [6.4 修改消息长度配额-G1](#64-修改消息长度配额-g1)
- [6.5 RRPC](#65-rrpc)
- [7. 其他服务依赖的API](#7-其他服务依赖的api)
- [7.1 查询KAFKA地址-G1](#71-查询kafka地址-g1)
- [8. 状态码](#8-状态码)
- [9. 数据结构](#9-数据结构)
- [9.1 ProductInfo](#91-productinfo)
- [9.2 DeviceInfo](#92-deviceinfo)
- [9.3 TopicInfo](#93-topicinfo)
- [10. 修改记录](#10-修改记录)

## 1 概述

```
iothub，IoT hub service.
```

## 2 修订记录

```
20181101 1101版本OPENAPI
```

## 3 产品管理API

### 3.1 创建产品

#### 描述

```
创建产品，同步
```

#### URL路径

```
/iothub
```

#### 请求方法

```
GET
```

#### 请求头部

```
* X-Product-Id
    * 描述：租户ID
    * 类型：String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes
```

#### QUERY STRING

```
* Action
    * 描述：操作代码
    * 类型：String
    * 默认值：无
    * 约束：固定值-CreateProduct
    * 是否必须：Yes
* Version
    * 描述：版本号
    * 类型：String
    * 默认值：无
    * 约束：固定值-2018-11-01
    * 是否必须：Yes
* ProductName
    * 描述：产品名称
    * 类型：String
    * 默认值：无
    * 约束：长度大于等于1，小于等于63
    * 是否必须：Yes
* Description
    * 描述：产品描述
    * 类型：String
    * 默认值：空
    * 约束：长度大于等于0，小于等于100
    * 是否必须：No
```

#### 请求BODY

```
无
```

#### 响应状态码

```
200：正常处理
其它：异常处理
```

#### 响应BODY

```
* Code
    * 描述：错误码
    * 类型: String
    * 是否必须：Yes
* Message
    * 描述：错误详细信息
    * 类型: String
    * 是否必须：Yes
* ProductInfo
    * 描述：产品信息json对象，响应状态码等于200的时候才有意义
    * 类型：ProductInfo object，详见附录-数据结构
    * 是否必须：Yes
```

### 3.2 查询产品

#### 描述

```
查询产品信息，同步
```

#### URL路径

```
/iothub
```

#### 请求方法

```
GET
```

#### 请求头部

```
* X-Product-Id
    * 描述：租户ID
    * 类型：String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes
```

#### QUERY STRING

```
* Action
    * 描述：操作代码
    * 类型：String
    * 默认值：无
    * 约束：固定值-QueryProduct
    * 是否必须：Yes
* Version
    * 描述：版本号
    * 类型：String
    * 默认值：无
    * 约束：固定值-2018-11-01
    * 是否必须：Yes
* ProductKey
    * 描述：产品KEY
    * 类型：String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes
```

#### 请求BODY

```
无
```

#### 响应状态码

```
200：正常处理
其它：异常处理
```

#### 响应BODY

```
* Code
    * 描述：错误码
    * 类型: String
    * 是否必须：Yes
* Message
    * 描述：错误详细信息
    * 类型: String
    * 是否必须：Yes
* ProductInfo
    * 描述：产品信息JSON对象，响应状态码等于200的时候才有意义
    * 类型：ProductInfo object，详见附录-数据结构
    * 是否必须：Yes
```

### 3.3 查询产品列表

#### 描述

```
查询产品列表，同步
```

#### URL路径

```
/iothub
```

#### 请求方法

```
GET
```

#### 请求头部

```
* X-Product-Id
    * 描述：租户ID
    * 类型：String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes
```

#### QUERY STRING

```
* Action
    * 描述：操作代码
    * 类型：String
    * 默认值：无
    * 约束：固定值-QueryProductList
    * 是否必须：Yes
* Version
    * 描述：版本号
    * 类型：String
    * 默认值：无
    * 约束：固定值-2018-11-01
    * 是否必须：Yes
* Offset
    * 描述：起始记录偏移
    * 类型：Integer
    * 默认值：0
    * 约束：合法取值区间[0, Integer.MAX]
    * 是否必须：No
* Limit
    * 描述：最大获取条数
    * 类型：Integer
    * 默认值：10
    * 约束：合法取值区间[1, 100]
    * 是否必须：No
* Keyword
    * 描述：查询关键字，支持ProductName，模糊匹配，不区分大小写
    * 类型：String
    * 默认值：无
    * 约束：长度大于等于1，小于等于63
    * 是否必须：No
```

#### 请求BODY

```
无
```

#### 响应状态码

```
200：正常处理
其它：异常处理
```

#### 响应BODY

```
* Code
    * 描述：错误码
    * 类型: String
    * 是否必须：Yes
* Message
    * 描述：错误详细信息
    * 类型: String
    * 是否必须：Yes
* TotalCount
    * 描述：产品总数，不考虑分页
    * 类型：Integer
    * 是否必须：Yes
* ProductInfoList
    * 描述：产品信息JSON对象列表，响应状态码等于200的时候才有意义
    * 类型：Array of ProductInfo object，详见附录-数据结构
    * 是否必须：Yes
```

### 3.4 更新产品

#### 描述

```
更新产品信息，同步
```

#### URL路径

```
/iothub
```

#### 请求方法

```
GET
```

#### 请求头部

```
* X-Product-Id
    * 描述：租户ID
    * 类型：String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes
```

#### QUERY STRING

```
* Action
    * 描述：操作代码
    * 类型：String
    * 默认值：无
    * 约束：固定值-UpdateProduct
    * 是否必须：Yes
* Version
    * 描述：版本号
    * 类型：String
    * 默认值：无
    * 约束：固定值-2018-11-01
    * 是否必须：Yes
* ProductKey
    * 描述：产品KEY
    * 类型：String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes
* ProductName
    * 描述：产品名称
    * 类型：String
    * 默认值：无
    * 约束：长度大于等于1，小于等于63
    * 是否必须：No
* Description
    * 描述：产品描述
    * 类型：String
    * 默认值：空
    * 约束：长度大于等于0，小于等于100
    * 是否必须：No

```

#### 请求BODY

```
无

```

#### 响应状态码

```
200：正常处理
其它：异常处理

```

#### 响应BODY

```
* Code
    * 描述：错误码
    * 类型: String
    * 是否必须：Yes
* Message
    * 描述：错误详细信息
    * 类型: String
    * 是否必须：Yes
* ProductInfo
    * 描述：产品信息JSON对象，响应状态码等于200的时候才有意义
    * 类型：ProductInfo object，详见附录-数据结构
    * 是否必须：Yes

```

### 3.5 删除产品

#### 描述

```
删除产品，异步，仅在产品无关联的设备和TOPIC类时才可以删除
删除操作不在数据库中删记录，只把状态标记为删除

```

#### URL路径

```
/iothub

```

#### 请求方法

```
GET

```

#### 请求头部

```
* X-Product-Id
    * 描述：租户ID
    * 类型：String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes

```

#### QUERY STRING

```
* Action
    * 描述：操作代码
    * 类型：String
    * 默认值：无
    * 约束：固定值-DeleteProduct
    * 是否必须：Yes
* Version
    * 描述：版本号
    * 类型：String
    * 默认值：无
    * 约束：固定值-2018-11-01
    * 是否必须：Yes
* ProductKey
    * 描述：产品KEY
    * 类型：String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes

```

#### 请求BODY

```
无

```

#### 响应状态码

```
200：正常处理
其它：异常处理

```

#### 响应BODY

```
* Code
    * 描述：错误码
    * 类型: String
    * 是否必须：Yes
* Message
    * 描述：错误详细信息
    * 类型: String
    * 是否必须：Yes

```

### 3.6 查询产品配额

#### 描述

```
查询产品数量配额，同步

```

#### URL路径

```
/iothub

```

#### 请求方法

```
GET

```

#### 请求头部

```
* X-Product-Id
    * 描述：租户ID
    * 类型：String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes
```

#### QUERY STRING

```
* Action
    * 描述：操作代码
    * 类型：String
    * 默认值：无
    * 约束：固定值-QueryProductQuota
    * 是否必须：Yes
* Version
    * 描述：版本号
    * 类型：String
    * 默认值：无
    * 约束：固定值-2018-11-01
    * 是否必须：Yes

```

#### 请求BODY

```
无

```

#### 响应状态码

```
200：正常处理
其它：异常处理

```

#### 响应BODY

```
* Code
    * 描述：错误码
    * 类型: String
    * 是否必须：Yes
* Message
    * 描述：错误详细信息
    * 类型: String
    * 是否必须：Yes
* UsedQuota
    * 描述：已占用数量配额，即已有产品数量，响应状态码等于200的时候才有意义
    * 类型: Integer
    * 是否必须：Yes
* Quota
    * 描述：产品数量配额，默认值20，响应状态码等于200的时候才有意义
    * 类型: Integer
    * 是否必须：Yes

```

### 3.7 修改产品配额-G1

#### 描述

```
修改产品数量配额，同步
G1接口，只面向运营平台

```

#### URL路径

```
/iothub

```

#### 请求方法

```
GET

```

#### 请求头部

```
* X-Origin-Service
    * 描述：调用方服务名
    * 类型：String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes

```

#### QUERY STRING

```
* Action
    * 描述：操作代码
    * 类型：String
    * 默认值：无
    * 约束：固定值-ModifyProductQuota
    * 是否必须：Yes
* Version
    * 描述：版本号
    * 类型：String
    * 默认值：无
    * 约束：固定值-2018-11-01
    * 是否必须：Yes
* TenantId
    * 描述：租户ID
    * 类型：String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes
* Quota
    * 描述：产品数量配额
    * 类型：Integer
    * 默认值：无
    * 约束：合法取值区间[0, Integer.MAX]
    * 是否必须：Yes

```

#### 请求BODY

```
无

```

#### 响应状态码

```
200：正常处理
其它：异常处理

```

#### 响应BODY

```
* Code
    * 描述：错误码
    * 类型: String
    * 是否必须：Yes
* Message
    * 描述：错误详细信息
    * 类型: String
    * 是否必须：Yes
* UsedQuota
    * 描述：已占用数量配额，即已有产品数量，响应状态码等于200的时候才有意义
    * 类型: Integer
    * 是否必须：Yes
* Quota
    * 描述：产品数量配额，默认值20，响应状态码等于200的时候才有意义
    * 类型: Integer
    * 是否必须：Yes

```

### 3.8 检查产品名

#### 描述

```
检查产品名，判断产品名是否有效（长度、字符编码等），是否重复
```

#### URL路径

```
/iothub


```

#### 请求方法

```
GET


```

#### 请求头部

```
* X-Product-Id
    * 描述：租户ID
    * 类型：String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes

```

#### QUERY STRING

```
* Action
    * 描述：操作代码
    * 类型：String
    * 默认值：无
    * 约束：固定值-CheckProductName
    * 是否必须：Yes
* Version
    * 描述：版本号
    * 类型：String
    * 默认值：无
    * 约束：固定值-2018-11-01
    * 是否必须：Yes

```

#### 请求BODY

```
无


```

#### 响应状态码

```
200：正常处理
其它：异常处理

```

#### 响应BODY

```
* Code
    * 描述：错误码
    * 类型: String
    * 是否必须：Yes
* Message
    * 描述：错误详细信息
    * 类型: String
    * 是否必须：Yes
* Status
    * 描述：描述产品名状态
    * 类型: String，"VALID": 产品名无效，"EXISTED": 产品名已经存在，"INVALID": 产品名无效
    * 是否必须：Yes

```



## 4 设备管理API

### 4.1 注册设备

#### 描述

```
申请注册设备，异步

```

#### URL路径

```
/iothub

```

#### 请求方法

```
POST

```

#### 请求头部

```
* X-Product-Id
    * 描述：租户ID
    * 类型：String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes

```

#### QUERY STRING

```
* Action
    * 描述：操作代码
    * 类型：String
    * 默认值：无
    * 约束：固定值-RegisterDevices
    * 是否必须：Yes
* Version
    * 描述：版本号
    * 类型：String
    * 默认值：无
    * 约束：固定值-2018-11-01
    * 是否必须：Yes
* ProductKey
    * 描述：产品KEY
    * 类型：String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes

```

#### 请求BODY

```
* Count
    * 描述：注册设备数
    * 类型: Integer
    * 默认值：无
    * 约束：合法取值区间[1, 1000]
    * 是否必须：Yes
* DeviceNames
    * 描述：设备名字数组，设备名字长度大于等于1，小于等于63
    * 类型: Array of string
    * 默认值：无
    * 约束：数组长度需与Count参数匹配
    * 是否必须：No

```

#### 响应状态码

```
200：正常处理
其它：异常处理

```

#### 响应BODY

```
* Code
    * 描述：错误码
    * 类型: String
    * 是否必须：Yes
* Message
    * 描述：错误详细信息
    * 类型: String
    * 是否必须：Yes
* ApplyId
    * 描述：申请ID，响应状态码等于200的时候才有意义
    * 类型: String
    * 是否必须：Yes

```

### 4.2 查询申请处理状态

#### 描述

```
查询注册申请处理状态，同步

```

#### URL路径

```
/iothub

```

#### 请求方法

```
GET

```

#### 请求头部

```
* X-Product-Id
    * 描述：租户ID
    * 类型：String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes

```

#### QUERY STRING

```
* Action
    * 描述：操作代码
    * 类型：String
    * 默认值：无
    * 约束：固定值-QueryApplyState
    * 是否必须：Yes
* Version
    * 描述：版本号
    * 类型：String
    * 默认值：无
    * 约束：固定值-2018-11-01
    * 是否必须：Yes
* ProductKey
    * 描述：产品KEY
    * 类型：String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes
* ApplyId
    * 描述：申请ID
    * 类型: String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes

```

#### 请求BODY

```
无

```

#### 响应状态码

```
200：正常处理
其它：异常处理

```

#### 响应BODY

```
* Code
    * 描述：错误码
    * 类型: String
    * 是否必须：Yes
* Message
    * 描述：错误详细信息
    * 类型: String
    * 是否必须：Yes
* State
    * 描述：状态，Processing-处理中，Complete-处理完成，响应状态码等于200的时候才有意义
    * 类型: String
    * 是否必须：Yes
* InvalidNames
    * 描述：若申请设备时提供了DeviceNames参数，则返回其中创建失败的DeviceName
    * 类型: Array of string
    * 是否必须：Yes

```

### 4.3 查询注册成功的设备列表

#### 描述

```
查询注册成功的设备列表，同步

```

#### URL路径

```
/iothub

```

#### 请求方法

```
GET

```

#### 请求头部

```
* X-Product-Id
    * 描述：租户ID
    * 类型：String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes
```

#### QUERY STRING

```
* Action
    * 描述：操作代码
    * 类型：String
    * 默认值：无
    * 约束：固定值-QueryDeviceListByApplyId
    * 是否必须：Yes
* Version
    * 描述：版本号
    * 类型：String
    * 默认值：无
    * 约束：固定值-2018-11-01
    * 是否必须：Yes
* ProductKey
    * 描述：产品KEY
    * 类型：String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes
* ApplyId
    * 描述：申请ID
    * 类型: String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes
* Offset
    * 描述：起始记录偏移
    * 类型：Integer
    * 默认值：0
    * 约束：合法取值区间[0, Integer.MAX]
    * 是否必须：No
* Limit
    * 描述：最大获取条数
    * 类型：Integer
    * 默认值：10
    * 约束：合法取值区间[1, 100]
    * 是否必须：No
* Keyword
    * 描述：查询关键字，支持DeviceName，模糊匹配，不区分大小写
    * 类型：String
    * 默认值：无
    * 约束：长度大于等于1，小于等于63
    * 是否必须：No
```

#### 请求BODY

```
无

```

#### 响应状态码

```
200：正常处理
其它：异常处理

```

#### 响应BODY

```
* Code
    * 描述：错误码
    * 类型: String
    * 是否必须：Yes
* Message
    * 描述：错误详细信息
    * 类型: String
    * 是否必须：Yes
* TotalCount
    * 描述：成功申请的设备总数，不考虑分页
    * 类型：Integer
    * 是否必须：Yes
* DeviceInfoList
    * 描述：设备信息JSON对象列表，响应状态码等于200的时候才有意义
    * 类型：Array of DeviceInfo object，详见附录-数据结构
    * 是否必须：Yes

```

### 4.4 查询设备

#### 描述

```
查询设备详细信息，同步

```

#### URL路径

```
/iothub
```

#### 请求方法

```
GET

```

#### 请求头部

```
* X-Product-Id
    * 描述：租户ID
    * 类型：String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes

```

#### QUERY STRING

```
* Action
    * 描述：操作代码
    * 类型：String
    * 默认值：无
    * 约束：固定值-QueryDeviceInfo
    * 是否必须：Yes
* Version
    * 描述：版本号
    * 类型：String
    * 默认值：无
    * 约束：固定值-2018-11-01
    * 是否必须：Yes
* ProductKey
    * 描述：产品KEY
    * 类型：String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes
* DeviceName
    * 描述：设备名字
    * 类型: String
    * 默认值：无
    * 约束：长度大于等于1，小于等于63
    * 是否必须：Yes

```

#### 请求BODY

```
无

```

#### 响应状态码

```
200：正常处理
其它：异常处理

```

#### 响应BODY

```
* Code
    * 描述：错误码
    * 类型: String
    * 是否必须：Yes
* Message
    * 描述：错误详细信息
    * 类型: String
    * 是否必须：Yes
* DeviceInfo
    * 描述：设备详细信息JSON对象
    * 类型：DeviceInfo object，详见附录-数据结构
    * 是否必须：Yes

```

### 4.5 查询设备列表

#### 描述

```
查询产品关联的设备列表，同步
```

#### URL路径

```
/iothub

```

#### 请求方法

```
GET

```

#### 请求头部

```
* X-Product-Id
    * 描述：租户ID
    * 类型：String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes
```

#### QUERY STRING

```
* Action
    * 描述：操作代码
    * 类型：String
    * 默认值：无
    * 约束：固定值-QueryDeviceList
    * 是否必须：Yes
* Version
    * 描述：版本号
    * 类型：String
    * 默认值：无
    * 约束：固定值-2018-11-01
    * 是否必须：Yes
* ProductKey
    * 描述：产品KEY
    * 类型：String
    * 默认值：无
    * 约束：无
    * 是否必须：No
* Offset
    * 描述：起始记录偏移
    * 类型：Integer
    * 默认值：0
    * 约束：合法取值区间[0, Integer.MAX]
    * 是否必须：No
* Limit
    * 描述：最大获取条数
    * 类型：Integer
    * 默认值：10
    * 约束：合法取值区间[1, 100]
    * 是否必须：No
* Keyword
    * 描述：查询关键字，支持DeviceName，模糊匹配，不区分大小写
    * 类型：String
    * 默认值：无
    * 约束：长度大于等于1，小于等于63
    * 是否必须：No
```

#### 请求BODY

```
无

```

#### 响应状态码

```
200：正常处理
其它：异常处理

```

#### 响应BODY

```
* Code
    * 描述：错误码
    * 类型: String
    * 是否必须：Yes
* Message
    * 描述：错误详细信息
    * 类型: String
    * 是否必须：Yes
* TotalCount
    * 描述：该产品的设备总数，不考虑分页
    * 类型：Integer
    * 是否必须：Yes
* DeviceInfoList
    * 描述：设备详细信息JSON对象列表，响应状态码等于200的时候才有意义
    * 类型：Array of DeviceInfo object，详见附录-数据结构
    * 是否必须：Yes

```

### 4.6 查询设备topic列表

#### 描述

```
查询设备的topic列表，同步
```

#### URL路径

```
/iothub

```

#### 请求方法

```
GET

```

#### 请求头部

```
* X-Product-Id
    * 描述：租户ID
    * 类型：String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes
```

#### QUERY STRING

```
* Action
    * 描述：操作代码
    * 类型：String
    * 默认值：无
    * 约束：固定值-QueryTopicList
    * 是否必须：Yes
* Version
    * 描述：版本号
    * 类型：String
    * 默认值：无
    * 约束：固定值-2018-11-01
    * 是否必须：Yes
* ProductKey
    * 描述：产品KEY
    * 类型：String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes
* DeviceName
    * 描述：设备名字
    * 类型: String
    * 默认值：无
    * 约束：长度大于等于1，小于等于63
    * 是否必须：Yes
* Offset
    * 描述：起始记录偏移
    * 类型：Integer
    * 默认值：0
    * 约束：合法取值区间[0, Integer.MAX]
    * 是否必须：No
* Limit
    * 描述：最大获取条数
    * 类型：Integer
    * 默认值：10
    * 约束：合法取值区间[1, 100]
    * 是否必须：No
* Keyword
    * 描述：查询关键字，支持TopicName，模糊匹配，不区分大小写
    * 类型：String
    * 默认值：无
    * 约束：长度大于等于1，小于等于128
    * 是否必须：No
```

#### 请求BODY

```
无

```

#### 响应状态码

```
200：正常处理
其它：异常处理

```

#### 响应BODY

```
* Code
    * 描述：错误码
    * 类型: String
    * 是否必须：Yes
* Message
    * 描述：错误详细信息
    * 类型: String
    * 是否必须：Yes
* TotalCount
    * 描述：设备TOPIC总数，不考虑分页
    * 类型：Integer
    * 是否必须：Yes
* TopicInfoList
    * 描述：设备topic对象列表，响应状态码等于200的时候才有意义
    * 类型：Array of TopicInfo object，详见附录-数据结构
    * 是否必须：Yes

```

### 4.7 删除设备

#### 描述

```
删除单个设备，异步

```

#### URL路径

```
/iothub

```

#### 请求方法

```
GET

```

#### 请求头部

```
* X-Product-Id
    * 描述：租户ID
    * 类型：String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes

```

#### QUERY STRING

```
* Action
    * 描述：操作代码
    * 类型：String
    * 默认值：无
    * 约束：固定值-DeleteDevice
    * 是否必须：Yes
* Version
    * 描述：版本号
    * 类型：String
    * 默认值：无
    * 约束：固定值-2018-11-01
    * 是否必须：Yes
* ProductKey
    * 描述：产品KEY
    * 类型：String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes
* DeviceName
    * 描述：设备名字
    * 类型: String
    * 默认值：无
    * 约束：长度大于等于1，小于等于63
    * 是否必须：Yes
```

#### 请求BODY

```
无

```

#### 响应状态码

```
200：正常处理
其它：异常处理

```

#### 响应BODY

```
* Code
    * 描述：错误码
    * 类型: String
    * 是否必须：Yes
* Message
    * 描述：错误详细信息
    * 类型: String
    * 是否必须：Yes

```



### 4.8 查询设备配额

#### 描述

```
查询具体产品的可注册设备总数配额，同步

```

#### URL路径

```
/iothub

```

#### 请求方法

```
GET

```

#### 请求头部

```
* X-Product-Id
    * 描述：租户ID
    * 类型：String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes

```

#### QUERY STRING

```
* Action
    * 描述：操作代码
    * 类型：String
    * 默认值：无
    * 约束：固定值-QueryDeviceQuota
    * 是否必须：Yes
* Version
    * 描述：版本号
    * 类型：String
    * 默认值：无
    * 约束：固定值-2018-11-01
    * 是否必须：Yes
* ProductKey
    * 描述：产品KEY
    * 类型：String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes

```

#### 请求BODY

```
无

```

#### 响应状态码

```
200：正常处理
其它：异常处理

```

#### 响应BODY

```
* Code
    * 描述：错误码
    * 类型: String
    * 是否必须：Yes
* Message
    * 描述：错误详细信息
    * 类型: String
    * 是否必须：Yes
* UsedQuota
    * 描述：已占用数量配额，即该产品已注册设备总数，响应状态码等于200的时候才有意义
    * 类型: Integer
    * 是否必须：Yes
* Quota
    * 描述：设备数量配额，默认值10000，响应状态码等于200的时候才有意义
    * 类型: Integer
    * 是否必须：Yes

```

### 4.9 修改设备配额-G1

#### 描述

```
修改产品可注册设备数量配额，同步
G1接口，只面向运营平台

```

#### URL路径

```
/iothub

```

#### 请求方法

```
GET

```

#### 请求头部

```
* X-Origin-Service
    * 描述：调用方服务名
    * 类型：String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes

```

#### QUERY STRING

```
* Action
    * 描述：操作代码
    * 类型：String
    * 默认值：无
    * 约束：固定值-ModifyDeviceQuota
    * 是否必须：Yes
* Version
    * 描述：版本号
    * 类型：String
    * 默认值：无
    * 约束：固定值-2018-11-01
    * 是否必须：Yes
* TenantId
    * 描述：租户ID
    * 类型：String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes
* ProductKey
    * 描述：产品KEY
    * 类型：String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes
* Quota
    * 描述：设备数量配额
    * 类型：Integer
    * 默认值：无
    * 约束：合法取值区间[0, Integer.MAX]
    * 是否必须：Yes

```

#### 请求BODY

```
无

```

#### 响应状态码

```
200：正常处理
其它：异常处理

```

#### 响应BODY

```
* Code
    * 描述：错误码
    * 类型: String
    * 是否必须：Yes
* Message
    * 描述：错误详细信息
    * 类型: String
    * 是否必须：Yes
* UsedQuota
    * 描述：已占用数量配额，即该产品已注册设备总数，响应状态码等于200的时候才有意义
    * 类型: Integer
    * 是否必须：Yes
* Quota
    * 描述：设备数量配额，默认值10000，响应状态码等于200的时候才有意义
    * 类型: Integer
    * 是否必须：Yes

```

### 4.10 检查DeviceName

#### 描述

```
检查 DeviceName，判断 DeviceName 是否有效（长度、字符编码等），是否重复
```

#### URL路径

```
/iothub


```

#### 请求方法

```
GET


```

#### 请求头部

```
* X-Product-Id
    * 描述：租户ID
    * 类型：String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes

```

#### QUERY STRING

```
* Action
    * 描述：操作代码
    * 类型：String
    * 默认值：无
    * 约束：固定值-CheckDeviceName
    * 是否必须：Yes
* Version
    * 描述：版本号
    * 类型：String
    * 默认值：无
    * 约束：固定值-2018-11-01
    * 是否必须：Yes

```

#### 请求BODY

```
无



```

#### 响应状态码

```
200：正常处理
其它：异常处理


```

#### 响应BODY

```
* Code
    * 描述：错误码
    * 类型: String
    * 是否必须：Yes
* Message
    * 描述：错误详细信息
    * 类型: String
    * 是否必须：Yes
* Status
    * 描述：描述产品名状态
    * 类型: String，"VALID": DeviceName 有效，"EXISTED": DeviceName 已经存在，"INVALID": DeviceName 无效
    * 是否必须：Yes

```



## 5 TOPIC类管理API

### 5.1 创建TOPIC类

#### 描述

```
创建产品TOPIC类，同步
```

#### URL路径

```
/iothub

```

#### 请求方法

```
GET

```

#### 请求头部

```
* X-Product-Id
    * 描述：租户ID
    * 类型：String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes
```

#### QUERY STRING

```
* Action
    * 描述：操作代码
    * 类型：String
    * 默认值：无
    * 约束：固定值-CreateTopicClass
    * 是否必须：Yes
* Version
    * 描述：版本号
    * 类型：String
    * 默认值：无
    * 约束：固定值-2018-11-01
    * 是否必须：Yes
* ProductKey
    * 描述：产品KEY
    * 类型：String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes
* TopicName
    * 描述：TOPIC类名字，格式须遵从TOPIC类命名规范
    * 类型：String
    * 默认值：无
    * 约束：长度大于等于1，小于等于128，且斜杠数量小于等于7
    * 是否必须：Yes
* TopicType
    * 描述：TOPIC类型，CUSTOM-产品自定义，SYSTEM-系统保留，OTA-空中升级
    * 类型：String
    * 是否必须：Yes
* Operation
    * 描述：设备对TOPIC类的许可操作
    * 类型: String
    * 默认值：无
    * 约束：可选值[SUB, PUB, ALL]，SUB-订阅，PUB-发布，ALL-订阅和发布
    * 是否必须：Yes
* Qos
    * 描述：设备操作TOPIC时所允许的最大QOS等级
    * 类型：Integer
    * 默认值：无
    * 约束：可选值[0, 1]，0-MQTT QOS等级0，1-MQTT QOS等级1
    * 是否必须：Yes
* Description
    * 描述：TOPIC类描述
    * 类型：String
    * 默认值：空
    * 约束：长度大于等于0，小于等于100
    * 是否必须：No
```

#### 请求BODY

```
无

```

#### 响应状态码

```
200：正常处理
其它：异常处理

```

#### 响应BODY

```
* Code
    * 描述：错误码
    * 类型: String
    * 是否必须：Yes
* Message
    * 描述：错误详细信息
    * 类型: String
    * 是否必须：Yes
* TopicId
    * 描述：TOPIC类ID，响应状态码等于200的时候才有意义
    * 类型: String
    * 是否必须：Yes

```

### 5.2 查询TOPIC类

#### 描述

```
查询产品自定义的TOPIC类详细信息，同步

```

#### URL路径

```
/iothub

```

#### 请求方法

```
GET

```

#### 请求头部

```
* X-Product-Id
    * 描述：租户ID
    * 类型：String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes

```

#### QUERY STRING

```
* Action
    * 描述：操作代码
    * 类型：String
    * 默认值：无
    * 约束：固定值-QueryTopicClass
    * 是否必须：Yes
* Version
    * 描述：版本号
    * 类型：String
    * 默认值：无
    * 约束：固定值-2018-11-01
    * 是否必须：Yes
* ProductKey
    * 描述：产品KEY
    * 类型：String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes
* TopicId
    * 描述：TOPIC类ID
    * 类型：String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes
```

#### 请求BODY

```
无

```

#### 响应状态码

```
200：正常处理
其它：异常处理

```

#### 响应BODY

```
* Code
    * 描述：错误码
    * 类型: String
    * 是否必须：Yes
* Message
    * 描述：错误详细信息
    * 类型: String
    * 是否必须：Yes
* TopicInfo
    * 描述：TOPIC类信息JSON对象，响应状态码等于200的时候才有意义
    * 类型：TopicInfo object，详见附录-数据结构
    * 是否必须：Yes

```

### 5.3 查询TOPIC类列表

#### 描述

```
查询产品自定义的TOPIC类列表，同步

```

#### URL路径

```
/iothub

```

#### 请求方法

```
GET

```

#### 请求头部

```
* X-Product-Id
    * 描述：租户ID
    * 类型：String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes

```

#### QUERY STRING

```
* Action
    * 描述：操作代码
    * 类型：String
    * 默认值：无
    * 约束：固定值-QueryTopicClassList
    * 是否必须：Yes
* Version
    * 描述：版本号
    * 类型：String
    * 默认值：无
    * 约束：固定值-2018-11-01
    * 是否必须：Yes
* ProductKey
    * 描述：产品KEY
    * 类型：String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes
* Offset
    * 描述：起始记录偏移
    * 类型：Integer
    * 默认值：0
    * 约束：合法取值区间[0, Integer.MAX]
    * 是否必须：No
* Limit
    * 描述：最大获取条数
    * 类型：Integer
    * 默认值：10
    * 约束：合法取值区间[1, 100]
    * 是否必须：No
* Keyword
    * 描述：查询关键字，支持TopicName，模糊匹配，不区分大小写
    * 类型：String
    * 默认值：无
    * 约束：长度大于等于1，小于等于128
    * 是否必须：No
```

#### 请求BODY

```
无

```

#### 响应状态码

```
200：正常处理
其它：异常处理

```

#### 响应BODY

```
* Code
    * 描述：错误码
    * 类型: String
    * 是否必须：Yes
* Message
    * 描述：错误详细信息
    * 类型: String
    * 是否必须：Yes
* TotalCount
    * 描述：产品自定义的TOPIC类总数，不考虑分页
    * 类型：Integer
    * 是否必须：Yes
* TopicClassInfoList
    * 描述：TOPIC类信息JSON对象列表，响应状态码等于200的时候才有意义
    * 类型：Array of TopicClassInfo object，详见附录-数据结构
    * 是否必须：Yes

```

### 5.4 更新TOPIC类

#### 描述

```
更新产品TOPIC类，同步
```

#### URL路径

```
/iothub

```

#### 请求方法

```
GET

```

#### 请求头部

```
* X-Product-Id
    * 描述：租户ID
    * 类型：String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes

```

#### QUERY STRING

```
* Action
    * 描述：操作代码
    * 类型：String
    * 默认值：无
    * 约束：固定值-UpdateTopicClass
    * 是否必须：Yes
* Version
    * 描述：版本号
    * 类型：String
    * 默认值：无
    * 约束：固定值-2018-11-01
    * 是否必须：Yes
* ProductKey
    * 描述：产品KEY
    * 类型：String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes
* TopicId
    * 描述：TOPIC类ID
    * 类型：String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes
* TopicName
    * 描述：TOPIC类名字，格式须遵从TOPIC类命名规范
    * 类型：String
    * 默认值：无
    * 约束：长度大于等于1，小于等于128，且斜杠数量小于等于7
    * 是否必须：No
* Operation
    * 描述：设备对TOPIC类的许可操作
    * 类型: String
    * 默认值：无
    * 约束：可选值[SUB, PUB, ALL]，SUB-订阅，PUB-发布，ALL-订阅和发布
    * 是否必须：No
* Qos
    * 描述：设备操作TOPIC时所允许的最大QOS等级
    * 类型：Integer
    * 默认值：无
    * 约束：可选值[0, 1]，0-MQTT QOS等级0，1-MQTT QOS等级1
    * 是否必须：No
* Description
    * 描述：TOPIC类描述
    * 类型：String
    * 默认值：空
    * 约束：长度大于等于0，小于等于100
    * 是否必须：No
```

#### 请求BODY

```
无

```

#### 响应状态码

```
200：正常处理
其它：异常处理

```

#### 响应BODY

```
* Code
    * 描述：错误码
    * 类型: String
    * 是否必须：Yes
* Message
    * 描述：错误详细信息
    * 类型: String
    * 是否必须：Yes

```

### 5.5 删除TOPIC类

#### 描述

```
删除产品TOPIC类，同步
```

#### URL路径

```
/iothub
```

#### 请求方法

```
GET

```

#### 请求头部

```
* X-Product-Id
    * 描述：租户ID
    * 类型：String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes

```

#### QUERY STRING

```
* Action
    * 描述：操作代码
    * 类型：String
    * 默认值：无
    * 约束：固定值-DeleteTopicClass
    * 是否必须：Yes
* Version
    * 描述：版本号
    * 类型：String
    * 默认值：无
    * 约束：固定值-2018-11-01
    * 是否必须：Yes
* ProductKey
    * 描述：产品KEY
    * 类型：String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes
* TopicId
    * 描述：TOPIC类ID
    * 类型：String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes
```

#### 请求BODY

```
无

```

#### 响应状态码

```
200：正常处理
其它：异常处理

```

#### 响应BODY

```
* Code
    * 描述：错误码
    * 类型: String
    * 是否必须：Yes
* Message
    * 描述：错误详细信息
    * 类型: String
    * 是否必须：Yes

```

### 5.6 查询TOPIC类配额

#### 描述

```
查询允许产品TOPIC类数量配额，同步
```

#### URL路径

```
/iothub
```

#### 请求方法

```
GET
```

#### 请求头部

```
* X-Product-Id
    * 描述：租户ID
    * 类型：String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes

```

#### QUERY STRING

```
* Action
    * 描述：操作代码
    * 类型：String
    * 默认值：无
    * 约束：固定值-QueryTopicClassQuota
    * 是否必须：Yes
* Version
    * 描述：版本号
    * 类型：String
    * 默认值：无
    * 约束：固定值-2018-11-01
    * 是否必须：Yes
* ProductKey
    * 描述：产品KEY
    * 类型：String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes

```

#### 请求BODY

```
无

```

#### 响应状态码

```
200：正常处理
其它：异常处理

```

#### 响应BODY

```
* Code
    * 描述：错误码
    * 类型: String
    * 是否必须：Yes
* Message
    * 描述：错误详细信息
    * 类型: String
    * 是否必须：Yes
* UsedQuota
    * 描述：已占用数量配额，即该产品已有TOPIC类数量，响应状态码等于200的时候才有意义
    * 类型: Integer
    * 是否必须：Yes
* Quota
    * 描述：TOPIC类数量配额，默认值50，响应状态码等于200的时候才有意义
    * 类型: Integer
    * 是否必须：Yes

```

### 5.7 修改TOPIC类配额-G1

#### 描述

```
修改允许产品的TOPIC类数量配额，同步
G1接口，只面向运营平台
```

#### URL路径

```
/iothub

```

#### 请求方法

```
GET

```

#### 请求头部

```
* X-Origin-Service
    * 描述：调用方服务名
    * 类型：String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes

```

#### QUERY STRING

```
* Action
    * 描述：操作代码
    * 类型：String
    * 默认值：无
    * 约束：固定值-ModifyTopicClassQuota
    * 是否必须：Yes
* Version
    * 描述：版本号
    * 类型：String
    * 默认值：无
    * 约束：固定值-2018-11-01
    * 是否必须：Yes
* TenantId
    * 描述：租户ID
    * 类型：String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes
* ProductKey
    * 描述：产品KEY
    * 类型：String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes
* Quota
    * 描述：TOPIC类数量配额
    * 类型：Integer
    * 默认值：无
    * 约束：合法取值区间[0, 100]
    * 是否必须：Yes
```

#### 请求BODY

```
无

```

#### 响应状态码

```
200：正常处理
其它：异常处理

```

#### 响应BODY

```
* Code
    * 描述：错误码
    * 类型: String
    * 是否必须：Yes
* Message
    * 描述：错误详细信息
    * 类型: String
    * 是否必须：Yes
* UsedQuota
    * 描述：已占用数量配额，即该产品已有TOPIC类数量，响应状态码等于200的时候才有意义
    * 类型: Integer
    * 是否必须：Yes
* Quota
    * 描述：TOPIC类数量配额，默认值50，响应状态码等于200的时候才有意义
    * 类型: Integer
    * 是否必须：Yes
```

### 5.8 检查TopicName

#### 描述

```
检查 TopicName，判断 TopicName 是否有效（长度、字符编码等），是否重复
只针对自定义 Topic
```

#### URL路径

```
/iothub


```

#### 请求方法

```
GET



```

#### 请求头部

```
* X-Product-Id
    * 描述：租户ID
    * 类型：String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes


```

#### QUERY STRING

```
* Action
    * 描述：操作代码
    * 类型：String
    * 默认值：无
    * 约束：固定值-CheckTopicName
    * 是否必须：Yes
* Version
    * 描述：版本号
    * 类型：String
    * 默认值：无
    * 约束：固定值-2018-11-01
    * 是否必须：Yes
* ProductKey
    * 描述：产品KEY
    * 类型：String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes
```

#### 请求BODY

```
无




```

#### 响应状态码

```
200：正常处理
其它：异常处理



```

#### 响应BODY

```
* Code
    * 描述：错误码
    * 类型: String
    * 是否必须：Yes
* Message
    * 描述：错误详细信息
    * 类型: String
    * 是否必须：Yes
* Status
    * 描述：描述产品名状态
    * 类型: String，"VALID": TopicName 有效，"EXISTED": TopicName 已经存在，"INVALID": TopicName 无效
    * 是否必须：Yes

```



## 6 消息通信API

### 6.1 发布消息

#### 描述

```
服务端发布消息，异步

```

#### URL路径

```
/iothub

```

#### 请求方法

```
POST

```

#### 请求头部

```
* X-Product-Id
    * 描述：租户ID
    * 类型：String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes

```

#### QUERY STRING

```
* Action
    * 描述：操作代码
    * 类型：String
    * 默认值：无
    * 约束：固定值-PublishMessage
    * 是否必须：Yes
* Version
    * 描述：版本号
    * 类型：String
    * 默认值：无
    * 约束：固定值-2018-11-01
    * 是否必须：Yes

```

#### 请求BODY

```
* ProductKey
    * 描述：产品KEY
    * 类型：String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes
* Topic
    * 描述：发布的具体TOPIC
    * 类型: String
    * 默认值：无
    * 约束：长度大于等于1，小于等于128，且斜杠数量小于等于7
    * 是否必须：Yes
* Qos
    * 描述：消息的QOS等级
    * 类型：Integer
    * 默认值：无
    * 约束：可选值[0, 1]，0-MQTT QOS等级0，1-MQTT QOS等级1
    * 是否必须：Yes
* Content
    * 描述：消息内容
    * 类型：String
    * 默认值：无
    * 约束：长度大于等于1，小于等于1024，Base64编码
    * 是否必须：Yes

```

#### 响应状态码

```
200：正常处理
其它：异常处理

```

#### 响应BODY

```
* Code
    * 描述：错误码
    * 类型: String
    * 是否必须：Yes
* Message
    * 描述：错误详细信息
    * 类型: String
    * 是否必须：Yes
* MessageId
    * 描述：成功发送消息后，云端返回的消息ID，响应状态码等于200的时候才有意义
    * 类型: String
    * 是否必须：Yes

```

### 6.2 查询消息内容

#### 描述

```
查询消息内容，同步

```

#### URL路径

```
/iothub

```

#### 请求方法

```
GET

```

#### 请求头部

```
* X-Product-Id
    * 描述：租户ID
    * 类型：String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes

```

#### QUERY STRING

```
* Action
    * 描述：操作代码
    * 类型：String
    * 默认值：无
    * 约束：固定值-QueryMessage
    * 是否必须：Yes
* Version
    * 描述：版本号
    * 类型：String
    * 默认值：无
    * 约束：固定值-2018-11-01
    * 是否必须：Yes
* ProductKey
    * 描述：产品KEY
    * 类型：String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes
* MessageId
    * 描述：消息ID
    * 类型: String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes

```

#### 请求BODY

```
无

```

#### 响应状态码

```
200：正常处理
其它：异常处理

```

#### 响应BODY

```
* Code
    * 描述：错误码
    * 类型: String
    * 是否必须：Yes
* Message
    * 描述：错误详细信息
    * 类型: String
    * 是否必须：Yes
* Content
    * 描述：消息内容，对应于发布请求中的Content参数，响应状态码等于200的时候才有意义
    * 类型: String
    * 是否必须：Yes

```

### 6.3 查询消息长度配额

#### 描述

```
查询服务端发布消息的最大长度配额，同步

```

#### URL路径

```
/iothub

```

#### 请求方法

```
GET

```

#### 请求头部

```
* X-Product-Id
    * 描述：租户ID
    * 类型：String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes

```

#### QUERY STRING

```
* Action
    * 描述：操作代码
    * 类型：String
    * 默认值：无
    * 约束：固定值-QueryMessageLengthQuota
    * 是否必须：Yes
* Version
    * 描述：版本号
    * 类型：String
    * 默认值：无
    * 约束：固定值-2018-11-01
    * 是否必须：Yes
* ProductKey
    * 描述：产品KEY
    * 类型：String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes

```

#### 请求BODY

```
无

```

#### 响应状态码

```
200：正常处理
其它：异常处理

```

#### 响应BODY

```
* Code
    * 描述：错误码
    * 类型: String
    * 是否必须：Yes
* Message
    * 描述：错误详细信息
    * 类型: String
    * 是否必须：Yes
* Quota
    * 描述：消息最大长度配额，默认值1024，响应状态码等于200的时候才有意义
    * 类型: Integer
    * 是否必须：Yes

```

### 6.4 修改消息长度配额-G1

#### 描述

```
修改具体产品发布消息最大长度配额，同步
G1接口，只面向运营平台

```

#### URL路径

```
/iothub

```

#### 请求方法

```
GET

```

#### 请求头部

```
* X-Origin-Service
    * 描述：调用方服务名
    * 类型：String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes

```

#### QUERY STRING

```
* Action
    * 描述：操作代码
    * 类型：String
    * 默认值：无
    * 约束：固定值-ModifyMessageLengthQuota
    * 是否必须：Yes
* Version
    * 描述：版本号
    * 类型：String
    * 默认值：无
    * 约束：固定值-2018-11-01
    * 是否必须：Yes
* TenantId
    * 描述：租户ID
    * 类型：String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes
* ProductKey
    * 描述：产品KEY
    * 类型：String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes
* Quota
    * 描述：消息最大长度配额
    * 类型：Integer
    * 默认值：无
    * 约束：合法取值区间[0, 4096]
    * 是否必须：Yes

```

#### 请求BODY

```
无

```

#### 响应状态码

```
200：正常处理
其它：异常处理

```

#### 响应BODY

```
* Code
    * 描述：错误码
    * 类型: String
    * 是否必须：Yes
* Message
    * 描述：错误详细信息
    * 类型: String
    * 是否必须：Yes
* Quota
    * 描述：消息最大长度配额，默认值1024，响应状态码等于200的时候才有意义
    * 类型: Integer
    * 是否必须：Yes

```

### 6.5 RRPC

#### 描述

```
服务端RRpc，同步

```

#### URL路径

```
/iothub

```

#### 请求方法

```
POST

```

#### 请求头部

```
* X-Product-Id
    * 描述：租户ID
    * 类型：String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes

```

#### QUERY STRING

```
* Action
    * 描述：操作代码
    * 类型：String
    * 默认值：无
    * 约束：固定值-RRpc
    * 是否必须：Yes
* Version
    * 描述：版本号
    * 类型：String
    * 默认值：无
    * 约束：固定值-2018-11-01
    * 是否必须：Yes

```

#### 请求BODY

```
* ProductKey
    * 描述：产品KEY
    * 类型：String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes
* DeviceName
    * 描述：设备名字
    * 类型: String
    * 默认值：无
    * 约束：长度大于等于1，小于等于63
    * 是否必须：Yes
* RRpcRequest
    * 描述：RRpc请求内容
    * 类型：String
    * 默认值：无
    * 约束：长度大于等于1，小于等于1024，Base64编码
    * 是否必须：Yes
* Timeout
    * 描述：等待设备响应RRPC的时间
    * 类型：Integer
    * 默认值：无
    * 约束：可选值[1, 25000]，单位是毫秒
    * 是否必须：Yes

```

#### 响应状态码

```
200：正常处理
其它：异常处理

```

#### 响应BODY

```
* Code
    * 描述：错误码
    * 类型: String
    * 是否必须：Yes
* Message
    * 描述：错误详细信息
    * 类型: String
    * 是否必须：Yes
* RRpcCode
    * 描述：调用返回码，标识请求状态，取值：SUCCESS--成功，TIMEOUT--设备响应超时，OFFLINE--设备离线，UNKNOW--系统异常，响应状态码等于200的时候才有意义
    * 类型: String
    * 是否必须：Yes
* RRpcResponse
    * 描述：调用成功后，设备的响应内容，Base64编码，在 RRpcCode 为 SUCCESS 时才有意义
    * 类型: String
    * 是否必须：Yes

```

## 7 其他服务依赖的API

### 7.1 查询KAFKA地址-G1

#### 描述

```
查询KAFKA地址，同步
G1接口，仅供规则引擎获取KAFKA地址

```

#### URL路径

```
/iothub

```

#### 请求方法

```
GET

```

#### 请求头部

```
* X-Origin-Service
    * 描述：调用方服务名
    * 类型：String
    * 默认值：无
    * 约束：无
    * 是否必须：Yes

```

#### QUERY STRING

```
* Action
    * 描述：操作代码
    * 类型：String
    * 默认值：无
    * 约束：固定值-GetKafkaAddress
    * 是否必须：Yes
* Version
    * 描述：版本号
    * 类型：String
    * 默认值：无
    * 约束：固定值-2018-11-01
    * 是否必须：Yes

```

#### 请求BODY

```
无

```

#### 响应状态码

```
200：正常处理
其它：异常处理

```

#### 响应BODY

```
* Code
    * 描述：错误码
    * 类型: String
    * 是否必须：Yes
* Message
    * 描述：错误详细信息
    * 类型: String
    * 是否必须：Yes
* KafkaAddress
    * 描述：KAFKA地址，响应状态码等于200的时候才有意义
    * 类型：String
    * 是否必须：Yes

```

## 8  状态码

```
HTTP状态码	Code	Message
200	Success	The api was called successfully.
400	HttpMethodNotSupported	The http method %s is not supported.
400	HttpHeaderRequired	The http header %s is required.
400	HttpHeaderValueError	The value %s of http header %s is error.
400	HttpParameterRequired	The http query parameter %s is required.
400	HttpParameterValueError	The value %s of http query parameter %s is error.
400	OperationNotPermitted	The operation is not permitted, by reason %s.
400	ResourceNotFound	The resource is not found, %s = %s.
400	ResourceAbnormal	The resource is abnormal, %s = %s.
400	ResourceQuotaExhausted	The resource quota is exhausted.

```

## 9 数据结构

### 9.1 ProductInfo

```
* ProductKey
    * 描述：产品KEY，云端分配给用户的随机字符串，唯一标识一个产品
    * 类型: String
    * 是否必须: Yes
* ProductName
    * 描述：产品名称
    * 类型：String
    * 是否必须：Yes
* Description
    * 描述：产品描述，默认为空
    * 类型：String
    * 是否必须：Yes
* DeviceCount
    * 描述：产品关联设备总数
    * 类型：Integer
    * 是否必须：Yes
* AccessPoints
    * 描述：设备接入点地址数组
    * 类型：Array of string
    * 是否必须：Yes
* CreateAt
    * 描述：创建时间，ISO8601（yyyy-MM-dd'T'HH:mm:ss'Z'）
    * 类型：String
    * 是否必须：Yes
* UpdateAt
    * 描述：更新时间，ISO8601（yyyy-MM-dd'T'HH:mm:ss'Z'）
    * 类型：String
    * 是否必须：Yes
```

### 9.2 DeviceInfo

```
* ProductKey
    * 描述：产品Key
    * 类型: String
    * 是否必须: Yes
* DeviceName
    * 描述：设备名字
    * 类型: String
    * 是否必须: Yes
* DeviceSecret
    * 描述：设备密钥，云端分配的随机字符串，一机一密
    * 类型：String
    * 是否必须：Yes
* State
    * 描述：在线状态，Online-在线，Offline-离线
    * 类型：String
    * 是否必须：Yes
* CreateAt
    * 描述：创建时间，ISO8601（yyyy-MM-dd'T'HH:mm:ss'Z'）
    * 类型：String
    * 是否必须：Yes
* UpdateAt
    * 描述：更新时间，ISO8601（yyyy-MM-dd'T'HH:mm:ss'Z'）
    * 类型：String
    * 是否必须：Yes
* LastActiveAt
    * 描述：上次活跃时间，ISO8601（yyyy-MM-dd'T'HH:mm:ss'Z'）
    * 类型：String
    * 是否必须：Yes
```

### 9.3 TopicInfo

```
* TopicId
    * 描述：TOPIC类ID
    * 类型：String
    * 是否必须：Yes
* Topic
    * 描述：具体的TOPIC字符串
    * 类型：String
    * 是否必须：Yes
* TopicType
    * 描述：TOPIC类型，CUSTOM-产品自定义，SYSTEM-系统保留，OTA-空中升级
    * 类型：String
    * 是否必须：Yes
* Operation
    * 描述：设备对TOPIC类的许可操作，SUB-订阅，PUB-发布，ALL-订阅和发布
    * 类型：String
    * 是否必须：Yes
* Qos
    * 描述：MQTT TOPIC QOS等级，0-MQTT QOS等级0，1-MQTT QOS等级1
    * 类型：Integer
    * 是否必须：Yes
* Description
    * 描述：TOPIC类描述
    * 类型：String
    * 是否必须：Yes
* MessageCount
    * 描述：发布消息数
    * 类型：Integer
    * 是否必须：Yes
```


### 9.4 TopicClassInfo

```
* TopicId
    * 描述：TOPIC类ID
    * 类型：String
    * 是否必须：Yes
* TopicName
    * 描述：TOPIC类名
    * 类型：String
    * 是否必须：Yes
* TopicType
    * 描述：TOPIC类型，CUSTOM-产品自定义，SYSTEM-系统保留，OTA-空中升级
    * 类型：String
    * 是否必须：Yes
* Operation
    * 描述：设备对TOPIC类的许可操作，SUB-订阅，PUB-发布，ALL-订阅和发布
    * 类型：String
    * 是否必须：Yes
* Qos
    * 描述：MQTT TOPIC QOS等级，0-MQTT QOS等级0，1-MQTT QOS等级1
    * 类型：Integer
    * 是否必须：Yes
* Description
    * 描述：TOPIC类描述
    * 类型：String
    * 是否必须：Yes
* CreateAt
    * 描述：创建时间，ISO8601（yyyy-MM-dd'T'HH:mm:ss'Z'）
    * 类型：String
    * 是否必须：Yes
* UpdateAt
    * 描述：更新时间，ISO8601（yyyy-MM-dd'T'HH:mm:ss'Z'）
    * 类型：String
    * 是否必须：Yes
```

## 10 修改记录

- DeviceInfo 数据结构增加 ProductKey（2018-12-19）

```
前端设备信息列表需要显示产品名称，这里增加 ProductKey，用于产品查询产品信息得到 ProductName
```

- RegisterDevices 接口增加 ProductKey（2018-12-19）

```
OpenApi 缺少这个字段
```

- Custom 和 System Topic 操作合并（2018-12-20）

- 增加 RRpc 接口

