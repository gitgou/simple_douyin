package handlers

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gitgou/simple_douyin/cmd/api/rpc"
	"github.com/gitgou/simple_douyin/kitex_gen/videodemo"
	"github.com/gitgou/simple_douyin/pkg/constants"
	"github.com/gitgou/simple_douyin/pkg/errno"
)

// Feed feed video data to user
func Feed(ctx context.Context, c *app.RequestContext) {
	var feedVar FeedRequest
	if err := c.Bind(&feedVar); err != nil {
		klog.Debugf("not get param lasttime: %d, token:%s", feedVar.LastestTime, feedVar.Token)
		//SendErrResponse(c, errno.ConvertErr(err))
	}

	claims, _ := JwtMiddleware.GetClaimsFromJWT(ctx, c)
	userId := int64(claims[constants.IdentityKey].(float64))
	videoList, nextTime, err := rpc.Feed(context.Background(), &videodemo.FeedRequest{
		LatestTime: feedVar.LastestTime,
		UserID: userId,
	})
	if err != nil {
		SendErrResponse(c, errno.ConvertErr(err))
		return
	}
	SendResponse(c, map[string]interface{}{
		constants.StatusCode: 0, constants.VideoList: videoList, constants.NextTime: nextTime,
	})
}
