package main

import (
	"context"

	"github.com/gitgou/simple_douyin/cmd/relation/pack"
	"github.com/gitgou/simple_douyin/cmd/relation/rpc"
	"github.com/gitgou/simple_douyin/cmd/relation/service"
	relationdemo "github.com/gitgou/simple_douyin/kitex_gen/relationdemo"
	"github.com/gitgou/simple_douyin/kitex_gen/userdemo"
	"github.com/gitgou/simple_douyin/pkg/errno"
	"github.com/golang/glog"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

// Relation implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) Relation(ctx context.Context, req *relationdemo.RelationRequest) (resp *relationdemo.RelationResponse, err error) {
	resp = new(relationdemo.RelationResponse)

	if err := service.NewRelationService(ctx).Relation(req); err != nil {
		glog.Error("Relation error. ", err.Error());
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// GetFollow implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFollow(ctx context.Context, req *relationdemo.GetFollowRequest) (resp *relationdemo.GetFollowResponse, err error) {
	resp = new(relationdemo.GetFollowResponse)

	followList := service.NewRelationService(ctx).GetFollowList(req.UserId)
	if followList == nil {
		resp.BaseResp = pack.BuildBaseResp(errno.Success)
		return resp, nil
	}

	followUserIds := service.GetFollowUserIds(followList)
	//Get Follower User Info
	userList, err := rpc.MGetUser(ctx, &userdemo.MGetUserRequest{
		UserIds:       followUserIds,
		RequestUserId: req.UserId})
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		glog.Error("Get FollowList GetUser error. ", err.Error());
		return resp, nil
	}

	resp.UserList = userList
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// GetFollower implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFollower(ctx context.Context, req *relationdemo.GetFollowerRequest) (resp *relationdemo.GetFollowerResponse, err error) {
	resp = new(relationdemo.GetFollowerResponse)

	followerList := service.NewRelationService(ctx).GetFollowerList(req.UserId)
	if followerList == nil {
		resp.BaseResp = pack.BuildBaseResp(errno.Success)
		return resp, nil
	}

	followerUserIds := service.GetFollowUserIds(followerList)
	//Get Follow User Info 获取粉丝用户信息
	userList, err := rpc.MGetUser(ctx, &userdemo.MGetUserRequest{
		UserIds:       followerUserIds,
		RequestUserId: req.UserId})
	if err != nil {
		glog.Error("Get FollowerList GetUser error. ", err.Error());
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.UserList = userList
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// GetFriend implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFriend(ctx context.Context, req *relationdemo.GetFriendRequest) (resp *relationdemo.GetFriendResponse, err error) {
	resp = new(relationdemo.GetFriendResponse)

	friendList := service.NewRelationService(ctx).GetFriendList(req.UserId)
	if friendList == nil {
		resp.BaseResp = pack.BuildBaseResp(errno.Success)
		return resp, nil
	}

	friendUserIds := service.GetFriendUserIds(req.UserId, friendList)
	//Get Follow User Info 获取粉丝用户信息
	userList, err := rpc.MGetUser(ctx, &userdemo.MGetUserRequest{
		UserIds:       friendUserIds,
		RequestUserId: req.UserId})
	if err != nil {
		glog.Error("Get FriendList GetUser error. ", err.Error());
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.UserList = pack.Friends(req.UserId, userList)
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// CheckFollowRelation implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) CheckFollowRelation(ctx context.Context, req *relationdemo.CheckFollowRelationRequest) (resp *relationdemo.CheckFollowRelationResponse, err error) {
	resp = new(relationdemo.CheckFollowRelationResponse)
	resp.Follow = service.NewRelationService(ctx).CheckFollowRelation(req)
	return resp, nil
}