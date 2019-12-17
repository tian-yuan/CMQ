#### 配置中心
本项目固定使用 consul 作为配置中心
不同的服务使用 iot.xxx.xxx.config 方式配置参数
比如:

```
topic-manager 服务
iot.topic.manager.config
{
  "log": {
    "maxSize": 1024,
    "...."
  }
}
```

##### 日志
日志统一使用 zap
配置信息如下：

```
{
  "log": {
    "level": "error",
    "development": false,
    "encoding": "json",
    "logFileDir": "",
    "outputPaths": [],
    "errorOutputPaths": [],
    "maxSize": 50,
    "maxBackups": 200,
    "maxAge": 10
  }
}
```

