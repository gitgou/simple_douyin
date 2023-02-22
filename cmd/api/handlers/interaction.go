package handlers

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gitgou/simple_douyin/cmd/api/rpc"
	"github.com/gitgou/simple_douyin/kitex_gen/interaction"
	"github.com/gitgou/simple_douyin/pkg/constants"
	"github.com/gitgou/simple_douyin/pkg/errno"
	"github.com/gitgou/simple_douyin/pkg/utils"
	"strconv"
)

// interaction interface

func Favorite(ctx context.Context, c *app.RequestContext) {
	var favoriteVar FavoriteParam
	klog.Infof("interaction | test")
	if err := c.Bind(&favoriteVar); err != nil {
		klog.Errorf("interaction Favor  get param err.%s ", err.Error())
		SendErrResponse(c, errno.ConvertErr(err))
		return
	}

	userId := utils.GetUserIdInToken(favoriteVar.Token)
	klog.Infof("interaction | test, ID:%d", userId)
	action_type, err := strconv.Atoi(favoriteVar.ActionType)
	if err != nil {
		klog.Errorf("Parameter err. userId:%d, %s", userId, err.Error())
		SendErrResponse(c, err)
		return
	}
	if err := rpc.Favorite(ctx, &interaction.DouyinFavoriteActionRequest{
		UserId		: 	userId,
		VideoId		:	favoriteVar.VideoId,
		ActionType	: 	action_type,}); err != nil {
		klog.Errorf("rpc interaction err. userId:%d, %s", userId, err.Error())
		SendErrResponse(c, err)
		return
	}

	SendResponse(c, map[string]interface{}{
		constants.StatusCode: 0})
}

func FavoriteList(ctx context.Context, c *app.RequestContext) {

	var favoriteListVar FavoriteListParam
	klog.Infof("interaction | test")
	if err := c.Bind(&favoriteListVar); err != nil {
		klog.Errorf("interaction FavorList get param err.%s ", err.Error())
		SendErrResponse(c, errno.ConvertErr(err))
		return
	}

	userId := utils.GetUserIdInToken(favoriteListVar.Token)
	klog.Infof("interaction | test, ID:%d", userId)
	if VideoList, err := rpc.FavoriteList(ctx, &interaction.DouyinFavoriteListRequest{Token: FavoriteListParam.Token,}); err != nil {
		klog.Errorf("rpc interaction err. userId:%d, %s", userId, err.Error())
		SendErrResponse(c, err)
		return
	}

	SendResponse(c, map[string]interface{}{
		constants.StatusCode: 0,
		constants.StatusMsg: "success",
		constants.VideoList: VideoList,
	})

}

func Comment(ctx context.Context, c *app.RequestContext) {
	var commentVar CommentParam
	klog.Infof("interaction comment | test")
	if err := c.Bind(&commentVar); err != nil {
		klog.Errorf("interaction comment List get param err.%s ", err.Error())
		SendErrResponse(c, errno.ConvertErr(err))
		return
	}

	userId := utils.GetUserIdInToken(commentVar.Token)
	klog.Infof("interaction | test, ID:%d", userId)
	video_id, err := strconv.Atoi(commentVar.VideoId)
	if err != nil {
		klog.Errorf("Parameter err. userId:%d, %s", userId, err.Error())
		SendErrResponse(c, err)
		return
	}
	action_type, err := strconv.Atoi(commentVar.ActionType)
	if err != nil {
		klog.Errorf("Parameter err. userId:%d, %s", userId, err.Error())
		SendErrResponse(c, err)
		return
	}
	comment_id, err := strconv.Atoi(commentVar.CommentId)
	if err != nil {
		klog.Errorf("Parameter err. userId:%d, %s", userId, err.Error())
		SendErrResponse(c, err)
		return
	}
	if Comment, err := rpc.Comment(ctx, &interaction.DouyinCommentActionRequest{
		UserId		:	userId,
		VideoId		:	video_id,
		ActionType	:	action_type,
		CommentText	:	commentVar.CommentText,
		CommentId	:	comment_id,}); err != nil {
		klog.Errorf("rpc interaction err. userId:%d, %s", userId, err.Error())
		SendErrResponse(c, err)
		return
	}

	SendResponse(c, map[string]interface{}{
		constants.StatusCode: 0,
		constants.StatusMsg: "success",
		constants.Comment: Comment,
	})
}

func CommentList(ctx context.Context, c *app.RequestContext) {
	var commentListVar CommentListParam
	klog.Infof("interaction commentlist | test")
	if err := c.Bind(&commentListVar); err != nil {
		klog.Errorf("interaction commentList get param err.%s ", err.Error())
		SendErrResponse(c, errno.ConvertErr(err))
		return
	}

	userId := utils.GetUserIdInToken(commentListVar.Token)
	klog.Infof("interaction | test, ID:%d", userId)
	video_id, err := strconv.Atoi(commentListVar.VideoId)
	if err != nil {
		klog.Errorf("Parameter err. userId:%d, %s", userId, err.Error())
		SendErrResponse(c, err)
		return
	}
	if Comment, err := rpc.CommentList(ctx, &interaction.DouyinCommentListRequest {
		Token		:	commentListVar.Token,
		VideoId		:	video_id,}); err != nil {
		klog.Errorf("rpc interaction commentlist err. userId:%d, %s", userId, err.Error())
		SendErrResponse(c, err)
		return
	}
	SendResponse(c, map[string]interface{}{
		constants.StatusCode: 0,
		constants.StatusMsg: "success",
		constants.CommentList: Comment,
	})
}


