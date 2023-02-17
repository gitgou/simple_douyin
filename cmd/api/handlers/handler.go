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

package handlers

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/gitgou/simple_douyin/pkg/errno"
)

type Response struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

func SendErrResponse(c *app.RequestContext, err error) {
	Err := errno.ConvertErr(err)
	c.JSON(consts.StatusOK, Response{
		StatusCode: Err.ErrCode,
		StatusMsg:  Err.ErrMsg,
	})
}

func SendResponse(c *app.RequestContext, data map[string]interface{}) {
	c.JSON(consts.StatusOK, data)
}

type FeedRequest struct {
	LastestTime int64  `json:"lastest_time"`
	Token       string `json:"token"`
}

type GetUserParam struct {
	UserId int64 `json:"user_id"`
	Token  int64 `json:"token"`
}

type UserParam struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type ChatParam struct {
	Token    string `json:"token"`
	ToUserId int64  `json:"to_user_id"`
}

type ChatActionParam struct {
	Token      string `json:"token"`
	ToUserId   int64  `json:"to_user_id"`
	ActionType int64  `json:"action_type"`
	Content    string `json:"content"`
}

type PulishParam struct {
	Token string `json:"token"`
	//Data  byte   `json:"data"`
	Title string `json:"title"`
}

type PublishListParam struct {
	Token  string `json:"token,omitempty"`
	UserId int64  `json:"user_id,omitempty"`
}

type RelationParam struct {
	Token      string `json:"token,omitempty"`
	ToUserId   int64  `json:"to_user_id,omitempty"`
	ActionType int64  `json:"action_type,omitempty"`
}

type FollowListParam struct {
	Token  string `json:"token,omitempty"`
	UserId int64  `json:"user_id,omitempty"`
}

type FollowerListParam struct {
	Token  string `json:"token,omitempty"`
	UserId int64  `json:"user_id,omitempty"`
}

type FriendListParam struct {
	Token  string `json:"token,omitempty"`
	UserId int64  `json:"user_id,omitempty"`
}
