
package db

import (
	"context"

	"github.com/gitgou/simple_douyin/pkg/constants"
	"log"
	interaction "github.com/gitgou/simple_douyin/kitex_gen/interactiondemo"
	"gorm.io/gorm"
	"time"
)


//须在数据库表中的点赞记录表中增加点赞记录的ID字段，以便于取消点赞记录
type Favoriate_record_Model struct {
	ID        	int64  `json:"id"`
	VideoID   	int64	 `json:"video_id"`
	UserID    	string `json:"user_id"`
	CreateAt    time.Time `json:"create_at"`
	UpdateAt    time.Time `json:"update_at"`
}

type VideoModel struct {
	gorm.Model
	ID		 int64  `json:"id"`
	UserID   int64  `json:"user_id"`
	Title    string `json:"title"`
	PlayURL  string `json:"play_url"`
	CoverURL string `json:"cover_url"`
}

func (n *Favoriate_record_Model) TableName() string {
	return constants.FavoriateRecordTableName
}

//用户点赞
func FavoriteAction(ctx context.Context, favoriate_record *Favoriate_record_Model) error {

	//创建点赞记录
	
	if err := DB.WithContext(ctx).Create(favoriate_record).Error; err != nil {
		return err
	}
	return nil
}

//查询当前视频用户是否点赞  
func IsFavoriteVideo(ctx context.Context, user_id int64, video_id int64) (bool, error) {

	record := Favoriate_record_Model{}

	if err := DB.WithContext(ctx).Where("user_id = ? AND video_id = ?", user_id, video).Find(&record).Error; err != nil {
		return false, err
	}

	if record == nil {
		return false, nil
	} else {
		return true, nil
	}
}


//用户取消点赞
func DeleteFavoriteAction(ctx context.Context, favoriate_record *Favoriate_record_Model) error {
	return DB.WithContext(ctx).Where("user_id = ? AND video_id = ?", favoriate_record.UserID, favoriate_record.VideoID).Delete(&Favoriate_record_Model{})
}

//查询该视频的点赞数(喜欢计数)
func QueryFavorite(ctx context.Context, int video_id) (int64, error) {
	var count int64
	if err := DB.WithContext(ctx).Model(&Favoriate_record_Model{}).Where("video_id = ?",video_id).Count(&count).Error; err != nil {
		return count, err
	}
	return count, nil
}
//返回查询的喜欢的视频列表(video_id的列表)
func QueryFavoriteVideoList(ctx context.Context, int user_id) ([]*VideoModel, error) {
	videoIdList := []*Favoriate_record_Model
	res := []*VideoModel
	if err := DB.WithContext(ctx).Model(&Favoriate_record_Model{}).Where("user_id = ?",user_id).Find(&videoIdList).Error; err != nil {
		return res, err
	}
	for _, v : range videoIdList {
		var temp *VideoModel
		if err := DB.WithContext(ctx).Model(&VideoModel{}).Where("id = ?",v.VideoID).Find(&temp).Error; err != nil {
			log.println("this videoid is wrong")
			return nil, err
		}
		res = append(res, temp)
	}
	return res, nil
}



