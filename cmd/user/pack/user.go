package pack

import (
	"context"

	"github.com/gitgou/simple_douyin/cmd/user/dal/db"
	"github.com/gitgou/simple_douyin/cmd/user/rpc"
	"github.com/gitgou/simple_douyin/kitex_gen/redisdemo"
	"github.com/gitgou/simple_douyin/kitex_gen/userdemo"
	"github.com/gitgou/simple_douyin/pkg/constants"
)

func User(m *db.UserModel) *userdemo.User {
	if m == nil {
		return nil
	}

	user := &userdemo.User{
		Id:   m.ID,
		Name: m.Name,
		Avatar: m.AvatarUrl,
		BackgroundImage: m.BackgroundImage,
		Signature: m.Signature,
	}
	user = GetUserRedisInfo(user);
	return user;
}

func Users(ms []*db.UserModel) []*userdemo.User {
	if ms == nil || len(ms) == 0 {
		return nil
	}
	users := make([]*userdemo.User, 0)
	for _, m := range ms {
		if n := User(m); n != nil {
			users = append(users, n)
		}
	}
	return users
}

func GetUserRedisInfo(user *userdemo.User)*userdemo.User{
	userInfo := rpc.GetRedisUserInfo(context.Background(), &redisdemo.GetUserInfoRequest{UserId: user.Id})
	for _, info := range userInfo{
		switch info.Key {
		case constants.RedisZSetKeyFollow : 
			user.FollowCount = int64(info.Value)
		case constants.RedisZSetKeyFollower:
			user.FollowerCount = int64(info.Value)
		case constants.RedisZSetKeyVideo :
			user.WorkCount = int64(info.Value)
		case constants.RedisZSetKeyFavorite:
			user.FavoriteCount = int64(info.Value)
		case constants.RedisZSetKeyFavorited :
			user.TotalFavorited = int64(info.Value)
		} 
	}
	
	return user
}