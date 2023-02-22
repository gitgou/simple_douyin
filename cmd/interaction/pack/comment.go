package pack

import (
	"context"
	"github.com/gitgou/simple_douyin/cmd/interaction/dal/db"
	"github.com/gitgou/simple_douyin/kitex_gen/interaction"
	"log"
)

func Comment(comment *db.CommentModel, userLook int64) *interaction.Comment {
	if comment == nil {
		return nil
	}
	user := db.GetInfoByUserID(context.Background(), comment.CommentUserID)
	isFollow, err := db.CheckIsFollow(context.Background(), comment.CommentUserID, userLook)
	if err != nil {
		log.println("db check Isfollow error")
	}
	FollowCount, FollowerCount, err := db.GetFollowAndFollowerCount(context.Background(), comment.CommentUserID)
	user_response := &interaction.User {
		Id		 	 	: user.ID,
		Name 	  		: user.Name,
		FollowCount		: FollowCount
		FollowerCount	: FollowerCount
		IsFollow 		: is_follow,
	} 
	return &interaction.Comment{
		Id				:   comment.ID,
		User			: 	user_response,
		content 		: 	comment.content,
		create_date		: 	comment.gorm.Model.CreatedAt,
	}
}

func Comments(comments []*db.CommentModel, userLook int64) []*interaction.comment {
	if comments == nil || len(comments) == 0 {
		return nil
	}
	res := make([]*interaction.comment, 0)
	for _, m := range comments {
		if n := Comment(m, userLook); n != nil {
			res = append(res, n)
		}
	}
	return res
}



