package cache

//TODO store user with map
import (
	"context"
	"log"

	"github.com/gitgou/simple_douyin/cmd/chat/dal/db"
	"github.com/gitgou/simple_douyin/cmd/chat/rpc"
	"github.com/gitgou/simple_douyin/cmd/chat/utils"
	"github.com/gitgou/simple_douyin/kitex_gen/redisdemo"
	"github.com/gitgou/simple_douyin/kitex_gen/userdemo"
	"github.com/gitgou/simple_douyin/pkg/constants"
	//"github.com/gitgou/simple_douyin/kitex_gen/chatdemo"
)

var (
	//Store Old Msg, Already store to db
	MapChat map[string][]*db.MessageModel // key :  from_user_id & to_user_id
	//new generate msg, not store to db 
	MapNewChat map[string][]*db.MessageModel
)

func initMsgSequenceId(){
	MessageSequenceId, err := db.GetMsgCount();
	if err != nil{
		log.Panic("Get Msg Count Fail", err.Error())
	}
	err = rpc.SetMsgId(context.Background(), &redisdemo.SetRequest{Key: constants.ChatMsgIdKey, 
		Value: MessageSequenceId, Expire: 0,})
	if err != nil{
		log.Panic("Set Msg Id Fail")
	}
}
// TODO 有问题
func StoreDB() {
	msgModels := make([]*db.MessageModel, 0) 
	for key, msgs := range MapNewChat {
		isOnline := rpc.CheckUserOnline(context.Background(), 
		 &userdemo.CheckUserOnlineRequest{UserIds: utils.SpliceChatKey(key)})
		//store to db
		for _, msg := range msgs{
			msgModels = append(msgModels, msg)
			// new chat msg to old
			if isOnline {
				MapChat[key] = append(MapChat[key], msg)
			}
		}

		//delete old chat data if user not online
		if ! isOnline {
			 	delete(MapChat, key)
		}

		delete(MapNewChat, key)
	}

	//store db
	if err := db.InsertMessages(msgModels); err != nil {
		log.Println("Store DB Fail. ", err.Error())
	}
	
	// clear old msg(stored to db) of user not online
	for key := range MapChat {
		isOnline := rpc.CheckUserOnline(context.Background(), 
		 &userdemo.CheckUserOnlineRequest{UserIds: utils.SpliceChatKey(key)})
		if ! isOnline{
			delete(MapChat, key)
		}
	}
}

func Login(userId int64)error{
	msgModels, err := db.GetUserMessages(userId)
	if err != nil {
		log.Println("login | get user msg fail. userId: ", userId);
		return err
	}

	for _, msg := range msgModels {
		chatKey := utils.GenChatKey(msg.ToUserId, msg.FromUserId)
		if _, exist := MapChat[chatKey]; exist{
			continue
		}
		MapChat[chatKey] = append(MapChat[chatKey], msg)
	}
	return nil
}
