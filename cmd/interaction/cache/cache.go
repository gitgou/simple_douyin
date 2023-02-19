package cache

//TODO store user with map
import (
	"log"
	"github.com/gitgou/simple_douyin/cmd/interaction/dal/db"
	"context"

)

type ActionRecord struct {
	record db.Favoriate_record_Model
	action int32          //action = 1, 点赞行为， action = 2， 取消行为
}


//缓存点赞数

var MapFavoriteCount map[int64]int64 // key : videoId, value : favorite_model
var SliceRecord []ActionRecord

func AppendRecord(video_id int64, user_id int64, action int32) {
	var record ActionRecord = ActionRecord {
		record : db.Favoriate_record_Model{
			VideoID : video_id,
			UserID 	: user_id
		}
		action : action
	}
	SliceRecord = append(SliceRecord, record)
}
func Favorite(video_id int64, user_id int64){

	if _, ok := range MapFavoriteCount; ok {
		MapLoginUser[video_id] = MapLoginUser[video.ID] + 1
	} else {
		count, err := db.QueryFavorite(context.Context, video_id)
		if err == nil {
			MapUser[video_id] = 1
		} else {
			MapUser[video_id] = count
		}
	}
	AppendRecord(video_id, user_id, 1)	 	
}

func CancelFavorite(video_id int64, user_id int64){

	if _, ok := range MapFavoriteCount; ok {
		MapLoginUser[video_id] = MapLoginUser[video.ID] - 1
	} else {
		count, err := db.QueryFavorite(context.Context, video_id)
		if err == nil {
			return 
		}
		MapUser[video_id] = count - 1
	}
	AppendRecord(video_id, user_id, 2)	 	 	
}

func StoreDB(){
	for _, v := range SliceRecord {
		if(v.action == 1) {
			db.create(&v.record)
		}
		else {
			db.delete(&v.record)
		}
	}
	if err := db.UpdateUsers(userModels); err != nil{
		log.Println("Store DB Fail. ");
	}
}

