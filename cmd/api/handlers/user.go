package handlers

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/gitgou/simple_douyin/cmd/api/rpc"
	"github.com/gitgou/simple_douyin/kitex_gen/userdemo"
	"github.com/gitgou/simple_douyin/pkg/constants"
	"github.com/gitgou/simple_douyin/pkg/errno"
)

// Feed feed video data to user
func GetUser(ctx context.Context, c *app.RequestContext) {
	var userParam GetUserParam
	if err := c.Bind(&userParam); err != nil {
		SendErrResponse(c, errno.ConvertErr(err))
		return
	}
	if userParam.UserId < 0 {
		SendErrResponse(c, errno.ParamErr)
		return
	}
	//TODO token 鉴权
	user, err := rpc.GetUser(context.Background(), &userdemo.GetUserRequest{
		UserId: userParam.UserId,
	})
	if err != nil {
		SendErrResponse(c, errno.ConvertErr(err))
		return
	}
	SendResponse(c, map[string]interface{}{
		constants.StatusCode: 0, constants.User: user})
}

func Register(ctx context.Context, c *app.RequestContext) {
	var userParam UserParam
	if err := c.Bind(&userParam); err != nil {
		SendErrResponse(c, errno.ConvertErr(err))
		return
	}
	if len(userParam.UserName) == 0 || len(userParam.Password) == 0 {
		SendErrResponse(c, errno.ParamErr)
		return
	}
	userId, err := rpc.CreateUser(context.Background(), &userdemo.CreateUserRequest{
		Name:     userParam.UserName,
		Password: userParam.Password,
	})
	if err != nil {
		SendErrResponse(c, errno.ConvertErr(err))
		return
	}
	SendResponse(c, map[string]interface{}{
		constants.StatusCode: 0, constants.Token: "", constants.UserID: userId})
}
