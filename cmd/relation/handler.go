package main

import (
	"context"
	relationdemo "github.com/gitgou/simple_douyin/kitex_gen/relationdemo"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

// Relation implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) Relation(ctx context.Context, req *relationdemo.RelationRequest) (resp *relationdemo.RelationResponse, err error) {
	// TODO: Your code here...
	return
}

// GetFollow implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFollow(ctx context.Context, req *relationdemo.GetFollowRequest) (resp *relationdemo.GetFollowResponse, err error) {
	// TODO: Your code here...
	return
}

// GetFollower implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFollower(ctx context.Context, req *relationdemo.GetFollowerRequest) (resp *relationdemo.GetFollowerResponse, err error) {
	// TODO: Your code here...
	return
}

// GetFriend implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFriend(ctx context.Context, req *relationdemo.GetFriendRequest) (resp *relationdemo.GetFriendResponse, err error) {
	// TODO: Your code here...
	return
}
