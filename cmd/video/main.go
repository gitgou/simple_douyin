package main

import (
	"log"
	videodemo "simple_douyin/kitex_gen/videodemo/feedservice"
)

func main() {
	svr := videodemo.NewServer(new(FeedServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
