package cache

//TODO store user with map
import (

	"github.com/gitgou/simple_douyin/cmd/chat/dal/db"
	//"github.com/gitgou/simple_douyin/kitex_gen/chatdemo"
)

var MapChat map[string][]*db.MessageModel // key : userId

func StoreDB() {
	msgModels := make([]*db.MessageModel, 0, len(MapChat))
	var index = 0
	for _, msgs := range MapChat {
		for _, msg := range msgs{
			msgModels[index] = msg
		}
	}
	//TODO
	/*
	if err := db.UpdateMessage(msgModels); err != nil {
		log.Println("Store DB Fail. ")
	}
	*/
}
