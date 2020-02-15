#### CMQ

CMQ 是一个开源的分布式 MQTT 消息平台，支持海量物联网设备一站式接入、MQTT&CoAP 多协议处理、低时延实时消息通信。通过开发自定义的规则引擎，可以将数据转发到其他消息中间件，比如 kafka 等；

整个平台使用 GO 开发，采用的开源的 go-micro 框架，支持微服务部署，理论上支持无限的水平扩容能力。

#### 架构



![image-20200215204323882](../../../../../Library/Application Support/typora-user-images/image-20200215204323882.png)

#### 组件

LoadBalance: 负载均衡，由外部提供的四层负载均衡，可以使用 LVS、HAPROXY 等四层负载均衡。

Hub: 设备链路网关，负责设备链路管理。

Registry: 设备验证服务，具体请参见设备管理文档。

TopicManager: Topic 管理服务，负责设备 Topic 订阅管理，设备订阅 Topic 落盘和加载。

TopicAcl: Topic 权限管理服务。

PublishEngin: 消息发布引擎。

PublishAcl: 消息发布 acl。

MessageDispatch: 消息分发组件。

#### 安装使用

使用 https://github.com/mattn/goreman 多进程管理工具启动服务，运行 ./bin/start.sh 即可