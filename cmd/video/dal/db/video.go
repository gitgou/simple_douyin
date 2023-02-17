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

type VideoModel struct {
	gorm.Model
	UserID   int64  `json:"user_id"`
	Title    string `json:"title"`
	PlayURL  string `json:"play_url"`
	CoverURL string `json:"cover_url"`
}

func (n *VideoModel) TableName() string {
	return constants.VideoTableName
}

// TODO 需要刷未被刷到的视频
//Feed Video to User
func FeedVideo(ctx context.Context) ([]*VideoModel, error) {
	res := make([]*VideoModel, 0)

	if err := DB.WithContext(ctx).Order(`create_at DESC`).Limit(10).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func PublishVideo(ctx context.Context, videoModel * VideoModel) error{
	return DB.WithContext(ctx).Create(videoModel).Error
}

func GetVideos(ctx context.Context, userId int64)([]*VideoModel, error){
	res := make([]*VideoModel, 0)
	if err := DB.WithContext(ctx).Where("user_id = ?", userId).Order(`create_at DESC`).Find(&res).Error; err != nil{
		return res, err
	}
	return res, nil
} 