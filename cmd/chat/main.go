package main

import (
	"log"

	"github.com/gitgou/simple_douyin/cmd/chat/cache"
	"github.com/gitgou/simple_douyin/cmd/chat/dal"
	"github.com/gitgou/simple_douyin/cmd/chat/rpc"
	"github.com/gitgou/simple_douyin/cmd/chat/cache/ticker"
	chatdemo "github.com/gitgou/simple_douyin/kitex_gen/chatdemo/chatservice"
)
func Init(){
	dal.Init()
	cache.Init()
	rpc.Init()
	ticker.Init()
}
func main() {
	Init()
	svr := chatdemo.NewServer(new(ChatServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
