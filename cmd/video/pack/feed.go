package pack

import (
	"github.com/gitgou/simple_douyin/cmd/video/dal/db"
	"github.com/gitgou/simple_douyin/kitex_gen/demofeed"
)

func Video(m *db.VideoModel) *demofeed.Video {
	if m == nil {
		return nil
	}

	return &demofeed.Video{
		Id:       m.ID,
		CoverUrl: m.CoverURL,
		PlayUrl:  m.PlayURL,
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
