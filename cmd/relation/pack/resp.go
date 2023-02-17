package pack

import (
	"errors"
	"time"

	"github.com/gitgou/simple_douyin/kitex_gen/userdemo"
	"github.com/gitgou/simple_douyin/pkg/errno"
)

// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *userdemo.BaseResp {
	if err == nil {
		return baseResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return baseResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return baseResp(s)
}

func baseResp(err errno.ErrNo) *userdemo.BaseResp {
	return &userdemo.BaseResp{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg, ServiceTime: time.Now().Unix()}
}
