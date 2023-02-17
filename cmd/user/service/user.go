package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"

	"github.com/gitgou/simple_douyin/cmd/user/cache"
	"github.com/gitgou/simple_douyin/cmd/user/dal/db"
	"github.com/gitgou/simple_douyin/cmd/user/pack"
	"github.com/gitgou/simple_douyin/cmd/user/rpc"
	"github.com/gitgou/simple_douyin/kitex_gen/userdemo"
	"github.com/gitgou/simple_douyin/pkg/errno"
)

type UserService struct {
	ctx context.Context
}

// NewUserService new UserService
func NewUserService(ctx context.Context) *UserService {
	return &UserService{ctx: ctx}
}
//Get User By Id
func (s *UserService) GetUser(req *userdemo.GetUserRequest) (*db.UserModel, error) {
	if user, exist := cache.MapUser[req.UserId]; exist {
		return &user.User, nil
	}
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
//Deal User Login  Func
func (s *UserService) Login(req *userdemo.LoginRequest) (*db.UserModel, error) {
	h := md5.New()
	if _, err := io.WriteString(h, req.Password); err != nil {
		return nil, err
	}
	passWord := fmt.Sprintf("%x", h.Sum(nil))
	userName := req.Name

	if user, exist := cache.MapLoginUser[req.Token]; exist {
		if passWord != user.User.Password {
			return nil, errno.AuthorizationFailedErr
		} else {
			return &user.User, errno.UserAlreadyExistErr
		}
	}

	users, err := db.QueryUser(s.ctx, userName)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, errno.AuthorizationFailedErr
	}
	u := users[0]
	if u.Password != passWord {
		return nil, errno.AuthorizationFailedErr
	}
	//cache Login user, reduce I/O
	cache.Login(req.Token, *u)
	rpc.Login(s.ctx, u)
	return u, nil
}

// MGetUser multiple get list of user info
func (s *UserService) MGetUser(req *userdemo.MGetUserRequest) ([]*userdemo.User, error) {
	modelUsers, err := db.MGetUsers(s.ctx, req.UserIds)
	if err != nil {
		return nil, err
	}
	return pack.Users(modelUsers), nil
}

func (s *UserService)CheckUserOnline(userIds []int64)(bool){
	for _, userId := range userIds{
		u, exist := cache.MapUser[userId]; 
		if !exist {
			return false;
		}
		if u.IsOnline == false {
			return false;
		}
	}
	return true;
}