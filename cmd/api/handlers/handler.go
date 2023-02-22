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

type LoginResponse struct{
	Response
	UserId int64 `json:"user_id,omitempty"`
	Token string `json:"token"`
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
	LastestTime int64  `query:"lastest_time"`
	Token       string `query:"token"`
}

type GetUserParam struct {
	UserId int64 `query:"user_id"`
	Token  string `query:"token"`
}

type UserParam struct {
	UserName string  `query:"username"`
	Password string `query:"password"`
}

type ChatParam struct {
	Token    string `query:"token"`
	ToUserId int64  `query:"to_user_id"`
}

type ChatActionParam struct {
	Token      string `query:"token"`
	ToUserId   int64  `query:"to_user_id"`
	ActionType int64  `query:"action_type"`
	Content    string `query:"content"`
}

type PulishParam struct {
	Token string `query:"token"`
	//Data  byte   `json:"data"`
	Title string `query:"title"`
}

type PublishListParam struct {
	Token  string `query:"token"`
	UserId int64  `query:"user_id"`
}

type RelationParam struct {
	Token      string `query:"token"`
	ToUserId   int64  `query:"to_user_id"`
	ActionType int64  `query:"action_type"`
}

type FollowListParam struct {
	Token  string `query:"token"`
	UserId int64  `query:"user_id"`
}

type FollowerListParam struct {
	Token  string `query:"token"`
	UserId int64  `query:"user_id"`
}

type FriendListParam struct {
	Token  string `query:"token"`
	UserId int64  `query:"user_id"`
}

type FavoriteParam struct {
	Token 		string 	`query:"token"`
	VideoId		string	`query:"video_id`
	ActionType	string	`query:"action_type`
}

type FavoriteListParam struct {
	Token 		string 	`query:"token"`
	UserId		string	`query:"user_id`
}

type CommentParam struct {
	Token 		string 	`query:"token"`
	VideoId		string	`query:"video_id`
	ActionType	string	`query:"action_type"`
	CommentText	string	`query:"comment_text"`
	CommentId	string	`query:"comment_id"`
}

type CommentListParam struct {
	Token 		string 	`query:"token"`
	VideoId		string	`query:"video_id`
}
