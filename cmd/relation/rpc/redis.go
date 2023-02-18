package rpc

import (
	"context"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/gitgou/simple_douyin/kitex_gen/redisdemo"
	"github.com/gitgou/simple_douyin/kitex_gen/redisdemo/redisservice"
	"github.com/gitgou/simple_douyin/pkg/constants"
	"github.com/gitgou/simple_douyin/pkg/middleware"
	"github.com/golang/glog"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

var redisClient redisservice.Client

func initRedisRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := redisservice.NewClient(
		constants.RedisServiceName,
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
	redisClient = c
}


func ZSetIncr(ctx context.Context, req *redisdemo.ZSETIncreRequest)(error){
	resp, err := redisClient.ZSetIncre(ctx, req)
	if err != nil{
		glog.Error("ZSet Incr err,", err.Error())
		return err 
	}
	if resp.BaseResp.StatusCode != 0 {
		glog.Error("ZSet Incre err,", resp.BaseResp.StatusMsg); 
		return err
	}
	return nil
}