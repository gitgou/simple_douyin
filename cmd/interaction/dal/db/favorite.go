package db

import (
	"context"

	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gitgou/simple_douyin/pkg/constants"
)

//须在数据库表中的点赞记录表中增加点赞记录的ID字段，以便于取消点赞记录
type FavoriteModel struct {
	ID        	int64  `json:"id"`
	VideoId   	int64	 `json:"video_id"`
	UserId    	int64 `json:"user_id"`
	CreatedAt    time.Time `json:"create_at"`
	UpdatedAt    time.Time `json:"update_at"`
}


func (n *FavoriteModel) TableName() string {
	return constants.FavoriteTableName
}

//用户点赞
func CreateFavorite(ctx context.Context, favorite *FavoriteModel) error {
	//创建点赞记录
	return DB.WithContext(ctx).Create(favorite).Error
}

//查询当前视频用户是否点赞  
func GetFavorite(ctx context.Context, userId int64, videoId int64) (*FavoriteModel) {

	res := make([]*FavoriteModel, 0)

	if err := DB.WithContext(ctx).Where("user_id = ? AND video_id = ?", userId, videoId).Find(&res).Error; err != nil {
		klog.Errorf("get Favorite Err. %s", err.Error())
		return nil
	}
	if len(res) > 0{
		return res[0]
	}else{
		return nil
	}
}


//用户取消点赞
func DeleteFavorite(ctx context.Context, userId int64, videoId int64) error {
	res := make([]*FavoriteModel, 0)
	return DB.WithContext(ctx).Where("user_id = ? AND video_id = ?", userId, videoId).Delete(res).Error
}

func GetFavoriteByUserId(ctx context.Context, userId int64)([]*FavoriteModel, error){
	res := make([]*FavoriteModel, 0)
	if err := DB.WithContext(ctx).Where("user_id = ? ", userId).Find(&res).Error; err != nil{
		return nil, err;
	}
	return res, nil;
}

