module github.com/tian-yuan/CMQ/message-dispatcher

go 1.12

require (
	github.com/go-akka/configuration v0.0.0-20190919102339-a31c845c4b1b // indirect
	github.com/go-redis/redis v6.15.6+incompatible // indirect
	github.com/gogap/config v0.0.0-20190801085516-e664631840ac // indirect
	github.com/gogap/env_json v0.0.0-20150503135429-86150085ddbe // indirect
	github.com/gogap/env_strings v0.0.1 // indirect
	github.com/gogap/logrus_mate v1.1.0 // indirect
	github.com/heralight/logrus_mate v1.0.1
	github.com/hoisie/redis v0.0.0-20160730154456-b5c6e81454e0 // indirect
	github.com/micro/go-micro v1.11.0
	github.com/micro/go-plugins v1.3.0
	github.com/natefinch/lumberjack v2.0.0+incompatible // indirect
	github.com/samuel/go-zookeeper v0.0.0-20180130194729-c4fab1ac1bec
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/cobra v0.0.5
	github.com/spf13/viper v1.4.0
	github.com/tian-yuan/iot-common v0.0.4
	golang.org/x/net v0.0.0-20190724013045-ca1201d0de80
)

replace (
	github.com/Sirupsen/logrus v1.0.5 => github.com/sirupsen/logrus v1.0.5
	github.com/Sirupsen/logrus v1.3.0 => github.com/sirupsen/logrus v1.0.6
	github.com/Sirupsen/logrus v1.4.0 => github.com/sirupsen/logrus v1.0.6
	github.com/Sirupsen/logrus v1.4.2 => github.com/sirupsen/logrus v1.4.2
)
