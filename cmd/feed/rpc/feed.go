package rpc

import (
	"context"
	"time"

	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/kitex_gen/userdemo"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/kitex_gen/userdemo/userservice"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/pkg/constants"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/pkg/errno"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/pkg/middleware"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

var userClient userservice.Client

// TODO 确定是不是 user Client
func initUserRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := userservice.NewClient(
		constants.UserServiceName,
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
	userClient = c
}

// MGetUser multiple get list of user info
func MGetUser(ctx context.Context, req *userdemo.MGetUserRequest) (map[int64]*userdemo.User, error) {
	resp, err := userClient.MGetUser(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	res := make(map[int64]*userdemo.User)
	for _, u := range resp.Users {
		res[u.UserId] = u
	}
	return res, nil
}
