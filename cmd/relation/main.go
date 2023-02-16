package main

import (
	relationdemo "github.com/gitgou/simple_douyin/kitex_gen/relationdemo/relationservice"
	"log"
)

func main() {
	svr := relationdemo.NewServer(new(RelationServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
