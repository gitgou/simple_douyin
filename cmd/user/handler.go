package main

import (
	"context"

	"github.com/gitgou/simple_douyin/cmd/user/pack"
	"github.com/gitgou/simple_douyin/cmd/user/service"
	demouser "github.com/gitgou/simple_douyin/kitex_gen/demouser"
	"github.com/gitgou/simple_douyin/pkg/errno"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// GetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUser(ctx context.Context, req *demouser.GetUserRequest) (resp *demouser.GetUserResponse, err error) {
	resp = new(demouser.GetUserResponse)
	if req.UserId <= 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	userModel, err := service.NewUserService(ctx).GetUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	user := pack.User(userModel)
	//TODO get follows data
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.User = user
	return resp, nil
}

// MGetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) MGetUser(ctx context.Context, req *demouser.MGetUserRequest) (resp *demouser.MGetUserResponse, err error) {
	// TODO: Your code here...
	return
}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *demouser.CreateUserRequest) (resp *demouser.CreateUserResponse, err error) {
	// TODO: Your code here...
	return
}

// CheckUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CheckUser(ctx context.Context, req *demouser.CheckUserRequest) (resp *demouser.CheckUserResponse, err error) {
	// TODO: Your code here...
	return
}
