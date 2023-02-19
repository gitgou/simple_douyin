package cache

//TODO store user with map
import (
	"context"
	"fmt"
	"sync"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gitgou/simple_douyin/cmd/chat/dal/db"
	"github.com/gitgou/simple_douyin/cmd/chat/rpc"
	"github.com/gitgou/simple_douyin/cmd/chat/utils"
	"github.com/gitgou/simple_douyin/cmd/user/cache"
	"github.com/gitgou/simple_douyin/kitex_gen/redisdemo"
	"github.com/gitgou/simple_douyin/kitex_gen/userdemo"
	"github.com/gitgou/simple_douyin/pkg/constants"
)

var (
	//Store Old Msg, Already store to db
	MapChat map[string][]*db.MessageModel // key :  from_user_id & to_user_id
	//new generate msg, not store to db
	MapNewChat map[string][]*db.MessageModel
	MutexChat sync.Mutex
)

func initMsgSequenceId() {
	MessageSequenceId, err := db.GetMsgCount()
	if err != nil {
		klog.Fatal("Get Msg Count Fail, %s", err.Error())
	}
	fmt.Println("MsgID: ", MessageSequenceId)
	err = rpc.SetMsgId(context.Background(), &redisdemo.SetRequest{Key: constants.ChatMsgIdKey,
		Value: MessageSequenceId, Expire: 0})
	if err != nil {
		klog.Fatal("Set Msg Id Fail")
	}
}

func StoreDB() {
	MutexChat.Lock()
	defer MutexChat.Unlock()
	msgModels := make([]*db.MessageModel, 0)
	for key, msgs := range MapNewChat {
		isOnline := rpc.CheckUserOnline(context.Background(),
			&userdemo.CheckUserOnlineRequest{UserIds: utils.SpliceChatKey(key)})
		//store to db
		for _, msg := range msgs {
			msgModels = append(msgModels, msg)
			// new chat msg to old
			if isOnline {
				MapChat[key] = append(MapChat[key], msg)
			}
		}

		//delete old chat data if user not online
		if !isOnline {
			delete(MapChat, key)
		}

		delete(MapNewChat, key)
	}

	//store db
	if len(msgModels) != 0 {
		if err := db.InsertMessages(msgModels); err != nil {
			klog.Errorf("Store DB Fail. %s", err.Error())
		}
	}
	// clear old msg(stored to db) of user not online
	for key := range MapChat {
		isOnline := rpc.CheckUserOnline(context.Background(),
			&userdemo.CheckUserOnlineRequest{UserIds: utils.SpliceChatKey(key)})
		if !isOnline {
			delete(MapChat, key)
		}
	}
}

func Login(userId int64) error {
	//拉用户聊天数据
	msgModels, err := db.GetUserMessages(userId)
	if err != nil {
		klog.Errorf("login | get user msg fail. userId: %d", userId)
		return err
	}

	cache.MutexUser.Lock()
	defer cache.MutexUser.Unlock()
	for _, msg := range msgModels {
		chatKey := utils.GenChatKey(msg.ToUserId, msg.FromUserId)
		if _, exist := MapChat[chatKey]; exist {
			continue
		}
		MapChat[chatKey] = append(MapChat[chatKey], msg)
	}
	return nil
}
