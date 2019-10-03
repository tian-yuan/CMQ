module github.com/tian-yuan/CMQ/controller

go 1.12

require (
	github.com/Sirupsen/logrus v1.4.2 // indirect
	github.com/go-akka/configuration v0.0.0-20190919102339-a31c845c4b1b // indirect
	github.com/go-sql-driver/mysql v1.4.1
	github.com/gogap/config v0.0.0-20190801085516-e664631840ac // indirect
	github.com/gogap/env_json v0.0.0-20150503135429-86150085ddbe // indirect
	github.com/gogap/env_strings v0.0.1 // indirect
	github.com/google/uuid v1.1.1
	github.com/hoisie/redis v0.0.0-20160730154456-b5c6e81454e0 // indirect
	github.com/sirupsen/logrus v1.2.0
	github.com/spf13/cobra v0.0.5
	github.com/spf13/viper v1.4.0
	golang.org/x/net v0.0.0-20190921015927-1a5e07d1ff72
)

replace (
	github.com/Sirupsen/logrus v1.2.0 => github.com/sirupsen/logrus v1.2.0
	github.com/Sirupsen/logrus v1.4.2 => github.com/sirupsen/logrus v1.4.2
)
