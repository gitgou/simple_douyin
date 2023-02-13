package main

import (
	"context"
	redisdemo "github.com/gitgou/simple_douyin/kitex_gen/redisdemo"
)

// RedisServiceImpl implements the last service interface defined in the IDL.
type RedisServiceImpl struct{}

// Set implements the RedisServiceImpl interface.
func (s *RedisServiceImpl) Set(ctx context.Context, req *redisdemo.SetRequest) (resp *redisdemo.SetResponse, err error) {
	// TODO: Your code here...
	resp = new(redisdemo.SetResponse) 

	return resp, nil
}
