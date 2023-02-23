package handlers

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gitgou/simple_douyin/cmd/api/rpc"
	"github.com/gitgou/simple_douyin/kitex_gen/interactiondemo"
	"github.com/gitgou/simple_douyin/pkg/constants"
	"github.com/gitgou/simple_douyin/pkg/errno"
	"github.com/gitgou/simple_douyin/pkg/utils"
)

func FavoriteAction(ctx context.Context, c *app.RequestContext) {
	var favoriteParam FavoriteActionParam
	klog.Infof("FavoriteAction test.")
	if err := c.Bind(&favoriteParam); err != nil {
		klog.Errorf("Get  Bind Param Err, %s", err.Error())
		SendErrResponse(c, errno.ConvertErr(err))
		return
	}
	//claims := jwt.ExtractClaims(ctx, c)
	//userId := int64(claims[constants.IdentityKey].(float64))
	//TODO token 鉴权
	userId := utils.GetUserIdInToken(favoriteParam.Token)
	err := rpc.FavoriteAction(ctx, &interactiondemo.FavoriteRequest{
		UserId: userId,
		ActionType: int32(favoriteParam.ActionType),
		VideoId: favoriteParam.VideoId,
	})
	if err != nil {
		klog.Errorf("Favorite Action Err, %s", err.Error())
		SendErrResponse(c, errno.ConvertErr(err))
		return
	}
	SendResponse(c, map[string]interface{}{
		constants.StatusCode: 0, })
}


func FavoriteList(ctx context.Context, c *app.RequestContext) {
	var favoriteListParam FavoriteListParam
	klog.Infof("FavoriteList test.")
	if err := c.Bind(&favoriteListParam); err != nil {
		klog.Errorf("Get  Bind Param Err, %s", err.Error())
		SendErrResponse(c, errno.ConvertErr(err))
		return
	}
	//claims := jwt.ExtractClaims(ctx, c)
	//userId := int64(claims[constants.IdentityKey].(float64))
	//TODO token 鉴权
	userId := utils.GetUserIdInToken(favoriteListParam.Token)
	videoList,err := rpc.GetFavoriteList(ctx, &interactiondemo.GetFavoriteListRequest{
		UserId: userId,
	})
	if err != nil {
		klog.Errorf("Favorite List Err, %s", err.Error())
		SendErrResponse(c, errno.ConvertErr(err))
		return
	}
	SendResponse(c, map[string]interface{}{
		constants.StatusCode: 0, 
		constants.VideoList : videoList, 
	})
}

func CommentAction(ctx context.Context, c *app.RequestContext) {
	var commentParam CommentActionParam
	klog.Infof("Comment Action test.")
	if err := c.Bind(&commentParam); err != nil {
		klog.Errorf("Get  Bind Param Err, %s", err.Error())
		SendErrResponse(c, errno.ConvertErr(err))
		return
	}
	//claims := jwt.ExtractClaims(ctx, c)
	//userId := int64(claims[constants.IdentityKey].(float64))
	//TODO token 鉴权
	userId := utils.GetUserIdInToken(commentParam.Token)
	comment, err := rpc.CommentAction(ctx, &interactiondemo.CommentRequest{
		UserId: userId,
		ActionType: int32(commentParam.ActionType),
		VideoId : commentParam.VideoId,
		CommentId: commentParam.CommentId,
		CommentText : commentParam.CommentText,
	})
	if err != nil {
		klog.Errorf("Comment Action Err, %s", err.Error())
		SendErrResponse(c, errno.ConvertErr(err))
		return
	}
	SendResponse(c, map[string]interface{}{
		constants.StatusCode: 0,
		constants.Comment : comment,
	})
}


func CommentList(ctx context.Context, c *app.RequestContext) {
	var commentListParam CommentListParam
	klog.Infof("Comment List test.")
	if err := c.Bind(&commentListParam); err != nil {
		klog.Errorf("Get Bind Param Err, %s", err.Error())
		SendErrResponse(c, errno.ConvertErr(err))
		return
	}
	//claims := jwt.ExtractClaims(ctx, c)
	//userId := int64(claims[constants.IdentityKey].(float64))
	//TODO token 鉴权
	userId := utils.GetUserIdInToken(commentListParam.Token)
	commentList,err := rpc.GetCommentList(ctx, &interactiondemo.GetCommentListRequest{
		UserId: userId,
		VideoId: commentListParam.VideoId,
	})
	if err != nil {
		klog.Errorf("Favorite List Err, %s", err.Error())
		SendErrResponse(c, errno.ConvertErr(err))
		return
	}
	SendResponse(c, map[string]interface{}{
		constants.StatusCode: 0, 
		constants.CommentList : commentList, 
	})
}


