package db

import (
	"context"
	"log"
	"time"

	"github.com/gitgou/simple_douyin/pkg/constants"
	"github.com/gitgou/simple_douyin/pkg/errno"
)

type MessageModel struct {
	ID         int64     `json:"id"`
	ToUserId   int64     `json:"to_user_id"`
	FromUserId int64     `json:"from_user_id"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"created_at"`
}

func (n *MessageModel) TableName() string {
	return constants.ChatMessageTableName
}

func GetChat(ctx context.Context, userId int64, toUserId int64, preMsgTime time.Time) ([]*MessageModel, error) {
	res := make([]*MessageModel, 0)
	if err := DB.WithContext(ctx).Where("to_user_id IN ? AND from_user_id IN ? AND created_at > ?", []int64{toUserId, userId}, []int64{toUserId, userId}, preMsgTime).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func GetNewMsg(ctx  context.Context,userId int64, toUserId int64) *MessageModel {
	res := make([]*MessageModel, 0)
	if err := DB.WithContext(ctx).Where("to_user_id IN ? AND from_user_id IN ? ", []int64{toUserId, userId}, []int64{toUserId, userId}).Order("created_at DESC").Limit(1).Find(&res).Error; err != nil {
		return nil
	}
	if len(res) > 0{
		return res[0]
	}else{
		return nil
	}
}
func GetMsgCount() (int64, error) {
	var msgCount int64 = 0
	if err := DB.Table(constants.ChatMessageTableName).Count(&msgCount).Error; err != nil {
		return 0, err
	}
	return msgCount, nil
}

func InsertMessages(msgModels []*MessageModel) error {
	if msgModels == nil || len(msgModels) == 0 {
		return errno.Success
	}
	return DB.WithContext(context.Background()).Create(msgModels).Error

}

func GetUserMessages(userId int64) ([]*MessageModel, error) {
	msgs := make([]*MessageModel, 0)
	//升序
	if err := DB.WithContext(context.Background()).Where("to_user_id = ? OR from_user_id = ?", userId, userId).Order("created_at ASC").Find(&msgs).Error; err != nil {
		log.Println("get toMe Msg Fail. userId: ", userId, " ", err.Error())
	}
	return msgs, nil

}
