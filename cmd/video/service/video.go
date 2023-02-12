package service

import (
	"context"

	"github.com/gitgou/simple_douyin/cmd/video/dal/db"
	"github.com/gitgou/simple_douyin/kitex_gen/videodemo"
)

type VideoService struct {
	ctx context.Context
}

// NewMGetNoteService new VideoService
func NewVideoService(ctx context.Context) *VideoService {
	return &VideoService{ctx: ctx}
}

func (s *VideoService) Feed(rep *videodemo.FeedRequest) ([]*db.VideoModel, error) {
	return db.FeedVideo(s.ctx)
}

func (s *VideoService) Publish(req * videodemo.PublishRequest)(error){
	return db.PublishVideo(s.ctx, & db.VideoModel{
		UserID: req.UserId,
		PlayURL: req.Url,
		Title: req.Title,
	})
}

func(s *VideoService) GetPublishList(req * videodemo.PublishListRequest)([]*db.VideoModel, error){

}