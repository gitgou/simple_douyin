package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"

	"github.com/gitgou/simple_douyin/cmd/interaction/dal/db"
	"github.com/gitgou/simple_douyin/cmd/interaction/pack"
	"github.com/gitgou/simple_douyin/cmd/interaction/cache"
	interaction "github.com/gitgou/simple_douyin/kitex_gen/interaction"
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
func (s *InteractionService) FavoriteAction(req *interaction.DouyinFavoriteActionRequest) error {
	//用户鉴权（当前视频是否点赞？）
	flag := db.IsFavoriteVideo(s.ctx, req.User.ID, req.VideoID)
	if 	flag {
		return errno.ParamErr
	}
	return db.FavoriteAction(s.ctx, &db.Favoriate_record_Model{
		VideoID 		: 		req.VideoId
		UserID			:		req.User.Id
	}) 
	
}

//取消点赞
func (s *InteractionService) CancelFavoriteAction(req *interaction.DouyinFavoriteActionRequest) error {
	flag := db.IsFavoriteVideo(s.ctx, req.User.ID, req.VideoID)
	if !flag {
		return errno.ParamErr
	}
	return db.CancelFavoriteAction(s.ctx,&db.Favoriate_record_Model{
		VideoID 		: 		req.VideoId
		UserID			:		req.User.Id
	})
}

//获取点赞的视频列表
func (s *InteractionService) GetFavoriteList(req *interaction.DouyinFavoriteListRequest) []*db.VideoModel, error {
	if req.UserId <= 0 {
		return nil, ParamErr
	}
	return db.QueryFavoriteVideoList(s.ctx, req.UserId)
}

//发布评论 
func (s *InteractionService) PublishComment(req *interaction.DouyinCommentActionRequest) error {
	return db.PublishComment(s.ctx, &db.CommentModel{
		VideoID 		: 	req.VideoId,
		CommentUserID 	:	req.User.Id,
		Content 		: 	req.CommentText,
	})
}

//取消评论
func (s *InteractionService) CancelPublishComment(req *interaction.DouyinCommentActionRequest) error {
	return db.DeleteComment(s.ctx, req.comment_id)
}

//获取视频评论
func (s *InteractionService) GetVideoComments(req *interaction.DouyinCommentListRequest) []*db.CommentModel,error {
	return db.FeedComments(s.ctx, req.video_id)
}