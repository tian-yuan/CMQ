module github.com/tian-yuan/CMQ/business-manager

go 1.12

require (
	github.com/eclipse/paho.mqtt.golang v1.2.0
	github.com/go-redis/redis v6.15.6+incompatible // indirect
	github.com/gogap/env_json v0.0.0-20150503135429-86150085ddbe // indirect
	github.com/gogap/env_strings v0.0.1 // indirect
	github.com/heralight/logrus_mate v1.0.1 // indirect
	github.com/hoisie/redis v0.0.0-20160730154456-b5c6e81454e0 // indirect
	github.com/micro/go-micro v1.11.0
	github.com/micro/go-plugins v1.3.0
	github.com/natefinch/lumberjack v2.0.0+incompatible // indirect
	github.com/opentracing/opentracing-go v1.1.0
	github.com/spf13/cobra v0.0.5
	github.com/tian-yuan/iot-common v0.0.7
	github.com/uber/jaeger-client-go v2.22.1+incompatible // indirect
	github.com/uber/jaeger-lib v2.2.0+incompatible // indirect
	golang.org/x/net v0.0.0-20200114155413-6afb5195e5aa
)

replace (
	github.com/Sirupsen/logrus v1.0.5 => github.com/sirupsen/logrus v1.0.5
	github.com/Sirupsen/logrus v1.3.0 => github.com/sirupsen/logrus v1.0.6
	github.com/Sirupsen/logrus v1.4.0 => github.com/sirupsen/logrus v1.0.6
	github.com/Sirupsen/logrus v1.4.2 => github.com/sirupsen/logrus v1.4.2
)
