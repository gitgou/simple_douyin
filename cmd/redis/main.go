package main

import (
	"net"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/gitgou/simple_douyin/cmd/redis/myredis"

	redisdemo "github.com/gitgou/simple_douyin/kitex_gen/redisdemo/redisservice"
	"github.com/gitgou/simple_douyin/pkg/bound"
	"github.com/gitgou/simple_douyin/pkg/constants"
	"github.com/gitgou/simple_douyin/pkg/middleware"

	tracer2 "github.com/gitgou/simple_douyin/pkg/tracer"

	etcd "github.com/kitex-contrib/registry-etcd"

	trace "github.com/kitex-contrib/tracer-opentracing"
)

func Init() {
	tracer2.InitJaeger(constants.RedisServiceName)
	myredis.Init()
}
func main() {
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", constants.RedisServiceAddress)
	if err != nil {
		panic(err)
	}
	Init()
	svr := redisdemo.NewServer(new(RedisServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.RedisServiceName}), // server name
		server.WithMiddleware(middleware.CommonMiddleware),                                              // middleware
		server.WithMiddleware(middleware.ServerMiddleware),
		server.WithServiceAddr(addr),                                       // address
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), // limit
		server.WithMuxTransport(),                                          // Multiplex
		server.WithSuite(trace.NewDefaultServerSuite()),                    // tracer
		server.WithBoundHandler(bound.NewCpuLimitHandler()),                // BoundHandler
		server.WithRegistry(r),                                             // registry
	)
	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
