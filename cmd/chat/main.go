package main

import (
	chatdemo "github.com/gitgou/simple_douyin/kitex_gen/chatdemo/chatservice"
	"log"
)

func main() {
	svr := chatdemo.NewServer(new(ChatServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
