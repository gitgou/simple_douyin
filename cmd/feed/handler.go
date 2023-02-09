package main

import (
	"context"
	demofeed "simple_douyin/kitex_gen/demofeed"
)

// FeedServiceImpl implements the last service interface defined in the IDL.
type FeedServiceImpl struct{}

// Feed implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) Feed(ctx context.Context, req *demofeed.FeedRequest) (resp *demofeed.FeedResponse, err error) {
	// TODO: Your code here...
	return
}
