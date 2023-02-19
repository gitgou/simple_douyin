package main

import (
	"log"
	interaction "simple_douyin-master/cmd/interaction/kitex_gen/interaction/interaction"
)

func main() {
	svr := interaction.NewServer(new(InteractionImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
