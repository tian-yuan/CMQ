package svc

import (
	"github.com/go-redis/redis"
	"github.com/tian-yuan/iot-common/util"
	"time"
)

var Global = &struct {
	RedisClient    *redis.ClusterClient
	SessionStorage util.SessionStorage
	SessionPrefix  string
	ReqTimeOut time.Duration
	RedisSessionTimeOut time.Duration
	RedisSessionRefresh time.Duration
} {}
