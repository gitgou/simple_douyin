package db

import (
	"context"
	"log"

	"github.com/gitgou/simple_douyin/pkg/constants"

	"gorm.io/gorm"
)

type MessageModel struct {
	gorm.Model
	ID        int64  `json:"id"`
	ToUserId      int64 `json:"to_user_id"`
	FromUserId int64 `json:"from_user_id"`
	Content  string `json:"content"`
	CreateTime string `json:"create_time"`
}

func (n *MessageModel) TableName() string {
	return constants.ChatMessageTableName
}

func GetChat(ctx context.Context,userId int64, toUserId int64)([]*MessageModel, error){
	res := make([]*MessageModel, 0)
	if err := DB.WithContext(ctx).Where("to_user_id = ? AND from_user_id = ?", toUserId, userId).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func GetMsgCount()(int64, error){
	var msgCount int64 = 0
	if err := DB.Count(&msgCount).Error; err != nil{
		return 0, err
	}
	return msgCount, nil
}

func InsertMessages(msgModels []*MessageModel)error{
	return DB.WithContext(context.Background()).Create(msgModels).Error
	
}

func GetUserMessages(userId int64)([]*MessageModel, error){
	fromMe := make([]*MessageModel, 0)
	toMe := make([]*MessageModel, 0)
	if err := DB.WithContext(context.Background()).Where("to_user_id = ? ", userId).Find(&toMe).Error; err != nil {
		log.Println("get toMe Msg Fail. userId: ", userId," ", err.Error())	
	}
	if err := DB.WithContext(context.Background()).Where("from_user_id = ? ", userId).Find(&fromMe).Error; err != nil {
		log.Println("get FromMe Msg Fail. userId: ", userId, " ", err.Error())	
	}
	return append(fromMe, toMe...), nil

}