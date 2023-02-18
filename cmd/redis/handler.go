package main

import (
	"context"
	"log"
	"time"

	"github.com/gitgou/simple_douyin/cmd/redis/myredis"
	redisdemo "github.com/gitgou/simple_douyin/kitex_gen/redisdemo"
	"github.com/golang/glog"
)

// RedisServiceImpl implements the last service interface defined in the IDL.
type RedisServiceImpl struct{}

// Set implements the RedisServiceImpl interface.
func (s *RedisServiceImpl) Set(ctx context.Context, req *redisdemo.SetRequest) (resp *redisdemo.SetResponse, err error) {
	resp = new(redisdemo.SetResponse) 
	_, err = myredis.Rdb.Set(context.Background(), req.Key, req.Value, time.Duration(req.Expire) * time.Second).Result()
	if err != nil{
		log.Println("Set Key error ", err.Error(), ", key: ", req.Key, ", Value: ", req.Value)
		resp.BaseResp = &redisdemo.BaseResp{StatusCode: 1, StatusMsg: err.Error()}
		return resp, err;
	}

	resp.BaseResp = &redisdemo.BaseResp{StatusCode: 0} 
	return resp, nil
}

// GetIncreId implements the RedisServiceImpl interface.
func (s *RedisServiceImpl) GetIncreId(ctx context.Context, req *redisdemo.GetIncreIdRequest) (resp *redisdemo.GetIncreIdResponse, err error) {
	resp = new(redisdemo.GetIncreIdResponse)
	v, err := myredis.Rdb.Incr(context.Background(), req.Key).Result()
	if err != nil{
		log.Println("Incre Key error ", err.Error(), ", key: ", req.Key) 
		resp.BaseResp = &redisdemo.BaseResp{StatusCode: 1, StatusMsg: err.Error()}
		return resp, err;
	}

	resp.BaseResp = &redisdemo.BaseResp{StatusCode: 0} 
	resp.Id = v 
	return resp, nil
}


// ZSetIncre implements the RedisServiceImpl interface.
func (s *RedisServiceImpl) ZSetIncre(ctx context.Context, req *redisdemo.ZSETIncreRequest) (resp *redisdemo.ZSETIncreResponse, err error) {
	resp = new(redisdemo.ZSETIncreResponse)
	_,err = myredis.Rdb.ZIncrBy(context.Background(), req.Key, float64(req.Increment), req.Menber).Result()
	if err != nil{
		glog.Error("Incre Key error ", err.Error(), ", key: ", req.Key) 
		resp.BaseResp = &redisdemo.BaseResp{StatusCode: 1, StatusMsg: err.Error()}
		return resp, err;
	}

	resp.BaseResp = &redisdemo.BaseResp{StatusCode: 0} 
	return resp, nil
}


// ZSetGetMember implements the RedisServiceImpl interface.
func (s *RedisServiceImpl) ZSetGetMember(ctx context.Context, req *redisdemo.ZSETGetMemberRequest) (resp *redisdemo.ZSETGetMemberResponse, err error) {
	resp = new(redisdemo.ZSETGetMemberResponse)
	v ,err := myredis.Rdb.ZScore(context.Background(), req.Key, req.Menber).Result()
	if err != nil{
		glog.Error("Incre Key error ", err.Error(), ", key: ", req.Key) 
		return resp, err;
	}

	resp.Value = float32(v)
	return resp, nil
}
