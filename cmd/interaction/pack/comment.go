package pack

import (
	"github.com/gitgou/simple_douyin/cmd/interaction/dal/db"
	"github.com/gitgou/simple_douyin/kitex_gen/interaction"
)

func Comment(comment *db.CommentModel) *interaction.Comment {
	if m == nil {
		return nil
	}
	user := db.GetInfoByUserID(comment.UserID)

	user_response := interaction.User {
		id		  : user.ID,
		name 	  : user.Name,
		//关注暂时没写
		is_follow : false,
	} 
	return &interaction.Comment{
		Id		:   m.ID,
		User	: 	user_response,
		content : comment.content,
		create_date: comment.gorm.Model.CreatedAt,
	}
}

func Comments(comments []*db.CommentModel) []*interaction.comment {
	if comments == nil || len(comments) == 0 {
		return nil
	}
	res := make([]*interaction.comment, 0)
	for _, m := range comments {
		if n := Comment(m); n != nil {
			res = append(res, n)
		}
	}
	return res
}

func Video(video *db.VideoModel) (*interaction.video) {
	if video ==nil {
		return nil
	}
	var favoritecount int64
	if err := db.DB.Model(&VideoModel{}).Where("video_id = ?",video_id).Count(&favoritecount).Error; err != nil {
		return nil
	}

	var commentcount int64
	if err := db.DB.Model(&commentModel{}).Where("video_id = ?",video_id).Count(&commentcount).Error; err != nil {
		return nil
	}

	Author := db.GetInfoByUserID(comment.UserID)
	user_response := interaction.User {
		id		  : Author.ID,
		name 	  : Author.Name,
		//关注暂时没写
		is_follow : false,
	}


	return &interaction.video{
		Id				:   video.ID,
		author			: 	Author,
		play_url 		: 	video.PlayURL,
		cover_url		:	video.CoverURL,
		favorite_count	:	favoritecount,
		comment_count	:	commentcount,
		is_favorite		:	true
		title			: 	video.Title
	}
}

func Videos(videos []*db.VideoModel) ([] *interaction.video) {
	if videos == nil || len(videos) == 0 {
		return nil
	}
	res := make([]*interaction.video, 0)
	for _, m := range videos {
		if n := Video(m); n != nil {
			res = append(res, n)
		}
	}
	return res
}

