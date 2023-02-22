package main

import (
	"context"
	interaction "simple_douyin-master/cmd/interaction/kitex_gen/interaction"
	"simple_douyin-master/cmd/interaction/pack"
	"simple_douyin-master/cmd/interaction/service"
	"github.com/gitgou/simple_douyin/pkg/errno"
)

// InteractionImpl implements the last service interface defined in the IDL.
type InteractionImpl struct{

}

// FavoriteAction implements the InteractionImpl interface.
func (s *InteractionImpl) FavoriteAction(ctx context.Context, req *interaction.DouyinFavoriteActionRequest) (resp *interaction.DouyinFavoriteActionResponse, err error) {
	// TODO: Your code here...
	resp = new(interaction.DouyinFavoriteActionResponse)
	//点赞操作(或者取消点赞)
	if req.ActionType == 1 {
		err := service.NewInteractionService(ctx).FavoriteAction(req)
		return pack.BuildBaseResp(err)	
	}else if req.ActionType == 2 {
		err := service.NewInteractionService(ctx).CancelFavoriteAction(req)
		return pack.BuildBaseResp(err)	
	}
}

// ShowFavoriteList implements the InteractionImpl interface.
func (s *InteractionImpl) ShowFavoriteList(ctx context.Context, req *interaction.DouyinFavoriteListRequest) (resp *interaction.DouyinFavoriteListResponse, err error) {
	// TODO: Your code here...
	resp = new(interaction.DouyinFavoriteListResponse)
	userid := req.UserId
	if userid <= 0 {
		resp.Resp = pack.BuildBaseResp(errno.ParamErr)
		return resp, ParamErr
	}
	resp.Resp = pack.BuildBaseResp()
	videoList, err := service.NewInteractionService(ctx).GetFavoriteList(req)
	if err != nil {
		return resp, err
	}
	resp.VideoList = pack.Videos(videoList, req.UserId)
	return resp, err
}

// CommentAction implements the InteractionImpl interface.
func (s *InteractionImpl) CommentAction(ctx context.Context, req *interaction.DouyinCommentActionRequest) (resp *interaction.DouyinCommentActionResponse, err error) {
	// TODO: Your code here...
	resp = new(interaction.DouyinCommentActionResponse)
	if req.ActionType == 1 {
		comment, err := service.NewInteractionService(ctx).PublishComment(req)
		if err != nil {
			resp.Resp = pack.BuildBaseResp(err)
			return resp, err
		}
		resp.Resp = pack.BuildBaseResp()
		resp.Comment = pack.Comment(comment, req.User.Id)
		return resp, nil
	} else if req.ActionType ==2 {
		err := service.NewInteractionService(ctx).CancelPublishComment(req)
		if err != nil {
			resp.Resp = pack.BuildBaseResp(err)
			return resp, err
		}
		resp.Resp = pack.BuildBaseResp()
		return resp,nil
	}
	resp.Resp = pack.BuildBaseResp(errno.ParamErr)
	return resp, ParamErr
}

// ShowCommentList implements the InteractionImpl interface.
func (s *InteractionImpl) ShowCommentList(ctx context.Context, req *interaction.DouyinCommentListRequest) (resp *interaction.DouyinCommentListResponse, err error) {
	// TODO: Your code here...
	resp = new(interaction.DouyinCommentActionResponse)
	comments, err := service.NewInteractionService(ctx).GetVideoComments(req)
	if err != nil {
		resp.Resp = pack.BuildBaseResp(err)
		return resp
	}
	resp.Resp = pack.BuildBaseResp()
	resp.comment_list = pack.Comments(comments, req.User.Id)
	return resp, nil
}
