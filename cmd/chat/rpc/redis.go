package rpc

import (
	"context"
	"log"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/gitgou/simple_douyin/kitex_gen/redisdemo"
	"github.com/gitgou/simple_douyin/kitex_gen/redisdemo/redisservice"
	"github.com/gitgou/simple_douyin/pkg/constants"
	"github.com/gitgou/simple_douyin/pkg/errno"
	"github.com/gitgou/simple_douyin/pkg/middleware"
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

func SetMsgId(ctx context.Context, req *redisdemo.SetRequest) (error) {
	resp, err := redisClient.Set(ctx, req)
	if err != nil {
		log.Println("Set Msg Id err,", err.Error())
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}

	return nil
}

func IncreMsgId(ctx context.Context, req *redisdemo.GetIncreIdRequest)(int64){
	resp, err := redisClient.GetIncreId(ctx, req)
	if err != nil{
		log.Println("Set Msg Id err,", err.Error())
		return 0 
	}
	if resp.BaseResp.StatusCode != 0 {
		log.Println("Incre Msg Id err,", resp.BaseResp.StatusMsg); 
		return 0 
	}
	return resp.Id
}