package cache

//TODO store user with map
import (
	"log"

	"github.com/gitgou/simple_douyin/cmd/user/dal/db"
)
type User struct{
	User db.UserModel;
	IsOnline bool;
}
var MapLoginUser map[string]User //key : token

var MapUser map[int64]User // key : userId


func Login(token string, user db.UserModel){
	MapLoginUser[token] = User{
		User: user,
		IsOnline: true,
	}
	MapUser[user.ID] = User{
		User: user,
		IsOnline: true,
	}	
}

func StoreDB(){
	userModels := make([]*db.UserModel, 0, len(MapUser))
	var index = 0
	for _, v := range MapUser {
		userModels[index] = &v.User;
	}
	if err := db.UpdateUsers(userModels); err != nil{
		log.Println("Store DB Fail. ");
	}
}