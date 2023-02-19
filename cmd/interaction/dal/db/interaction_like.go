
package db

import (
	"context"

	"github.com/gitgou/simple_douyin/pkg/constants"
	"log"

	"gorm.io/gorm"
)


//须在数据库表中的点赞记录表中增加点赞记录的ID字段，以便于取消点赞记录
type Favoriate_record_Model struct {
	gorm.Model
	
	ID        int64  `json:"id"`
	VideoID   int64	 `json:"video_id"`
	UserID    string `json:"user_id"`
}

type VideoModel struct {
	gorm.Model
	ID       int64  `json:"id"`
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

	//1.创建点赞记录
	
	if err := DB.WithContext(ctx).Create(favoriate_record).Error; err != nil {
		return err
	}
	return nil

	//2. 根据video_id查询视频记录
	//发现数据库表中没有点赞计数，暂时不往后面写了
	/*res := make([]*VideoModel, 0)
	if err := DB.WithContext(ctx).Where("user_id = ?", userId).Order(`create_at DESC`).Find(&res).Error; err != nil {
		return res, err
	}

	//3. 更新视频的点赞计数

	res.
	
	if err:= DB.Model(&VideoModel).Where("video_id = ?",  )

	return DB.WithContext(ctx).Create(commentModel).Error*/
}


//用户取消点赞(须在前面进行鉴权，user_id == commentModel.user_id)
func DeleteFavoriteAction(ctx context.Context, commentModel *CommentModel) error {
	return DB.WithContext(ctx).Where("id = ?", commentModel.ID).Delete(&CommentModel{})
}

//查询该视频的点赞数
func QueryFavorite(ctx context.Context, int video_id) (int64, error) {
	var count int64
	if err := DB.WithContext(ctx).Model(&Favoriate_record_Model{}).Where("video_id = ?",video_id).Count(&count).Error; err != nil {
		return count, err
	}
	return count, nil
}
//返回查询的喜欢的视频列表(video_id的列表)
func QueryFavoriteVideoList(ctx context.Context, int user_id) ([]VideoModel, error) {
	videoList := []Favoriate_record_Model
	res := []VideoModel
	if err := DB.WithContext(ctx).Model(&Favoriate_record_Model{}).Where("user_id = ?",user_id).Find(&videoList).Error; err != nil {
		return res, err
	}
	for _, v : range videoList {
		var temp VideoModel
		if err := DB.WithContext(ctx).Model(&VideoModel{}).Where("id = ?",v.video_id).Find(&temp).Error; err != nil {
			log.println("this video id is wrong")
		}
		res = append(res, temp)
	}
	return res, nil
}