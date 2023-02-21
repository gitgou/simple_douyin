package rpc

import (
	"context"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/gitgou/simple_douyin/kitex_gen/relationdemo"
	"github.com/gitgou/simple_douyin/kitex_gen/relationdemo/relationservice"
	"github.com/gitgou/simple_douyin/kitex_gen/userdemo"
	"github.com/gitgou/simple_douyin/pkg/constants"
	"github.com/gitgou/simple_douyin/pkg/errno"
	"github.com/gitgou/simple_douyin/pkg/middleware"
	"github.com/cloudwego/kitex/pkg/klog"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

var relationClient relationservice.Client

func initRelationRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := relationservice.NewClient(
		constants.RelationServiceName,
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
	relationClient = c
}

func Relation(ctx context.Context, req *relationdemo.RelationRequest) error {
	resp, err := relationClient.Relation(ctx, req)
	if err != nil {
		klog.Error("relation err, ", err.Error())
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		klog.Error("relation err, ",resp.BaseResp.StatusMsg) 
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}

	return nil

}

func GetFollowList(ctx context.Context, req *relationdemo.GetFollowRequest)([]*userdemo.User, error){
	resp, err := relationClient.GetFollow(ctx, req)
	if err != nil {
		klog.Error("get follow list err, ", err.Error())
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		klog.Error("get follow err, ",resp.BaseResp.StatusMsg) 
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}

	return resp.UserList, nil
}

func GetFollowerList(ctx context.Context, req *relationdemo.GetFollowerRequest)([]*userdemo.User, error){
	resp, err := relationClient.GetFollower(ctx, req)
	if err != nil {
		klog.Error("get follower err, ", err.Error())
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		klog.Error("get follower err, ",resp.BaseResp.StatusMsg) 
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}

	return resp.UserList, nil

}

func GetFriendList(ctx context.Context, req *relationdemo.GetFriendRequest)([]*relationdemo.FriendUser, error){
	klog.Infof("Get FriendList ")
	resp, err := relationClient.GetFriend(ctx, req)
	klog.Infof("Get FriendList ")
	if err != nil {
		klog.Error("get friend err, ", err.Error())
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		klog.Error("get friend err, ",resp.BaseResp.StatusMsg) 
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}

	return resp.UserList, nil

}