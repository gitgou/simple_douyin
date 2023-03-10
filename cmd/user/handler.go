package main

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gitgou/simple_douyin/cmd/user/pack"
	"github.com/gitgou/simple_douyin/cmd/user/service"
	userdemo "github.com/gitgou/simple_douyin/kitex_gen/userdemo"
	"github.com/gitgou/simple_douyin/pkg/errno"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// GetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUser(ctx context.Context, req *userdemo.GetUserRequest) (resp *userdemo.GetUserResponse, err error) {
	resp = new(userdemo.GetUserResponse)
	if req.UserId <= 0 {
		klog.Errorf("Get User ParamErr, %d", req.UserId)
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	userModel, err := service.NewUserService(ctx).GetUser(req)
	if err != nil {
		klog.Errorf("Get User Err, %s", err.Error())
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
func (s *UserServiceImpl) MGetUser(ctx context.Context, req *userdemo.MGetUserRequest) (resp *userdemo.MGetUserResponse, err error) {
	resp = new(userdemo.MGetUserResponse)

	if len(req.UserIds) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	users, err := service.NewUserService(ctx).MGetUser(req)
	if err != nil {
		klog.Errorf("MGet User Err, %s", err.Error())
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.Users = users
	return resp, nil
}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *userdemo.CreateUserRequest) (resp *userdemo.CreateUserResponse, err error) {
	resp = new(userdemo.CreateUserResponse)
	if len(req.Name) == 0 || len(req.Password) == 0 {
		klog.Errorf("Create User Err, ParamErr")
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	userId, err := service.NewUserService(ctx).CreateUser(req)
	if err != nil { 
		klog.Errorf("Create User Err, Err: %s", err.Error())
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.UserId = userId
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *userdemo.LoginRequest) (resp *userdemo.LoginResponse, err error) {
	resp = new(userdemo.LoginResponse)

	if len(req.Name) == 0 || len(req.Password) == 0 {
		klog.Errorf("Login User Err, ParamErr")
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	user, err := service.NewUserService(ctx).Login(req)
	if err != nil {
		klog.Errorf("Login User Err, %s", err.Error())
		if user != nil {
			resp.UserId = user.ID
		}
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	
	resp.UserId = user.ID
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// CheckUserOnline implements the UserServiceImpl interface.
func (s *UserServiceImpl) CheckUserOnline(ctx context.Context, req *userdemo.CheckUserOnlineRequest) (resp *userdemo.CheckUserOnlineResponse, err error) {
	resp = new(userdemo.CheckUserOnlineResponse)

	isOnline := service.NewUserService(ctx).CheckUserOnline(req.UserIds)
	if ! isOnline {
		resp.BaseResp = pack.BuildBaseResp(errno.UserNotExistErr)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}
