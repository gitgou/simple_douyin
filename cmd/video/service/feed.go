package service

import (
	"context"

	"github.com/gitgou/simple_douyin/cmd/video/dal/db"
	"github.com/gitgou/simple_douyin/kitex_gen/demofeed"
)

type FeedService struct {
	ctx context.Context
}

// NewMGetNoteService new FeedService
func NewFeedService(ctx context.Context) *FeedService {
	return &FeedService{ctx: ctx}
}

func (s *FeedService) Feed(rep *demofeed.FeedRequest) ([]*db.VideoModel, error) {
	return db.GetVideo(s.ctx)
}