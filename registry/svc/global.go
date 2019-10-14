package svc

import (
	"github.com/go-redis/redis"
	"time"
)

var Global = &struct {
	RedisClient    *redis.ClusterClient
	ReqTimeOut time.Duration
	RedisSessionTimeOut time.Duration
	RedisSessionRefresh time.Duration

	DeviceSvc *DeviceSvc
} {}
