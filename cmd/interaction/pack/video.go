package pack

import (
	"context"

	"github.com/gitgou/simple_douyin/cmd/interaction/dal/db"
	"github.com/gitgou/simple_douyin/cmd/interaction/rpc"
	"github.com/gitgou/simple_douyin/kitex_gen/videodemo"
)

func Videos(favoriteList []*db.FavoriteModel) []*videodemo.Video {
	if favoriteList == nil || len(favoriteList) == 0 {
		return nil
	}
	
	videoIds := make([]int64, 0, len(favoriteList))
	for _, m := range favoriteList {
		videoIds = append(videoIds, m.VideoId)
	}

	return rpc.GetVideoList(context.Background(), &videodemo.GetVideoListRequest{VideoId: videoIds})
}
