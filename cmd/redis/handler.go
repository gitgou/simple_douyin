package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gitgou/simple_douyin/cmd/redis/myredis"
	redisdemo "github.com/gitgou/simple_douyin/kitex_gen/redisdemo"
	"github.com/gitgou/simple_douyin/pkg/constants"
	"github.com/golang/glog"
)

// RedisServiceImpl implements the last service interface defined in the IDL.
type RedisServiceImpl struct{}

// Set implements the RedisServiceImpl interface.
func (s *RedisServiceImpl) Set(ctx context.Context, req *redisdemo.SetRequest) (resp *redisdemo.SetResponse, err error) {
	resp = new(redisdemo.SetResponse) 
	_, err = myredis.Rdb.Set(ctx, req.Key, req.Value, time.Duration(req.Expire) * time.Second).Result()
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
	v, err := myredis.Rdb.Incr(ctx, req.Key).Result()
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
	_,err = myredis.Rdb.ZIncrBy(ctx, req.Key, float64(req.Increment), req.Member).Result()
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
	v ,err := myredis.Rdb.ZScore(ctx, req.Key, req.Member).Result()
	if err != nil{
		glog.Error("Incre Key error ", err.Error(), ", key: ", req.Key) 
		return resp, err;
	}

	resp.Value = float32(v)
	return resp, nil
}

// GetUserInfo implements the RedisServiceImpl interface.
func (s *RedisServiceImpl) GetUserInfo(ctx context.Context, req *redisdemo.GetUserInfoRequest) (resp *redisdemo.GetUserInfoResponse, err error) {
	resp = new(redisdemo.GetUserInfoResponse)
	userId := req.UserId
	//需要什么信息在下面添加
	//Follow_Count
	followCount, err := myredis.Rdb.ZScore(ctx, constants.RedisZSetKeyFollow, fmt.Sprintf("%x", userId)).Result()
	if err != nil{
		glog.Error("Get Zset Score error ", err.Error(), ", key: ", constants.RedisZSetKeyFollow) 
	}else{
		resp.UserInfo = append(resp.UserInfo, &redisdemo.UserInfo{Key: constants.RedisZSetKeyFollow, Value: float32(followCount)})
	}
	//Follower_Count 
	followerCount, err := myredis.Rdb.ZScore(ctx, constants.RedisZSetKeyFollower, fmt.Sprintf("%x", userId)).Result()
	if err != nil{
		glog.Error("Get Zset Score error ", err.Error(), ", key: ", constants.RedisZSetKeyFollower) 
	}else{
		resp.UserInfo = append(resp.UserInfo, &redisdemo.UserInfo{Key: constants.RedisZSetKeyFollower, Value: float32(followerCount)})
	}

	//Favorite_Count 
	favoriteCount, err := myredis.Rdb.ZScore(ctx, constants.RedisZSetKeyFavorite, fmt.Sprintf("%x", userId)).Result()
	if err != nil{
		glog.Error("Get Zset Score error ", err.Error(), ", key: ", constants.RedisZSetKeyFavorite) 
	}else{
		resp.UserInfo = append(resp.UserInfo, &redisdemo.UserInfo{Key: constants.RedisZSetKeyFollower, Value: float32(favoriteCount)})
	}
	//Favorite_Count 
	favoritedCount, err := myredis.Rdb.ZScore(ctx, constants.RedisZSetKeyFavorited, fmt.Sprintf("%x", userId)).Result()
	if err != nil{
		glog.Error("Get Zset Score error ", err.Error(), ", key: ", constants.RedisZSetKeyFavorited) 
	}else{
		resp.UserInfo = append(resp.UserInfo, &redisdemo.UserInfo{Key: constants.RedisZSetKeyFollower, Value: float32(favoritedCount)})
	}
	
	//Video_Count 
	videoCount, err := myredis.Rdb.ZScore(ctx, constants.RedisZSetKeyVideo, fmt.Sprintf("%x", userId)).Result()
	if err != nil{
		glog.Error("Get Zset Score error ", err.Error(), ", key: ", constants.RedisZSetKeyVideo) 
	}else{
		resp.UserInfo = append(resp.UserInfo, &redisdemo.UserInfo{Key: constants.RedisZSetKeyVideo, Value: float32(videoCount)})
	}
	return resp, nil
}
