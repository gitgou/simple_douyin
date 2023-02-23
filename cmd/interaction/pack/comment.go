package pack

import (
	"context"

	"github.com/gitgou/simple_douyin/cmd/interaction/rpc"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gitgou/simple_douyin/cmd/interaction/dal/db"
	"github.com/gitgou/simple_douyin/kitex_gen/interactiondemo"
	"github.com/gitgou/simple_douyin/kitex_gen/userdemo"
)

func Comment(comment *db.CommentModel, user *userdemo.User) *interactiondemo.Comment {
	if comment == nil || user == nil{
		klog.Errorf("get user err. userId: %d", comment.UserId)
		return nil
	}

	return &interactiondemo.Comment{
		Id: comment.ID,
		Content: comment.Content,
		CreateDate: comment.CreatedAt.String(),
		User: user,	
	}
}

func Comments(comments []*db.CommentModel) []*interactiondemo.Comment {
	if comments == nil || len(comments) == 0 {
		return nil
	}
	res := make([]*interactiondemo.Comment, 0, len(comments))
	userIds := GetCommentUserIds(comments)
	users, _ := rpc.MGetUser(context.Background(), &userdemo.MGetUserRequest{UserIds: userIds})
	for _, m := range comments {
		if n := Comment(m, users[m.UserId]); n != nil {
			res = append(res, n)
		}
	}
	return res
}

func GetCommentUserIds(comments []*db.CommentModel)([]int64){
	if comments == nil || len(comments) == 0 {
		return nil
	}

	userIds := make([]int64, 0, len(comments))
	mapUserIds := make(map[int64]struct{}, 0)
	for _, comment := range comments{
		if _, exist := mapUserIds[comment.UserId]; !exist{
			userIds = append(userIds, comment.UserId)
			mapUserIds[comment.UserId] = struct{}{}
		}
	}
	return userIds
}

