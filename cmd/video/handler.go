package main

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gitgou/simple_douyin/cmd/video/pack"
	"github.com/gitgou/simple_douyin/cmd/video/service"
	videodemo "github.com/gitgou/simple_douyin/kitex_gen/videodemo"
	"github.com/gitgou/simple_douyin/pkg/errno"
)

// FeedServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// Feed implements the FeedServiceImpl interface.
func (s *VideoServiceImpl) Feed(ctx context.Context, req *videodemo.FeedRequest) (resp *videodemo.FeedResponse, err error) {
	resp = new(videodemo.FeedResponse)
	if req.UserID <= 0 {
		klog.Error("Feed Error, ParamErr.", req.UserID)
	}
	
	videoModels, err := service.NewVideoService(ctx).Feed(req)
	if err != nil {
		klog.Error("Feed Error.", err.Error())
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, err
	}
	videos := pack.Videos(videoModels)

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.VideoList = videos
	if len(videoModels) == 0 {
		resp.NextTime = 0
	} else {
		resp.NextTime = videoModels[0].CreatedAt.Unix()
	}
	return resp, nil
}

// Publish implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Publish(ctx context.Context, req *videodemo.PublishRequest) (resp *videodemo.PublishResponse, err error) {
	resp = new(videodemo.PublishResponse)
	if req.UserId <= 0 {
		pack.BuildBaseResp(errno.ParamErr)
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		klog.Error("Publish Error. ParamErr")
		return resp, err
	}

	if err := service.NewVideoService(ctx).Publish(req); err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		klog.Error("Publish Error. ", err.Error())
		return resp, err
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// PublishList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishList(ctx context.Context, req *videodemo.PublishListRequest) (resp *videodemo.PublishListResponse, err error) {
	resp = new(videodemo.PublishListResponse)
	if req.UserId <= 0 {
		pack.BuildBaseResp(errno.ParamErr)
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		klog.Error("PublishList Error. ")
		return resp, err
	}

	videoModels, err := service.NewVideoService(ctx).GetPublishList(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		klog.Error("PublishList Error. ", err.Error())
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	videos := pack.Videos(videoModels)
	resp.Videos = videos
	return resp, nil
}

// GetVideoList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetVideoList(ctx context.Context, req *videodemo.GetVideoListRequest) (resp *videodemo.GetVideoListResponse, err error) {
	// TODO: Your code here...
	return
}
