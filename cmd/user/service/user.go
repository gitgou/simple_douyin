package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"

	"github.com/gitgou/simple_douyin/cmd/user/dal/db"
	"github.com/gitgou/simple_douyin/cmd/user/pack"
	"github.com/gitgou/simple_douyin/kitex_gen/userdemo"
	"github.com/gitgou/simple_douyin/pkg/errno"
)

type UserService struct {
	ctx context.Context
}

// NewMGetNoteService new FeedService
func NewUserService(ctx context.Context) *UserService {
	return &UserService{ctx: ctx}
}

func (s *UserService) GetUser(req *userdemo.GetUserRequest) (*db.UserModel, error) {
	return db.GetUser(s.ctx, req.UserId)
}

func (s *UserService) CreateUser(req *userdemo.CreateUserRequest) (int64, error) {
	users, err := db.QueryUser(s.ctx, req.Name)
	if err != nil {
		return 0, err
	}
	if len(users) != 0 {
		return 0, errno.UserAlreadyExistErr
	}

	h := md5.New()
	if _, err = io.WriteString(h, req.Password); err != nil {
		return 0, err
	}
	password := fmt.Sprintf("%x", h.Sum(nil))
	return db.CreateUser(s.ctx, &db.UserModel{
		Name:     req.Name,
		Password: password})

}

func (s *UserService) CheckUser(req *userdemo.CheckUserRequest) (int64, error) {
	h := md5.New()
	if _, err := io.WriteString(h, req.Password); err != nil {
		return 0, err
	}
	passWord := fmt.Sprintf("%x", h.Sum(nil))

	userName := req.Name
	users, err := db.QueryUser(s.ctx, userName)
	if err != nil {
		return 0, err
	}
	if len(users) == 0 {
		return 0, errno.AuthorizationFailedErr
	}
	u := users[0]
	if u.Password != passWord {
		return 0, errno.AuthorizationFailedErr
	}

	//TODO store user info in cache memory
	return u.ID, nil
}

// MGetUser multiple get list of user info
func (s *UserService) MGetUser(req *userdemo.MGetUserRequest) ([]*userdemo.User, error) {
	modelUsers, err := db.MGetUsers(s.ctx, req.UserIds)
	if err != nil {
		return nil, err
	}
	return pack.Users(modelUsers), nil
}
