package main

import (
	"log"

	"github.com/gitgou/simple_douyin/cmd/video/dal/db"
	userdemo "github.com/gitgou/simple_douyin/kitex_gen/userdemo/userservice"
)

func main() {
	db.Init()
	
	svr := userdemo.NewServer(new(UserServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
