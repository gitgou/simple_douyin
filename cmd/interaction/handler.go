package main

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gitgou/simple_douyin/cmd/chat/rpc"
	"github.com/gitgou/simple_douyin/cmd/interaction/pack"
	"github.com/gitgou/simple_douyin/cmd/interaction/service"
	"github.com/gitgou/simple_douyin/kitex_gen/interactiondemo"
	interaction "github.com/gitgou/simple_douyin/kitex_gen/interactiondemo"
	"github.com/gitgou/simple_douyin/kitex_gen/userdemo"
	"github.com/gitgou/simple_douyin/pkg/errno"
)

// InteractionImpl implements the last service interface defined in the IDL.
type InteractionserviceImpl struct {
}

// FavoriteAction implements the InteractionserviceImpl interface.
func (s *InteractionserviceImpl) FavoriteAction(ctx context.Context, req *interactiondemo.FavoriteRequest) (resp *interactiondemo.FavoriteResponse, err error) {
	resp = new(interaction.FavoriteResponse)
	//点赞操作(或者取消点赞)
	if err := service.NewInteractionService(ctx).FavoriteAction(req); err != nil {
		klog.Errorf("Favorite Action. Err:%s", err.Error())
		resp.Resp = pack.BuildBaseResp(err)
		return resp, err
	}
	resp.Resp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// FavoriteList implements the InteractionserviceImpl interface.
func (s *InteractionserviceImpl) FavoriteList(ctx context.Context, req *interactiondemo.GetFavoriteListRequest) (resp *interactiondemo.GetFavoriteListResponse, err error) {
	resp = new(interactiondemo.GetFavoriteListResponse)
	favoriteList, err := service.NewInteractionService(ctx).GetFavoriteList(ctx, req.UserId)
	if err != nil {
		klog.Errorf("FavoriteList | %s", err.Error())
		resp.Resp = pack.BuildBaseResp(err)
		return resp, err
	}
	resp.VideoList = pack.Videos(favoriteList)
	resp.Resp = pack.BuildBaseResp(errno.Success)
	return resp, err
}

// CommentAction implements the InteractionserviceImpl interface.
func (s *InteractionserviceImpl) CommentAction(ctx context.Context, req *interactiondemo.CommentRequest) (resp *interactiondemo.CommentResponse, err error) {
	resp = new(interactiondemo.CommentResponse)
	comment, err := service.NewInteractionService(ctx).Comment(ctx, req); 
	if err != nil{
		klog.Errorf("Comment Action Fail, %s", err.Error())
		resp.Resp = pack.BuildBaseResp(err)
		return resp, err
	}
	
	resp.Resp = pack.BuildBaseResp(errno.Success)
	user, _ := rpc.GetUser(ctx, &userdemo.GetUserRequest{UserId: comment.UserId})
	resp.Comment = pack.Comment(comment, user) 
	return resp, nil
}

// CommentList implements the InteractionserviceImpl interface.
func (s *InteractionserviceImpl) CommentList(ctx context.Context, req *interactiondemo.GetCommentListRequest) (resp *interactiondemo.GetCommentListResponse, err error) {
	resp = new(interactiondemo.GetCommentListResponse)
	commentList, err := service.NewInteractionService(ctx).GetCommentList(ctx, req.VideoId)
	if err != nil {
		resp.Resp = pack.BuildBaseResp(err)
		return resp, err
	}
	resp.Resp = pack.BuildBaseResp(errno.Success)
	resp.CommentList = pack.Comments(commentList)
	return resp, nil
}
