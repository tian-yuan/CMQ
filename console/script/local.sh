export ENV=test
export NODE_CONFIG_PLUGIN=env
export NODE_CAMEL_CASE=true

export WEB__PROTOCOL=http
export WEB__PORT=18080
export WEB__TRUST_PROXY=true

export SERVICES__SERVER_DISCOVERY=consul
export SERVICES__CONSUL__NODES__0__HOST=localhost
export SERVICES__CONSUL__NODES__0__PORT=8500
export SERVICES__DEVICE_SERVICE_NAME=push-stat

export SERVICES__SERVER_DISCOVERY=services
export SERVICES__SERVICES_INFO=push-stat:localhost:9091

node ./app.js
