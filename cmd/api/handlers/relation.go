package handlers

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/golang/glog"

	"github.com/gitgou/simple_douyin/cmd/api/rpc"
	"github.com/gitgou/simple_douyin/kitex_gen/relationdemo"
	"github.com/gitgou/simple_douyin/pkg/constants"
	"github.com/gitgou/simple_douyin/pkg/errno"
	"github.com/hertz-contrib/jwt"
)

func Relation(ctx context.Context, c *app.RequestContext) {
	var relationVar RelationParam
	if err := c.Bind(&relationVar); err != nil {
		glog.Error("relation get param err. ", err.Error());
		SendErrResponse(c, errno.ConvertErr(err))
		return
	}

	claims := jwt.ExtractClaims(ctx, c)
	userId := int64(claims[constants.IdentityKey].(float64))

	if err := rpc.Relation(ctx, &relationdemo.RelationRequest{UserId: userId, ActionType: relationVar.ActionType,
		ToUserId: relationVar.ToUserId,}); err != nil{
			glog.Error("rpc relation err. ", err.Error());
			SendErrResponse(c, err)
			return	
		}
	
	SendResponse(c, map[string]interface{}{
		constants.StatusCode: 0})
}

func FollowList(ctx context.Context, c *app.RequestContext) {
	var followVar FollowListParam
	if err := c.Bind(&followVar); err != nil {
		glog.Error("follow get param err. ", err.Error());
		SendErrResponse(c, errno.ConvertErr(err))
		return
	}

	claims := jwt.ExtractClaims(ctx, c)
	userId := int64(claims[constants.IdentityKey].(float64))
	followList, err := rpc.GetFollowList(ctx, &relationdemo.GetFollowRequest{UserId: userId, })
	if err != nil{
		glog.Error("rpc get follow error, ", err.Error())
		SendErrResponse(c, err)
		return 
	}
	SendResponse(c, map[string]interface{}{
		constants.StatusCode : 0,
		constants.UserList : followList,
	})	
}

func FollowerList(ctx context.Context, c *app.RequestContext) {
	var followerVar FollowerListParam
	if err := c.Bind(&followerVar); err != nil {
		glog.Error("follower get param err. ", err.Error());
		SendErrResponse(c, errno.ConvertErr(err))
		return
	}

	claims := jwt.ExtractClaims(ctx, c)
	userId := int64(claims[constants.IdentityKey].(float64))
	followerList, err := rpc.GetFollowerList(ctx, &relationdemo.GetFollowerRequest{UserId: userId, })
	if err != nil{
		glog.Error("rpc get follower error, ", err.Error())
		SendErrResponse(c, err)
		return 
	}
	SendResponse(c, map[string]interface{}{
		constants.StatusCode : 0,
		constants.UserList : followerList,
	})	

	
}

func FriendList(ctx context.Context, c *app.RequestContext){
	var friendVar FriendListParam
	if err := c.Bind(&friendVar); err != nil {
		glog.Error("friend get param err. ", err.Error());
		SendErrResponse(c, errno.ConvertErr(err))
		return
	}

	claims := jwt.ExtractClaims(ctx, c)
	userId := int64(claims[constants.IdentityKey].(float64))
	friendList, err := rpc.GetFriendList(ctx, &relationdemo.GetFriendRequest{UserId: userId, })
	if err != nil{
		glog.Error("rpc get friend error, ", err.Error())
		SendErrResponse(c, err)
		return 
	}
	SendResponse(c, map[string]interface{}{
		constants.StatusCode : 0,
		constants.UserList : friendList,
	})	

}