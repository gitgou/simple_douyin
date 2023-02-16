package rpc

import (
	"context"

	"github.com/gitgou/simple_douyin/cmd/user/dal/db"
)

//Login API,if other service need, add service in this func
func Login(ctx context.Context, user *db.UserModel){
	ChatLogin(ctx, user.ID)
}