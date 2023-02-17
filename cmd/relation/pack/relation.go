package pack

import (
	"github.com/gitgou/simple_douyin/kitex_gen/relationdemo"
	"github.com/gitgou/simple_douyin/kitex_gen/userdemo"
)


func Friends(requestUserId int64, userList []*userdemo.User)[]*relationdemo.FriendUser{

	if userList == nil || len(userList) == 0 {
		return nil
	}
	friendList := make([]*relationdemo.FriendUser, 0, len(userList))
	for i , u := range userList {
		friendList[i].User = u
		friendList[i].Message = "" // TODO 获取最新消息
		friendList[i].MsgType = 1 // TODO 
	}
	return friendList
}
