package service

import (
	"context"
	"fmt"

	"github.com/gitgou/simple_douyin/cmd/relation/dal/db"
	"github.com/gitgou/simple_douyin/cmd/relation/rpc"
	"github.com/gitgou/simple_douyin/kitex_gen/redisdemo"
	"github.com/gitgou/simple_douyin/kitex_gen/relationdemo"
	"github.com/gitgou/simple_douyin/pkg/constants"
	"github.com/gitgou/simple_douyin/pkg/errno"
)

type RelationService struct {
	ctx context.Context
}

// NewRelationService new RelationService
func NewRelationService(ctx context.Context) *RelationService {
	return &RelationService{ctx: ctx}
}

func (s *RelationService) Relation(req *relationdemo.RelationRequest) error {
	if req.ActionType == int64(relationdemo.RelationActionType_ACTION_FOLLOW) {
		return s.follow(req.UserId, req.ToUserId)
	} else if req.ActionType == int64(relationdemo.RelationActionType_ACTION_CANCEL_FOLLOW) {
		return s.cancelFollow(req.UserId, req.ToUserId)
	}

	return errno.ParamErr
}

func (s *RelationService)follow(userId int64, toUserId int64) error {
	followModel := db.GetFollowRelation(userId, toUserId)
	if followModel != nil {
		return errno.UserIsAlreadyFollowErr
	}
	if err := db.CreateFollow(userId, toUserId); err != nil{
		return err;
	}

	//add follow_count & follower_count
	rpc.ZSetIncr(s.ctx, &redisdemo.ZSETIncreRequest{Key: constants.RedisZSetKeyFollow,
		 Member: fmt.Sprintf("%x", userId), Increment: 1,}) 
	rpc.ZSetIncr(s.ctx, &redisdemo.ZSETIncreRequest{Key: constants.RedisZSetKeyFollower,
		 Member: fmt.Sprintf("%x", toUserId), Increment: 1,})
	return nil
}

func (s *RelationService)cancelFollow(userId int64, toUserId int64) error {
	if err := db.DeleteFollow(userId, toUserId); err != nil{
		return err
	}
	//reduce follow_count & follower_count
	rpc.ZSetIncr(s.ctx, &redisdemo.ZSETIncreRequest{Key: constants.RedisZSetKeyFollow,
		 Member: fmt.Sprintf("%x", userId), Increment: -1,}) 
	rpc.ZSetIncr(s.ctx, &redisdemo.ZSETIncreRequest{Key: constants.RedisZSetKeyFollower,
		 Member: fmt.Sprintf("%x", toUserId), Increment: -1,})
	return nil
}

func (s *RelationService) GetFollowList(userId int64) []*db.FollowModel {
	return db.GetFollowList(userId)
}

func (s *RelationService) GetFollowerList(userId int64) []*db.FollowModel {
	return db.GetFollowerList(userId)
}

func (s *RelationService) GetFriendList(userId int64) []*db.FriendModel {
	return db.GetFriendList(userId)
}

func (s *RelationService)CheckFollowRelation(req * relationdemo.CheckFollowRelationRequest)bool{
	if req.RelationType == int64(relationdemo.CheckFollowRelationType_FOLLOW){
		return db.GetFollowRelation(req.UserId, req.ToUserId) != nil 
	}else{
		return db.GetFollowRelation(req.ToUserId, req.UserId) != nil 
	}
}

func GetFollowUserIds(followList []*db.FollowModel) []int64 {
	var userIds = make([]int64, 0, len(followList))
	for i, f := range followList {
		userIds[i] = f.FollowerId
	}
	return userIds
}
//requestUserId 是请求好友用户的 USER_ID
func GetFriendUserIds(requestUserId int64, friendList []*db.FriendModel) []int64 {
	var userIds = make([]int64, 0, len(friendList))
	for i, f := range friendList {
		if f.PriFriendId != requestUserId {
			userIds[i] = f.PriFriendId
		}else{
			userIds[i] = f.SecFriendId
		}
	}
	return userIds

}
