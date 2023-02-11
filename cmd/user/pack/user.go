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

/*
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
*/
