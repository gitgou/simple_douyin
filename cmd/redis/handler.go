package main

import (
	"context"
	"log"
	"time"
	redisdemo "github.com/gitgou/simple_douyin/kitex_gen/redisdemo"
	"github.com/gitgou/simple_douyin/cmd/redis/myredis"
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
