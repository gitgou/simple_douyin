package main

import (
	videodemo "github.com/gitgou/simple_douyin/kitex_gen/videodemo/videoservice"
	"github.com/gitgou/simple_douyin/cmd/video/dal"
	"log"
)

func main() {
	dal.Init()
	svr := videodemo.NewServer(new(VideoServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
