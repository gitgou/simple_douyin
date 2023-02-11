package handlers

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/gitgou/simple_douyin/cmd/api/rpc"
	"github.com/gitgou/simple_douyin/kitex_gen/demouser"
	"github.com/gitgou/simple_douyin/pkg/constants"
	"github.com/gitgou/simple_douyin/pkg/errno"
)

// Feed feed video data to user
func GetUser(ctx context.Context, c *app.RequestContext) {
	var userInfoVar UserRequest
	if err := c.Bind(&userInfoVar); err != nil {
		SendErrResponse(c, errno.ConvertErr(err))
		return
	}
	//TODO token 鉴权
	user, err := rpc.GetUser(context.Background(), &demouser.GetUserRequest{
		UserId: userInfoVar.UserId,
	})
	if err != nil {
		SendErrResponse(c, errno.ConvertErr(err))
		return
	}
	SendResponse(c, map[string]interface{}{
		constants.StatusCode: 0, constants.User: user})
}
