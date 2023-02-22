package main

import (
	"context"
	chatdemo "github.com/gitgou/simple_douyin/kitex_gen/chatdemo"
	"github.com/gitgou/simple_douyin/cmd/chat/service"
	"github.com/gitgou/simple_douyin/cmd/chat/pack"
	"github.com/gitgou/simple_douyin/pkg/errno"
)

// ChatServiceImpl implements the last service interface defined in the IDL.
type ChatServiceImpl struct{}

// GetChat implements the ChatServiceImpl interface.
func (s *ChatServiceImpl) GetChat(ctx context.Context, req *chatdemo.ChatRequest) (resp *chatdemo.ChatResponse, err error) {
	resp = new(chatdemo.ChatResponse)

	msgList, err := service.NewChatService(ctx).GetChat(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.MessageList = msgList
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// ChatAction implements the ChatServiceImpl interface.
func (s *ChatServiceImpl) ChatAction(ctx context.Context, req *chatdemo.ChatActionRequest) (resp *chatdemo.ChatActionResponse, err error) {
	resp = new(chatdemo.ChatActionResponse)

	if err := service.NewChatService(ctx).ChatAction(req); err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// Login implements the ChatServiceImpl interface.
func (s *ChatServiceImpl) Login(ctx context.Context, req *chatdemo.ChatLoginRequest) (resp *chatdemo.ChatLoginResponse, err error) {
	resp = new(chatdemo.ChatLoginResponse)

	if err := service.NewChatService(ctx).Login(req.UserId); err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// GetNewMsg implements the ChatServiceImpl interface.
func (s *ChatServiceImpl) GetNewMsg(ctx context.Context, req *chatdemo.GetNewMsgRequest) (resp *chatdemo.GetNewMsgResponse, err error) {
	resp = new(chatdemo.GetNewMsgResponse)

	msg:= service.NewChatService(ctx).GetNewMsg(req.UserId, req.ToUserId)
	if msg == nil {
		return resp, nil
	}
	resp.Msg = msg.Content;
	if msg.FromUserId == req.UserId {
		resp.MsgType = 1
	}else{
		resp.MsgType = 0 // TODO 解决魔术数字,设置枚举，时间紧
	}
	return resp, nil
}
