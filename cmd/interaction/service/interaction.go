package service

import (
	"context"
	"fmt"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gitgou/simple_douyin/cmd/interaction/dal/db"
	"github.com/gitgou/simple_douyin/cmd/interaction/rpc"
	"github.com/gitgou/simple_douyin/kitex_gen/interactiondemo"
	"github.com/gitgou/simple_douyin/kitex_gen/redisdemo"
	"github.com/gitgou/simple_douyin/kitex_gen/videodemo"
	"github.com/gitgou/simple_douyin/pkg/constants"
	"github.com/gitgou/simple_douyin/pkg/errno"
)

type InteractionService struct {
	ctx context.Context
}

// NewMGetNoteService new FeedService
func NewInteractionService(ctx context.Context) *InteractionService {
	return &InteractionService{ctx: ctx}
}

// 点赞
func (s *InteractionService) FavoriteAction(req *interactiondemo.FavoriteRequest) error {
	if req.ActionType == int32(interactiondemo.FavoriteActionType_FAVORITE_ACTION_FAVORITE) {
		return s.favorite(req.UserId, req.VideoId)
	} else {
		return s.cancelFavorite(req.UserId, req.VideoId)
	}

}
func (s *InteractionService) favorite(userId int64, videoId int64) error {
	favorite := db.GetFavorite(s.ctx, userId, videoId)
	if favorite != nil {
		klog.Errorf("is already favorite. %s")
		return errno.ParamErr
	}
	videoList := rpc.GetVideoList(context.Background(), &videodemo.GetVideoListRequest{VideoId: []int64{videoId}})
	if len(videoList) < 1{
		klog.Errorf("get video not exist. userId:%d, videoId:%d.", userId, videoId)
		return errno.ServiceErr
	}

	if err := db.CreateFavorite(s.ctx, &db.FavoriteModel{
		VideoId: videoId,
		UserId:  userId,
	}); err != nil{
		return err;
	}

	//add follow_count & follower_count
	rpc.ZSetIncr(s.ctx, &redisdemo.ZSETIncreRequest{Key: constants.RedisZSetKeyFavorite,
		Member: fmt.Sprintf("%x", userId), Increment: 1})
	rpc.ZSetIncr(s.ctx, &redisdemo.ZSETIncreRequest{Key: constants.RedisZSetKeyFavorited,
		Member: fmt.Sprintf("%x", videoList[0].Author.Id), Increment: 1})
	return nil

}

// 取消点赞
func (s *InteractionService) cancelFavorite(userId int64, videoId int64) error {
	favorite := db.GetFavorite(s.ctx, userId, videoId)
	if favorite == nil {
		return errno.ParamErr
	}
	if err := db.DeleteFavorite(s.ctx, userId, videoId);err != nil{
		return err;
	}
	
	videoList := rpc.GetVideoList(context.Background(), &videodemo.GetVideoListRequest{VideoId: []int64{videoId}})
	if len(videoList) < 1{
		klog.Error("get video not exist. userId:%d, videoId:%d.", userId, videoId)
		return errno.ServiceErr
	}

	rpc.ZSetIncr(s.ctx, &redisdemo.ZSETIncreRequest{Key: constants.RedisZSetKeyFavorite,
		Member: fmt.Sprintf("%x", userId), Increment: - 1})
	rpc.ZSetIncr(s.ctx, &redisdemo.ZSETIncreRequest{Key: constants.RedisZSetKeyFavorited,
		Member: fmt.Sprintf("%x", videoList[0].Author.Id), Increment: - 1})
	return nil
}

// 获取点赞的视频列表
func (s *InteractionService) GetFavoriteList(ctx context.Context, userId int64) ([]*db.FavoriteModel, error) {
	return db.GetFavoriteByUserId(s.ctx, userId)
}

func (s *InteractionService) Comment(ctx context.Context, req *interactiondemo.CommentRequest) (*db.CommentModel, error) {
	if req.ActionType == int32(interactiondemo.COMMENTActionType_COMMENT_ACTION_COMMENT) {
		return s.publishComment(req.UserId, req.VideoId, req.CommentText)
	} else {
		err := s.cancelComment(req.UserId, req.CommentId)
		return nil, err
	}
}

// 发布评论
func (s *InteractionService) publishComment(userId int64, videoId int64, commentText string) (*db.CommentModel, error) {
	return db.CreateComment(s.ctx, &db.CommentModel{
		VideoId: videoId,
		UserId:  userId,
		Content: commentText,
	})
}

// 取消评论
func (s *InteractionService) cancelComment(userId int64, commentId int64) error {
	return db.DeleteComment(s.ctx, userId, commentId)
}

// 获取视频评论
func (s *InteractionService) GetCommentList(ctx context.Context, videoId int64) ([]*db.CommentModel, error) {
	return db.GetCommentList(s.ctx, videoId)
}
