package db

import (
	"context"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gitgou/simple_douyin/pkg/constants"
	"github.com/gitgou/simple_douyin/pkg/errno"
	"gorm.io/gorm"
)

type UserModel struct {
	ID              int64     `json:"id"`
	Name            string    `json:"name"`
	AvatarUrl       string    `json:"avatar_url"`
	Password        string    `json:"password"`
	BackgroundImage string    `json:"background_image"`
	Signature       string    `json:"signature"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	DeletedAt        gorm.DeletedAt 
}

func (n *UserModel) TableName() string {
	return constants.UserTableName
}

func GetUser(ctx context.Context, userId int64) (*UserModel, error) {
	var userModel UserModel
	if err := DB.WithContext(ctx).Where("id = ?", userId).Find(&userModel).Error; err != nil {
		return nil, err
	}
	return &userModel, nil
}

func CreateUsers(ctx context.Context, users []*UserModel) error {
	return DB.WithContext(ctx).Create(users).Error
}

func CreateUser(ctx context.Context, user *UserModel) (int64, error) {
	if user == nil{
		klog.Errorf("create user nil. ")
		return 0, errno.ParamErr
	}
	result := DB.WithContext(ctx).Create(user)
	return user.ID, result.Error
}
func QueryUser(ctx context.Context, name string) ([]*UserModel, error) {
	res := make([]*UserModel, 0)
	if err := DB.WithContext(ctx).Where("name = ?", name).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// MGetUsers multiple get list of user info
func MGetUsers(ctx context.Context, userIDs []int64) ([]*UserModel, error) {
	res := make([]*UserModel, 0)
	if len(userIDs) == 0 {
		return res, nil
	}

	if err := DB.WithContext(ctx).Where("id in ?", userIDs).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}


func UpdateUsers(userModels []*UserModel) error {
	return DB.WithContext(context.Background()).Save(&userModels).Error
}
