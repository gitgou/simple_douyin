package service

import (
	"context"
	"time"

	"github.com/gitgou/simple_douyin/cmd/chat/cache"
	"github.com/gitgou/simple_douyin/cmd/chat/dal/db"
	"github.com/gitgou/simple_douyin/cmd/chat/pack"
	"github.com/gitgou/simple_douyin/cmd/chat/rpc"
	"github.com/gitgou/simple_douyin/cmd/chat/utils"
	"github.com/gitgou/simple_douyin/kitex_gen/chatdemo"
	"github.com/gitgou/simple_douyin/kitex_gen/userdemo"
	"github.com/gitgou/simple_douyin/pkg/errno"
)

type ChatService struct {
	ctx context.Context
}

// NewChatService 
func NewChatService(ctx context.Context) *ChatService {
	return &ChatService{ctx: ctx}
}
//Get Chat request
func (s *ChatService)GetChat(req *chatdemo.ChatRequest)([]*chatdemo.Message, error) {
	msgModels, err := db.GetChat(s.ctx, req.UserId, req.ToUserId)
	if err != nil{
		return nil, err
	}
	return pack.Messages(msgModels), nil
}
func (s* ChatService)ChatAction(req *chatdemo.ChatActionRequest)(error){
	_ ,err := rpc.GetUser(s.ctx, &userdemo.GetUserRequest{
		UserId: req.ToUserId,
	});
	if err != nil {
		return errno.UserAlreadyExistErr;
	}
	if req.ActionType != int32(chatdemo.ChatAction_SEND_MSG) {
		return errno.ParamErr
	}
	chatKey := utils.GenChatKey(req.UserId, req.ToUserId);
	cache.MapChat[chatKey] = append(cache.MapChat[chatKey], &db.MessageModel{
		ID : 1, //TODO 
		ToUserId: req.ToUserId,
		FromUserId: req.UserId,
		Content: req.Content,
		CreateTime: time.Now().Format(time.Kitchen),
	})
	//TODO 发送给对方
	return nil
}