package pack

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gitgou/simple_douyin/cmd/video/dal/db"
	"github.com/gitgou/simple_douyin/cmd/video/rpc"
	"github.com/gitgou/simple_douyin/kitex_gen/userdemo"
	"github.com/gitgou/simple_douyin/kitex_gen/videodemo"
)

func Video(m *db.VideoModel, user *userdemo.User) *videodemo.Video {
	if m == nil {
		return nil
	}

	return &videodemo.Video{
		Id:       int64(m.ID),
		CoverUrl: m.CoverURL,
		PlayUrl:  m.PlayURL,
		Title:    m.Title,
		Author:   user,
	}
}

func Videos(ms []*db.VideoModel) []*videodemo.Video {
	if ms == nil || len(ms) == 0 {
		return nil
	}

	//get UserInfo
	uIds := UserIds(ms)
	users, err := rpc.MGetUser(context.Background(), &userdemo.MGetUserRequest{UserIds: uIds})
	if err != nil {
		klog.Error("MGet User Err. ", err.Error())
	}

	videos := make([]*videodemo.Video, 0)
	for i, m := range ms {
		if n := Video(m, users[uIds[i]]); n != nil {
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
