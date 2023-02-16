package handlers

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	//"github.com/gitgou/simple_douyin/cmd/api/rpc"
	"github.com/gitgou/simple_douyin/pkg/constants"
	"github.com/gitgou/simple_douyin/pkg/errno"
	"github.com/hertz-contrib/jwt"
)

func Relation(ctx context.Context, c *app.RequestContext) {
	var relationVar RelationParam
	if err := c.Bind(&relationVar); err != nil {
		SendErrResponse(c, errno.ConvertErr(err))
		return
	}

	if len(relationVar.Token) == 0  {
		SendErrResponse(c, errno.ParamErr)
	}
	claims := jwt.ExtractClaims(ctx, c)
	userId := int64(claims[constants.IdentityKey].(float64))
	
}

func FollowList(ctx context.Context, c *app.RequestContext) {

	
}

func FollowerList(ctx context.Context, c *app.RequestContext) {

	
}

func FriendList(ctx context.Context, c *app.RequestContext){

}