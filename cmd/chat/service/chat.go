package service

import (
	"context"
	"time"

	"github.com/gitgou/simple_douyin/cmd/chat/dal/db"
	"github.com/gitgou/simple_douyin/cmd/chat/pack"
	"github.com/gitgou/simple_douyin/cmd/chat/rpc"
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
	msgModels, err := db.GetChat(s.ctx, req.UserId, req.ToUserId, time.UnixMicro(req.PreMsgTime)) //毫秒转化为秒
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
		return errno.UserNotExistErr;
	}
	if req.ActionType != int32(chatdemo.ChatAction_SEND_MSG) {
		return errno.ParamErr
	}

	return db.InsertMessages([]*db.MessageModel{
		{
			FromUserId: req.UserId,
			ToUserId: req.ToUserId,
			Content: req.Content,
		},
	})

	/* TODO 
	chatKey := utils.GenChatKey(req.UserId, req.ToUserId);
	//atomic.AddInt64(&cache.MessageSequenceId, 1)
	cache.MutexChat.Lock()
	defer cache.MutexChat.Unlock() 
	cache.MapNewChat[chatKey] = append(cache.MapNewChat[chatKey], &db.MessageModel{
		ID : rpc.IncreMsgId(s.ctx,&redisdemo.GetIncreIdRequest{Key: constants.ChatMsgIdKey}), //TODO 
		ToUserId: req.ToUserId,
		FromUserId: req.UserId,
		Content: req.Content,
		CreatedAt: time.Now(),
	})*/
}

func (s *ChatService)GetNewMsg(userId int64, toUserId int64)(*db.MessageModel){
	return db.GetNewMsg(s.ctx, userId, toUserId)
}

func (s * ChatService)Login(userId int64)error{
	//return cache.Login(userId)
	return nil
}