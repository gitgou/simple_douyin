package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"

	"github.com/gitgou/simple_douyin/cmd/interaction/dal/db"
	"github.com/gitgou/simple_douyin/cmd/interaction/pack"
	"github.com/gitgou/simple_douyin/cmd/interaction/cache"
	"github.com/gitgou/simple_douyin/kitex_gen/interaction"
	"github.com/gitgou/simple_douyin/pkg/errno"
)

type InteractionService struct {
	ctx context.Context
}

// NewMGetNoteService new FeedService
func NewInteractionService(ctx context.Context) *InteractionService {
	return &InteractionService{ctx: ctx}
}

//点赞
func (s *UserService) FavoriteAction(req *interaction.DouyinFavoriteActionRequest) error {
	//缓存点赞数
	return cache.Favorite(req.video_id,req.user.id)
}

//取消点赞
func (s *UserService) CancelFavoriteAction(req *interaction.DouyinFavoriteActionRequest) error {
	return cache.CancelFavorite(req.video_id,req.user.id)
}

//获取点赞的列表
func (s *UserService) PublishComment(req *interaction.DouyinFavoriteListRequest) error {
	return db.PublishComment(s.ctx, CommentModel{
		UserID 	: req.user.id,
		VideoID : req.video_id,
		Content : req.comment_text,
	})
}

//发布评论 

func (s *UserService) PublishComment(req *interaction.DouyinCommentActionRequest) error {
	return db.QueryFavoriteVideoList(s.ctx, req.user.id)
}

//取消评论
func (s *UserService) CancelPublishComment(req *interaction.DouyinCommentActionRequest) error {
	return db.DeleteComment(s.ctx, req.comment_id)
}

//获取视频评论
func (s *UserService) GetVideoComments(req *interaction.DouyinCommentListRequest) []*db.CommentModel,error {
	return db.FeedComments(s.ctx, req.video_id)
}