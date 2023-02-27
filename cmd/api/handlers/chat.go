package handlers

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gitgou/simple_douyin/cmd/api/rpc"
	"github.com/gitgou/simple_douyin/kitex_gen/chatdemo"
	"github.com/gitgou/simple_douyin/pkg/constants"
	"github.com/gitgou/simple_douyin/pkg/errno"
)


func GetChat(ctx context.Context, c *app.RequestContext) {
	var chatParam ChatParam
	if err := c.Bind(&chatParam); err != nil {
		SendErrResponse(c, errno.ConvertErr(err))
		return
	}
	claims, _ := JwtMiddleware.GetClaimsFromJWT(ctx, c)
	userId := int64(claims[constants.IdentityKey].(float64))
	klog.Infof("Get Chat, %d.", userId)
	msgs , err := rpc.GetChat(ctx, &chatdemo.ChatRequest{UserId: userId, ToUserId: chatParam.ToUserId, PreMsgTime: chatParam.PreMsgTime})
	if err != nil{
		klog.Errorf("Get Chat Err, userId:%d, err:%s", userId, err)
		SendErrResponse(c, err)
		return 
	}

	SendResponse(c, map[string]interface{}{
		constants.StatusCode: 0,
		constants.MsgList : msgs,
	})
}

func ChatAction(ctx context.Context, c *app.RequestContext) {
	klog.Errorf("ChatAction, test")
	var chatActionParam ChatActionParam 
	if err := c.Bind(&chatActionParam); err != nil {
		SendErrResponse(c, errno.ConvertErr(err))
		return
	}
	claims, _ := JwtMiddleware.GetClaimsFromJWT(ctx, c)
	userId := int64(claims[constants.IdentityKey].(float64))
	klog.Infof("Chat Action| send msg, userId: %d", userId)
	err := rpc.ChatAction(ctx, &chatdemo.ChatActionRequest{
		UserId: userId, 
		ToUserId: chatActionParam.ToUserId,
		Content: chatActionParam.Content,
		ActionType: int32(chatActionParam.ActionType),})
	if err != nil{
		SendErrResponse(c, err)
		return 
	}

	SendResponse(c, map[string]interface{}{
		constants.StatusCode: 0,
	})
}