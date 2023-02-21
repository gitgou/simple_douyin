package db

import (
	"github.com/gitgou/simple_douyin/pkg/constants"
	"time"

)

type FollowModel struct {
	//ID int64 `json:"id"`
	FollowId   int64  `json:"follow_id"`
	FollowerId int64  `json:"follower_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time 	`json:"updated_at"`
}


func (n *FollowModel) TableName() string {
	return constants.FollowTableName
}
// userId : 粉丝Id , toUserId : 被关注者Id 
func GetFollowRelation(userId int64, toUserId int64)*FollowModel{
	var res []*FollowModel
	if err := DB.Where("follow_id = ? AND follower_id = ?", userId, toUserId).Find(&res).Error; err != nil{
		return nil;
	}
	if len(res) > 0 {
		return res[0]
	}else{
		return nil
	}
}
//关注列表
func GetFollowList(userId int64)[]*FollowModel{
	var res []*FollowModel
	if err := DB.Where("follow_id = ? ", userId).Find(&res).Error; err != nil{
		return nil;
	}
	return res

}
//粉丝列表
func GetFollowerList(userId int64)[]*FollowModel{
	var res []*FollowModel
	if err := DB.Where("follower_id = ? ", userId).Find(&res).Error; err != nil{
		return nil;
	}
	return res

}

func CreateFollow(userId int64, toUserId int64)(error){
	return DB.Create(&FollowModel{
		FollowId: userId,
		FollowerId: toUserId,
	}).Error
}

func DeleteFollow(userId int64, toUserId int64)(error){
	return DB.Where("follow_id = ? AND follower_id = ?", userId, toUserId).Delete(&FollowModel{}).Error
}