package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gitgou/simple_douyin/cmd/user/cache"
	"github.com/gitgou/simple_douyin/cmd/user/dal/db"
	"github.com/gitgou/simple_douyin/cmd/user/pack"
	"github.com/gitgou/simple_douyin/cmd/user/rpc"
	"github.com/gitgou/simple_douyin/kitex_gen/userdemo"
	"github.com/gitgou/simple_douyin/pkg/errno"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	ctx context.Context
}

// NewUserService new UserService
func NewUserService(ctx context.Context) *UserService {
	return &UserService{ctx: ctx}
}

// Get User By Id
func (s *UserService) GetUser(req *userdemo.GetUserRequest) (*db.UserModel, error) {
	cache.MutexUser.Lock()
	defer cache.MutexUser.Unlock()
	user := cache.GetUserById(req.UserId)
	if user != nil {
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
	// TODO md5存储密码不安全、易被破解；同时 哈希算法不可逆
	// 加盐哈希算法，（多次哈希或者添加盐值再哈希）
	// Bcrypt 可生成随机盐值的单向Hash加密算法
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost) //加密处理
	if err != nil {
		fmt.Println(err)
	}
	encodePWD := string(hash) // 保存在数据库的密码，虽然每次生成都不同，只需保存一份即可
	return db.CreateUser(s.ctx, &db.UserModel{
		Name:     req.Name,
		Password: encodePWD})

}

// Deal User Login  Func
func (s *UserService) Login(req *userdemo.LoginRequest) (*db.UserModel, error) {
	// TODO password 使用 bcrypt 算法


	userName := req.Name
	// TODO cache insert to redis
	if userId, exist := cache.MapLoginUser[req.Name]; exist {
		user, _ := cache.MapUser[userId]
		if err := bcrypt.CompareHashAndPassword([]byte(user.User.Password), []byte(req.Password)); err != nil{ //验证（对比）
			klog.Errorf("Login | password err, password:%s, real:%s, %s", req.Password, user.User.Password, err)
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
		klog.Errorf("not found user, userName: %s", userName)
		return nil, errno.AuthorizationFailedErr
	}
	u := users[0]
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password)); err != nil{ //验证（对比）
		klog.Errorf("Login | password err, password:%s, real:%s, %s", req.Password, u.Password, err)
		return nil, errno.AuthorizationFailedErr
	}
	//cache Login user to redis, reduce I/O 
	cache.Login(*u)
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
// Redis 缓存近期登陆用户及大V用户
func (s *UserService) CheckUserOnline(userIds []int64) bool {
	cache.MutexUser.Lock()
	defer cache.MutexUser.Unlock()
	for _, userId := range userIds {
		u := cache.GetUserById(userId)
		if u == nil {
			return false
		}
		if u.IsOnline == false {
			return false
		}
	}
	return true
}
