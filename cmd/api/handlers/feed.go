package handlers 

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"cmd/api/rpc"
	"kitex_gen/demofeed"
	"pkg/constants"
	"pkg/errno"
	"github.com/hertz-contrib/jwt"

)
//Feed feed video data to user
func Feed(ctx context.Context, c *app.RequestContext) {
	var feedVar FeedRequest
	if err := c.Bind(&feedVar); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	
	claims := jwt.ExtractClaims(ctx, c)
	userID := int64(claims[constants.IdentityKey].(float64))
	err := rpc.Feed(context.Background(), &notedemo.CreateNoteRequest{
		UserId:  userID,
		Content: noteVar.Content, Title: noteVar.Title,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, nil)
}
