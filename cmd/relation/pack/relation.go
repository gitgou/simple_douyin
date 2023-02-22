package pack

import (
	"context"

	"github.com/gitgou/simple_douyin/cmd/relation/rpc"
	"github.com/gitgou/simple_douyin/kitex_gen/chatdemo"
	"github.com/gitgou/simple_douyin/kitex_gen/relationdemo"
	"github.com/gitgou/simple_douyin/kitex_gen/userdemo"
)


func Friends(requestUserId int64, userList []*userdemo.User)[]*relationdemo.FriendUser{

	if userList == nil || len(userList) == 0 {
		return nil
	}
	friendList := make([]*relationdemo.FriendUser, 0, len(userList))
	for _ , u := range userList {
		newMsg, newMsgType := rpc.GetNewChat(context.Background(), &chatdemo.GetNewMsgRequest{UserId:requestUserId, ToUserId: u.Id}) 
		friendList = append(friendList,  &relationdemo.FriendUser{
			Id: u.Id,
			Name: u.Name,
			IsFollow: u.IsFollow,
			FollowCount: u.FollowCount,
			FollowerCount: u.FollowerCount,
			FavoriteCount: u.FavoriteCount,
			WorkCount: u.WorkCount,
			TotalFavorited: u.TotalFavorited,
			BackgroundImage: u.BackgroundImage,
			Signature: u.Signature,
			Avatar: u.Avatar,
			Message: newMsg,
			MsgType: newMsgType,})
	}
	return friendList
}
