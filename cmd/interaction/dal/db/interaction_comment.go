
package db

import (
	"context"

	"github.com/gitgou/simple_douyin/pkg/constants"

	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	AvatarUrl string `json:"avatar_url"`
	Password  string `json:"password"`
}

type CommentModel struct {
	gorm.Model
	ID        int64  `json:"id"`
	UserID    string `json:"user_id"`
	VideoID   int64	 `json:"video_id"`
	Content   string `json:"content"`
}

func (n *UserModel) TableName() string {
	return constants.CommentTableName
}

//用户发布评论
func PublishComment(ctx context.Context, commentModel *CommentModel) error {
	return DB.WithContext(ctx).Create(commentModel).Error
}


//用户删除评论(须在前面进行鉴权，user_id == commentModel.user_id)
func DeleteComment(ctx context.Context, comment_id int64) error {
	return DB.WithContext(ctx).Where("id = ?", comment_id).Delete(&CommentModel{})
}

//推送跟视频相关的全部评论
func FeedComments(ctx context.Context, int video_id) ([]*CommentModel, error) {
	res := make([]*CommentModel, 0)
	if err := DB.WithContext(ctx).Where("video_id = ?", video_id).Order(`create_at DESC`).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

func CountComments(ctx context.Context, int video_id) (int 64, error) {
	var count int64
	if err := DB.WithContext(ctx).Model(&commentModel{}).Where("video_id = ?",video_id).Count(&count).Error; err != nil {
		return count, err
	}
	return count, nil
}

//根据用户id获取用户信息 

func GetInfoByUserID(ctx context.Context, int user_id) (user *UserModel, error) {
	var user *UserModel 
	if err := DB.WithContext(ctx).Where("user_id = ?", user_id).Find(user).Error; err != nil {
		return user, err
	}
	return user, nil
}