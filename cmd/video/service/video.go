package service

import (
	"context"
	"fmt"
	"time"

	"github.com/gitgou/simple_douyin/cmd/relation/rpc"
	"github.com/gitgou/simple_douyin/cmd/video/dal/db"
	"github.com/gitgou/simple_douyin/kitex_gen/redisdemo"
	"github.com/gitgou/simple_douyin/kitex_gen/videodemo"
	"github.com/gitgou/simple_douyin/pkg/constants"
)

type VideoService struct {
	ctx context.Context
}

// NewMGetNoteService new VideoService
func NewVideoService(ctx context.Context) *VideoService {
	return &VideoService{ctx: ctx}
}

func (s *VideoService) Feed(req *videodemo.FeedRequest) ([]*db.VideoModel, error) {
	if req.LatestTime == 0 {
		req.LatestTime = time.Now().Unix()
	}
	return db.FeedVideo(s.ctx, time.Unix(req.LatestTime, 0))
}

func (s *VideoService) Publish(req * videodemo.PublishRequest)(error){
	if err := db.PublishVideo(s.ctx, & db.VideoModel{
		UserID: req.UserId,
		PlayURL: req.Url,
		Title: req.Title,
	}); err != nil{
		return err;
	}
	return rpc.ZSetIncr(s.ctx, &redisdemo.ZSETIncreRequest{Key: constants.RedisZSetKeyVideo,
						Member: fmt.Sprintf("%x", req.UserId),
						Increment: 1,})
}

func(s *VideoService) GetPublishList(req * videodemo.PublishListRequest)([]*db.VideoModel, error){
	return db.GetVideos(s.ctx, req.UserId)
}