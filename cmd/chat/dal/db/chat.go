
package db

import (
	"context"

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