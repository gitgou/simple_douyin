package main

import (
	"github.com/gitgou/simple_douyin/cmd/redis/myredis"
	redisdemo "github.com/gitgou/simple_douyin/kitex_gen/redisdemo/redisservice"
	"log"
)

func main() {
	myredis.Init();	
	
	svr := redisdemo.NewServer(new(RedisServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
