package myredis

import (
	"github.com/redis/go-redis/v9"
)

var (
	RedisIp    = "127.0.0.1"
	RedisPort  = "6379"
	Rdb        *redis.Client
)

func Init() {
	Rdb = redis.NewClient(&redis.Options{Addr: RedisIp + ":" + RedisPort, Password: ""})
}

