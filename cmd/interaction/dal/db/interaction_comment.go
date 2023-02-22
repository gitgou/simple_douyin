
package db

import (
	"context"

	"github.com/gitgou/simple_douyin/pkg/constants"
	interaction "github.com/gitgou/simple_douyin/kitex_gen/interactiondemo"

	"gorm.io/gorm"
)

func (n *UserModel) TableName() string {
	return constants.CommentTableName
}


type CommentModel struct {
	gorm.Model
	ID        		int64   `json:"id"`
	VideoID   		int64	`json:"video_id"`
	CommentUserID   string  `json:"commentuser_id"`
	Content 		string 	`json:"content"`
}
type UserModel struct {
	ID              int64     `json:"id"`
	Name            string    `json:"name"`
	AvatarUrl       string    `json:"avatar_url"`
	Password        string    `json:"password"`
	BackgroundImage string    `json:"background_image"`
	Signature       string    `json:"signature"`
	CreateAt        time.Time `json:"create_at"`
	DeleteAt        time.Time `json:"delete_at"`
	UpdateAt        time.Time `json:"update_at"`
}

type FollowModel struct {
	//ID int64 `json:"id"`
	FollowId   int64  `json:"follow_id"`
	FollowerId int64  `json:"follower_id"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time 	`json:"update_at"`
}

//用户发布评论
func PublishComment(ctx context.Context, commentModel *CommentModel) *CommentModel, error {

	err := DB.WithContext(ctx).Create(commentModel).Error
	if err != nil {
		return nil, err
	} else {
		return &commentModel, nil
	}
}


//用户删除评论(须在前面进行鉴权，user_id == commentModel.user_id)
func DeleteComment(ctx context.Context, comment_id int64) error {
	return DB.WithContext(ctx).Where("id = ?", comment_id).Delete(&CommentModel{})
}

//推送跟视频相关的全部评论
func FeedComments(ctx context.Context, video_id int64) ([]*CommentModel, error) {
	res := make([]*CommentModel, 0)
	if err := DB.WithContext(ctx).Where("video_id = ?", video_id).Order(`create_at DESC`).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

//视频评论计数
func CountComments(ctx context.Context, int video_id) (int 64, error) {
	var count int64
	if err := DB.WithContext(ctx).Model(&CommentModel{}).Where("video_id = ?",video_id).Count(&count).Error; err != nil {
		return count, err
	}
	return count, nil
}

//根据用户id获取用户信息 
func GetInfoByUserID(ctx context.Context,  user_id int64) (user *UserModel, error) {
	var user *UserModel 
	if err := DB.WithContext(ctx).Where("user_id = ?", user_id).Find(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func CheckIsFollow(ctx context.Context, userid int64, follower_id int64) (bool, error){
	record := FollowModel{}

	if err := DB.WithContext(ctx).Where("follow_id = ? AND follower_id = ?", userid, follower_id).Find(&record).Error; err != nil {
		return false, err
	}
	if record == nil {
		return false, nil
	} else {
		return true, nil
	}
}

func GetFollowAndFollowerCount(ctx context.Context, userid int64) (int64, int64, error) {
	var FollowCount, FollowerCount int64
	if err := DB.WithContext(ctx).Model(&FollowModel{}).Where("follower_id = ?",userid).Count(&FollowCount).Error; err != nil {
		return FollowCount, FollowerCount, err
	}
	if err := DB.WithContext(ctx).Model(&FollowModel{}).Where("follower_id = ?",userid).Count(&FollowerCount).Error; err != nil {
		return FollowCount, FollowerCount, err
	}
	return FollowCount, FollowerCount, err
}