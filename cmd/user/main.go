package main

import (
	"log"

	"github.com/gitgou/simple_douyin/cmd/chat/cache"
	"github.com/gitgou/simple_douyin/cmd/chat/cache/ticker"
	"github.com/gitgou/simple_douyin/cmd/video/dal"
	userdemo "github.com/gitgou/simple_douyin/kitex_gen/userdemo/userservice"
)
func Init(){
	ticker.Init()
	dal.Init()
	cache.Init()	
}
func main() {
	Init()
	svr := userdemo.NewServer(new(UserServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
