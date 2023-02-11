package main

import (
	"context"

	"github.com/gitgou/simple_douyin/cmd/video/pack"
	"github.com/gitgou/simple_douyin/cmd/video/service"
	videodemo "github.com/gitgou/simple_douyin/kitex_gen/videodemo"
	"github.com/gitgou/simple_douyin/pkg/errno"
)

// FeedServiceImpl implements the last service interface defined in the IDL.
type FeedServiceImpl struct{}

// Feed implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) Feed(ctx context.Context, req *videodemo.FeedRequest) (resp *videodemo.FeedResponse, err error) {
	resp = new(videodemo.FeedResponse)
	if req.UserID <= 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	videoModels, err := service.NewFeedService(ctx).Feed(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	videos := pack.Videos(videoModels)
	//uIds := pack.UserIds(videoModels)
	//TODO get UserInfo
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.VideoList = videos
	resp.NextTime = 0 // TODO
	return resp, nil
}
