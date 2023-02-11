package handlers

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/gitgou/simple_douyin/pkg/constants"
	"github.com/gitgou/simple_douyin/pkg/errno"
	"github.com/hertz-contrib/jwt"
)

func Publish(ctx context.Context, c *app.RequestContext) {
	var publishVar PulishParam
	if err := c.Bind(&publishVar); err != nil {
		SendErrResponse(c, errno.ConvertErr(err))
		return
	}

	if len(publishVar.Token) == 0 || len(publishVar.Title) == 0 {
		SendErrResponse(c, errno.ParamErr)
	}

	claims := jwt.ExtractClaims(ctx, c)
	userID := int64(claims[constants.IdentityKey].(float64))

	//TODO video byte to url
}

func PublishList(ctx context.Context, c *app.RequestContext) {

}
