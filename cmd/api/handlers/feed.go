package handlers

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/gitgou/simple_douyin/cmd/api/rpc"
	"github.com/gitgou/simple_douyin/kitex_gen/demofeed"
	"github.com/gitgou/simple_douyin/pkg/constants"
	"github.com/gitgou/simple_douyin/pkg/errno"
	"github.com/hertz-contrib/jwt"
)

// Feed feed video data to user
func Feed(ctx context.Context, c *app.RequestContext) {
	var feedVar FeedRequest
	if err := c.Bind(&feedVar); err != nil {
		SendErrResponse(c, errno.ConvertErr(err))
		return
	}
	claims := jwt.ExtractClaims(ctx, c)
	userID := int64(claims[constants.IdentityKey].(float64))

	videoList, nextTime, err := rpc.Feed(context.Background(), &demofeed.FeedRequest{
		UserID:     userID,
		LatestTime: feedVar.LastestTime,
	})
	if err != nil {
		SendErrResponse(c, errno.ConvertErr(err))
		return
	}
	SendResponse(c, map[string]interface{}{
		constants.StatusCode: 0, constants.VideoList: videoList, constants.NextTime: nextTime,
	})
}
