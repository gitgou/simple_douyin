package handlers

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/kitex/pkg/klog"

	"github.com/gitgou/simple_douyin/cmd/api/rpc"
	"github.com/gitgou/simple_douyin/kitex_gen/relationdemo"
	"github.com/gitgou/simple_douyin/pkg/constants"
	"github.com/gitgou/simple_douyin/pkg/errno"
	"github.com/gitgou/simple_douyin/pkg/utils"
)

func Relation(ctx context.Context, c *app.RequestContext) {
	var relationVar RelationParam
	klog.Infof("Relation| test")
	if err := c.Bind(&relationVar); err != nil {
		klog.Errorf("relation get param err.%s ", err.Error())
		SendErrResponse(c, errno.ConvertErr(err))
		return
	}

	userId := utils.GetUserIdInToken(relationVar.Token)
	klog.Infof("Relation| test, ID:%d", userId)
	if err := rpc.Relation(ctx, &relationdemo.RelationRequest{UserId: userId, ActionType: relationVar.ActionType,
		ToUserId: relationVar.ToUserId}); err != nil {
		klog.Errorf("rpc relation err. userId:%d, %s", userId, err.Error())
		SendErrResponse(c, err)
		return
	}

	SendResponse(c, map[string]interface{}{
		constants.StatusCode: 0})
}

func FollowList(ctx context.Context, c *app.RequestContext) {
	var followVar FollowListParam
	if err := c.Bind(&followVar); err != nil {
		klog.Errorf("follow get param err. %s", err.Error())
		SendErrResponse(c, errno.ConvertErr(err))
		return
	}

	//claims := jwt.ExtractClaims(ctx, c)
	//userId := int64(claims[constants.IdentityKey].(float64))
	userId := utils.GetUserIdInToken(followVar.Token)
	followList, err := rpc.GetFollowList(ctx, &relationdemo.GetFollowRequest{UserId: userId})
	if err != nil {
		klog.Errorf("rpc get follow error,userId:%d, %s", userId, err.Error())
		SendErrResponse(c, err)
		return
	}
	SendResponse(c, map[string]interface{}{
		constants.StatusCode: 0,
		constants.UserList:   followList,
	})
}

func FollowerList(ctx context.Context, c *app.RequestContext) {
	var followerVar FollowerListParam
	if err := c.Bind(&followerVar); err != nil {
		klog.Errorf("follower get param err. %s", err.Error())
		SendErrResponse(c, errno.ConvertErr(err))
		return
	}

	//claims := jwt.ExtractClaims(ctx, c)
	//userId := int64(claims[constants.IdentityKey].(float64))
	userId := utils.GetUserIdInToken(followerVar.Token)
	followerList, err := rpc.GetFollowerList(ctx, &relationdemo.GetFollowerRequest{UserId: userId})
	if err != nil {
		klog.Errorf("rpc get follower error, ID:%d, %s", userId, err.Error())
		SendErrResponse(c, err)
		return
	}
	SendResponse(c, map[string]interface{}{
		constants.StatusCode: 0,
		constants.UserList:   followerList,
	})

}

func FriendList(ctx context.Context, c *app.RequestContext) {
	var friendVar FriendListParam
	if err := c.Bind(&friendVar); err != nil {
		klog.Errorf("friend get param err. %s", err.Error())
		SendErrResponse(c, errno.ConvertErr(err))
		return
	}

	//claims := jwt.ExtractClaims(ctx, c)
	//userId := int64(claims[constants.IdentityKey].(float64))
	userId := utils.GetUserIdInToken(friendVar.Token)
	klog.Infof("friendlist test| %s, userId:%d", friendVar.Token, userId)
	friendList, err := rpc.GetFriendList(ctx, &relationdemo.GetFriendRequest{UserId: userId})
	if err != nil {
		klog.Errorf("rpc get friend error,Id:%d, %s", userId, err.Error())
		SendErrResponse(c, err)
		return
	}
	SendResponse(c, map[string]interface{}{
		constants.StatusCode: 0,
		constants.UserList:   friendList,
	})

}
