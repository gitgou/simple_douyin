package rpc

import (
	"context"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/gitgou/simple_douyin/kitex_gen/interactiondemo"
	"github.com/gitgou/simple_douyin/kitex_gen/interactiondemo/interactionservice"
	"github.com/gitgou/simple_douyin/kitex_gen/videodemo"
	"github.com/gitgou/simple_douyin/pkg/constants"
	"github.com/gitgou/simple_douyin/pkg/errno"
	"github.com/gitgou/simple_douyin/pkg/middleware"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

var interactionClient interactionservice.Client

func initInteractionRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := interactionservice.NewClient(
		constants.InteractionServiceName,
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r),                            // resolver
	)
	if err != nil {
		panic(err)
	}
	interactionClient = c
}

func FavoriteAction(ctx context.Context, req *interactiondemo.FavoriteRequest) error {
	resp, err := interactionClient.FavoriteAction(ctx, req)
	if err != nil {
		klog.Error("favoriteaction err, ", err.Error())
		return err
	}
	if resp.Resp.StatusCode != 0 {
		klog.Error("favoriteaction err, ", resp.Resp.StatusMsg)
		return errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMsg)
	}

	return nil

}

func GetFavoriteList(ctx context.Context, req *interactiondemo.GetFavoriteListRequest) ([]*videodemo.Video, error) {
	resp, err := interactionClient.FavoriteList(ctx, req)
	if err != nil {
		klog.Error("favorite list err, ", err.Error())
		return nil, err
	}
	if resp.Resp.StatusCode != 0 {
		klog.Error("favorite list err, ", resp.Resp.StatusMsg)
		return nil, errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMsg)
	}

	return resp.VideoList, nil

}

func CommentAction(ctx context.Context, req *interactiondemo.CommentRequest) (*interactiondemo.Comment, error) {
	resp, err := interactionClient.CommentAction(ctx, req)
	if err != nil {
		klog.Error("comment action err, ", err.Error())
		return nil, err
	}
	if resp.Resp.StatusCode != 0 {
		klog.Error("comment action err, ", resp.Resp.StatusMsg)
		return nil, errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMsg)
	}

	return resp.Comment, nil

}

func GetCommentList(ctx context.Context, req *interactiondemo.GetCommentListRequest) ([]*interactiondemo.Comment, error) {
	resp, err := interactionClient.CommentList(ctx, req)
	if err != nil {
		klog.Error("comment list err, ", err.Error())
		return nil, err
	}
	if resp.Resp.StatusCode != 0 {
		klog.Error("comment action err, ", resp.Resp.StatusMsg)
		return nil, errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMsg)
	}

	return resp.CommentList, nil

}
