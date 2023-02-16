package main

import (
	"log"

	"github.com/gitgou/simple_douyin/cmd/video/rpc"
	"github.com/gitgou/simple_douyin/cmd/video/dal"
	videodemo "github.com/gitgou/simple_douyin/kitex_gen/videodemo/videoservice"
)
func Init(){
	dal.Init()
	rpc.InitRPC()
}
func main() {
	dal.Init()
	svr := videodemo.NewServer(new(VideoServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
