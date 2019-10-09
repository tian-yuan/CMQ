module github.com/tian-yuan/CMQ/hub

go 1.12

require (
	github.com/go-redis/redis v6.15.6+incompatible
	github.com/gogap/env_json v0.0.0-20150503135429-86150085ddbe // indirect
	github.com/gogap/env_strings v0.0.1 // indirect
	github.com/heralight/logrus_mate v1.0.1
	github.com/hoisie/redis v0.0.0-20160730154456-b5c6e81454e0 // indirect
	github.com/micro/go-micro v1.11.0 // indirect
	github.com/micro/go-plugins v1.3.0 // indirect
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/cobra v0.0.5
	github.com/spf13/viper v1.4.0
	github.com/tian-yuan/iot-common v0.0.1
	golang.org/x/net v0.0.0-20190724013045-ca1201d0de80
	k8s.io/apimachinery v0.0.0-20190727130956-f97a4e5b4abc
)

replace (
	github.com/Sirupsen/logrus v1.0.5 => github.com/sirupsen/logrus v1.0.5
	github.com/Sirupsen/logrus v1.3.0 => github.com/sirupsen/logrus v1.0.6
	github.com/Sirupsen/logrus v1.4.0 => github.com/sirupsen/logrus v1.0.6
	github.com/Sirupsen/logrus v1.4.2 => github.com/sirupsen/logrus v1.4.2
)
