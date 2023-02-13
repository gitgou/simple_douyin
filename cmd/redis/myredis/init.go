package myredis

import (
	"context"
	"fmt"
	"time"
	"github.com/redis/go-redis/v9"
)

var (
	RedisIp    = "127.0.0.1"
	RedisPort  = "6379"
	expireTime = 600
	rdb        *redis.Client
)

func Init() {
	rdb = redis.NewClient(&redis.Options{Addr: RedisIp + ":" + RedisPort, Password: ""})
	test()
}

func test()error{
	_, err := rdb.Set(context.Background(), "user", 1, time.Duration(time.Duration(expireTime).Seconds())).Result()
	if err != nil{
		fmt.Println("init error ")
		return err;
	}
	return nil
}
