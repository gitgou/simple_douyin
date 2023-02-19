package db

import (
	"github.com/gitgou/simple_douyin/pkg/constants"
	"github.com/henrylee2cn/ameda/test/time"

)

type FriendModel struct {
	//ID int64 `json:"id"`
	PriFriendId   int64  `json:"primary_friend_id"`
	SecFriendId int64  `json:"second_friend_id"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time 	`json:"update_at"`
}


func (n *FriendModel) TableName() string {
	return constants.FriendTableName
}

func GetFriendList(userId int64)[]*FriendModel{
	var FriendModels = make([]*FriendModel, 0)
	DB.Where("primary_friend_id = ? OR second_friend_id = ?", userId, userId).Find(&FriendModels)
	
	return FriendModels
}