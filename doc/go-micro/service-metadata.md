#### service discovery metadata

go-micro 中非常重要的一个模块为服务发现模块，服务发现模块采用的是插件模式，支持 zookeeper、etcd 等
保存到 zookeeper、etcd 中的服务器节点信息如下：

![image-20191125165644288](../images/image-20191125165644288.png)

其中 endpoints 为接口信息，nodes 为节点信息，包括地址等信息

