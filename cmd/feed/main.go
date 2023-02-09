package main

import (
	"log"
	demofeed "simple_douyin/kitex_gen/demofeed/feedservice"
)

func main() {
	svr := demofeed.NewServer(new(FeedServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
