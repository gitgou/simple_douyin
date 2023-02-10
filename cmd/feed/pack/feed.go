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

package pack

import (
	"github.com/gitgou/simple_douyin/cmd/feed/dal/db"
	"github.com/gitgou/simple_douyin/kitex_gen/demofeed"
)

func Video(m *db.VideoModel) *demofeed.Video {
	if m == nil {
		return nil
	}

	return &demofeed.Video{
		Id:       m.ID,
		CoverURL: m.CoverURL,
		PlayURL:  m.PlayURL,
		Title:    m.Title,
	}
}

func Videos(ms []*db.VideoModel) []*demofeed.Video {
	if ms == nil || len(ms) == 0 {
		return nil
	}
	videos := make([]*demofeed.Video, 0)
	for _, m := range ms {
		if n := Video(m); n != nil {
			videos = append(videos, n)
		}
	}
	return videos
}

func UserIds(ms []*db.VideoModel) []int64 {
	uIds := make([]int64, 0)
	if len(ms) == 0 {
		return uIds
	}
	uIdMap := make(map[int64]struct{})
	for _, m := range ms {
		if m != nil {
			uIdMap[m.UserID] = struct{}{}
		}
	}

	for uId := range uIdMap {
		uIds = append(uIds, uId)
	}
	return uIds
}
