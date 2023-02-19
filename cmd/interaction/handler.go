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
	if req.action_type == 1 {
		err := service.NewInteractionService(ctx).FavoriteAction(req)
		return pack.BuildBaseResp(err)	
	}else if req.action_type == 2 {
		err := service.NewInteractionService(ctx).CancelFavoriteAction(req)
		return pack.BuildBaseResp(err)	
	}
}

// ShowFavoriteList implements the InteractionImpl interface.
func (s *InteractionImpl) ShowFavoriteList(ctx context.Context, req *interaction.DouyinFavoriteListRequest) (resp *interaction.DouyinFavoriteListResponse, err error) {
	// TODO: Your code here...
	resp = new(interaction.DouyinFavoriteListResponse)
	userid := req.user_id
	if userid <= 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp()
	resp.video_list, err := pack.Videos(service.NewInteractionService(ctx).ShowFavoriteList(req))
	return resp, err
}

// CommentAction implements the InteractionImpl interface.
func (s *InteractionImpl) CommentAction(ctx context.Context, req *interaction.DouyinCommentActionRequest) (resp *interaction.DouyinCommentActionResponse, err error) {
	// TODO: Your code here...
	resp = new(interaction.DouyinCommentActionResponse)
	if req.action_type == 1 {
		err := service.NewInteractionService(ctx).PublishComment(req)
		if err != nil {
			resp.BaseResp = pack.BuildBaseResp(err)
			return resp
		}
		resp.BaseResp = pack.BuildBaseResp()
		return resp,nil
	} else if req.action_type ==2 {
		err := service.NewInteractionService(ctx).CancelPublishComment(req)
		if err != nil {
			resp.BaseResp = pack.BuildBaseResp(err)
			return resp
		}
		resp.BaseResp = pack.BuildBaseResp()
		return resp,nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
}

// ShowCommentList implements the InteractionImpl interface.
func (s *InteractionImpl) ShowCommentList(ctx context.Context, req *interaction.DouyinCommentListRequest) (resp *interaction.DouyinCommentListResponse, err error) {
	// TODO: Your code here...
	resp = new(interaction.DouyinCommentActionResponse)
	comments, err := service.NewInteractionService(ctx).GetVideoComments(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp
	}
	resp.BaseResp = pack.BuildBaseResp()
	resp.comment_list = pack.Comments(comments)
	return resp, nil
}
