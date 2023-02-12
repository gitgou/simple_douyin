package main

import (
	videodemo "github.com/gitgou/simple_douyin/kitex_gen/videodemo/videoservice"
	"log"
)

func main() {
	svr := videodemo.NewServer(new(VideoServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
