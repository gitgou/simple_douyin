// Copyright 2021 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package db

import (
	"context"

	"github.com/gitgou/simple_douyin/pkg/constants"

	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	AvatarUrl string `json:"avatar_url"`
	Password  string `json:"password"`
}

func (n *UserModel) TableName() string {
	return constants.VideoTableName
}

func GetUser(ctx context.Context, userId int64) (*UserModel, error) {
	var userModel UserModel
	if err := DB.WithContext(ctx).Where("id =", userId).Find(&userModel).Error; err != nil {
		return nil, err
	}
	return &userModel, nil
}

func CreateUsers(ctx context.Context, users []*UserModel) error {
	return DB.WithContext(ctx).Create(users).Error
}

func CreateUser(ctx context.Context, user *UserModel) (int64, error) {
	result := DB.WithContext(ctx).Create(user)
	return user.ID, result.Error
}
func QueryUser(ctx context.Context, name string) ([]*UserModel, error) {
	res := make([]*UserModel, 0)
	if err := DB.WithContext(ctx).Where("user_name = ?", name).Find(&res).Error; err != nil {
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

//TODO update
