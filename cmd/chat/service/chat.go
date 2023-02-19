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
	"github.com/gitgou/simple_douyin/kitex_gen/redisdemo"
	"github.com/gitgou/simple_douyin/kitex_gen/userdemo"
	"github.com/gitgou/simple_douyin/pkg/constants"
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
	//atomic.AddInt64(&cache.MessageSequenceId, 1)
	cache.MutexChat.Lock()
	defer cache.MutexChat.Unlock() 
	cache.MapNewChat[chatKey] = append(cache.MapNewChat[chatKey], &db.MessageModel{
		ID : rpc.IncreMsgId(s.ctx,&redisdemo.GetIncreIdRequest{Key: constants.ChatMsgIdKey}), //TODO 
		ToUserId: req.ToUserId,
		FromUserId: req.UserId,
		Content: req.Content,
		CreateAt: time.Now(),
	})
	//TODO 发送给对方
	return nil
}

func (s * ChatService)Login(userId int64)error{
	return cache.Login(userId)
}