package pack

import (
	"github.com/gitgou/simple_douyin/cmd/chat/dal/db"
	"github.com/gitgou/simple_douyin/kitex_gen/chatdemo"
)

func Message(msgModel *db.MessageModel) *chatdemo.Message{
	if msgModel == nil {
		return nil
	}

	return &chatdemo.Message{
		Id:   msgModel.ID,
		ToUserId: msgModel.ToUserId,
		FromUserId: msgModel.FromUserId,
		Content: msgModel.Content,
		CreateTime: msgModel.CreatedAt.Unix(),
	}
}

func Messages(msgModels []*db.MessageModel)[]*chatdemo.Message{

	if msgModels == nil || len(msgModels) == 0 {
		return nil
	}
	msgs := make([]*chatdemo.Message, 0)
	for _, m := range msgModels {
		if n := Message(m); n != nil {
			msgs = append(msgs, n)
		}
	}
	return msgs
}
