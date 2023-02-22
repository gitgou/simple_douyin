package pack

import (
	"context"
	"github.com/gitgou/simple_douyin/cmd/interaction/dal/db"
	"github.com/gitgou/simple_douyin/kitex_gen/interaction"
	"log"
)


/* type VideoModel struct {
	gorm.Model
	ID
	UserID   int64  `json:"user_id"`
	Title    string `json:"title"`
	PlayURL  string `json:"play_url"`
	CoverURL string `json:"cover_url"`
} */


/* type Video struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                            // 视频唯一标识
	Author        *User  `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`                                     // 视频作者信息
	PlayUrl       string `protobuf:"bytes,3,opt,name=play_url,json=playUrl,proto3" json:"play_url,omitempty"`                    // 视频播放地址
	CoverUrl      string `protobuf:"bytes,4,opt,name=cover_url,json=coverUrl,proto3" json:"cover_url,omitempty"`                 // 视频封面地址
	FavoriteCount int64  `protobuf:"varint,5,opt,name=favorite_count,json=favoriteCount,proto3" json:"favorite_count,omitempty"` // 视频的点赞总数
	CommentCount  int64  `protobuf:"varint,6,opt,name=comment_count,json=commentCount,proto3" json:"comment_count,omitempty"`    // 视频的评论总数
	IsFavorite    bool   `protobuf:"varint,7,opt,name=is_favorite,json=isFavorite,proto3" json:"is_favorite,omitempty"`          // true-已点赞，false-未点赞
	Title         string `protobuf:"bytes,8,opt,name=title,proto3" json:"title,omitempty"`                                       // 视频标题
} */



func Video(video *db.VideoModel, userLook int64) (*interaction.Video, error) {
	if video ==nil {
		return nil
	}
	var FavoriteVideoCount int64
	FavoriteVideoCount, err := db.QueryFavorite(context.Background(), video.ID)
	if err != nil {
		return nil, err
	}
	var IsFavorite bool 
	IsFavorite, err := db.IsFavoriteVideo(context.Background(), video.ID, userLook)
	if err != nil {
		return nil, err
	}
	var Commentcount int64
	Commentcount, err := CountComments(context.Background(), video.ID)
	if err != nil {
		return nil, err
	}
	Author := db.GetInfoByUserID(video.UserID)
	isFollow, err := db.CheckIsFollow(context.Background(), video.UserID, userLook)
	if err != nil {
		log.println("db check Isfollow error")
		return nil, err
	}
	FollowCount, FollowerCount, err := db.GetFollowAndFollowerCount(context.Background(), video.UserID)
	if err != nil {
		log.println("db check FollowAndFollowerCount error")
		return nil, err
	}
	user_response := &interaction.User {
		Id		 	 	: Author.ID,
		Name 	  		: Author.Name,
		FollowCount		: FollowCount
		FollowerCount	: FollowerCount
		IsFollow 		: is_follow,
	} 
	return &interaction.video{
		Id				:   video.ID,
		Author			: 	user_response,
		PlayUrl 		: 	video.PlayURL,
		CoverUrl		:	video.CoverURL,
		FavoriteCount	:	FavoriteVideoCount,
		CommentCount	:	Commentcount,
		IsFavorite		:	IsFavorite,
		Title			: 	video.Title,
	}
}

func Videos(videos []*db.VideoModel, userLook int64) ([]*interaction.Video) {
	if videos == nil || len(videos) == 0 {
		return nil
	}
	res := make([]*interaction.Video, 0)
	for _, m := range videos {
		if n := Video(m, userLook); n != nil {
			res = append(res, n)
		}
	}
	return res
}