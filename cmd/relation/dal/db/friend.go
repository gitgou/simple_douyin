package db

import (
	"github.com/gitgou/simple_douyin/pkg/constants"
	"time"
)

type FriendModel struct {
	//ID int64 `json:"id"`
	PrimaryFriendId int64     `json:"primary_friend_id"`
	SecondFriendId int64     `json:"second_friend_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (n *FriendModel) TableName() string {
	return constants.FriendTableName
}

func GetFriendList(userId int64) []*FriendModel {
	var FriendModels = make([]*FriendModel, 0)
	DB.Where("primary_friend_id = ? OR second_friend_id = ?", userId, userId).Find(&FriendModels)

	return FriendModels
}
