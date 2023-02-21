package cache

//TODO store user with map
import (
	"sync"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gitgou/simple_douyin/cmd/user/dal/db"
)

type User struct {
	User     db.UserModel
	IsOnline bool
}

var (
	MapLoginUser map[string]int64 //key : token

	MapUser map[int64]User // key : userId

	MutexUser sync.Mutex
)

func Login(user db.UserModel) {
	MutexUser.Lock()
	defer MutexUser.Unlock()
	MapUser[user.ID] = User{
		User:     user,
		IsOnline: true,
	}
	MapLoginUser[user.Name] = user.ID
}

func GetUserById(userId int64) *User {
	if user, exist := MapUser[userId]; exist {
		return &user
	}
	return nil
}

func StoreDB() {
	MutexUser.Lock()
	defer MutexUser.Unlock()
	klog.Info("map user: ", len(MapUser))
	userModels := make([]*db.UserModel, 0, len(MapUser))
	for _, v := range MapUser {
		userModels = append(userModels, &v.User)
	}
	if len(userModels) == 0 {
		return
	}
	if err := db.UpdateUsers(userModels); err != nil {
		klog.Errorf("Store DB Fail. %s", err.Error())
	}
}
//TODO 登出时删除内存用户