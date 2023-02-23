package db

import (
	"context"
	"time"

	"github.com/gitgou/simple_douyin/pkg/constants"
)

type CommentModel struct {
	ID        int64     `json:"id"`
	VideoId   int64     `json:"video_id"`
	UserId    int64 	`json:"user_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (n *CommentModel) TableName() string {
	return constants.CommentTableName
}

// 用户发布评论
func CreateComment(ctx context.Context, commentModel *CommentModel) (*CommentModel, error) {
	err := DB.WithContext(ctx).Create(commentModel).Error
	if err != nil {
		return nil, err
	} else {
		return commentModel, nil
	}
}

// 用户删除评论
func DeleteComment(ctx context.Context, userId int64, commentId int64) error {
	return DB.WithContext(ctx).Where("user_id = ? AND id = ?", userId, commentId).Delete(&CommentModel{}).Error
}

// 推送跟视频相关的全部评论
func GetCommentList(ctx context.Context, videoId int64) ([]*CommentModel,error) {
	res := make([]*CommentModel, 0)
	if err := DB.WithContext(ctx).Where("video_id = ?", videoId).Order(`created_at DESC`).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
