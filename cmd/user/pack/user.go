package pack

import (
	"github.com/gitgou/simple_douyin/cmd/user/dal/db"
	"github.com/gitgou/simple_douyin/kitex_gen/demouser"
)

func User(m *db.UserModel) *demouser.User {
	if m == nil {
		return nil
	}

	return &demouser.User{
		Id:   m.ID,
		Name: m.Name,
	}
}

func Users(ms []*db.UserModel) []*demouser.User {
	if ms == nil || len(ms) == 0 {
		return nil
	}
	users := make([]*demouser.User, 0)
	for _, m := range ms {
		if n := User(m); n != nil {
			users = append(users, n)
		}
	}
	return users
}
