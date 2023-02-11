package service

import (
	"context"

	"github.com/gitgou/simple_douyin/cmd/user/dal/db"
	"github.com/gitgou/simple_douyin/kitex_gen/demouser"
)

type UserService struct {
	ctx context.Context
}

// NewMGetNoteService new FeedService
func NewUserService(ctx context.Context) *UserService {
	return &UserService{ctx: ctx}
}

func (s *UserService) GetUser(req *demouser.GetUserRequest) (*db.UserModel, error) {
	return db.GetUser(s.ctx, req.UserId)
}
