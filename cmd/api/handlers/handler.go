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

/*
	type Comment struct {
		Id         int64  `json:"id,omitempty"`
		User       User   `json:"user"`
		Content    string `json:"content,omitempty"`
		CreateDate string `json:"create_date,omitempty"`
	}

	type User struct {
		Id            int64  `json:"id,omitempty"`
		Name          string `json:"name,omitempty"`
		FollowCount   int64  `json:"follow_count,omitempty"`
		FollowerCount int64  `json:"follower_count,omitempty"`
		IsFollow      bool   `json:"is_follow,omitempty"`
	}

	type Message struct {
		Id         int64  `json:"id,omitempty"`
		Content    string `json:"content,omitempty"`
		CreateTime string `json:"create_time,omitempty"`
	}
type MessageSendEvent struct {
	UserId     int64  `json:"user_id,omitempty"`
	ToUserId   int64  `json:"to_user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}

type MessagePushEvent struct {
	FromUserId int64  `json:"user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}
*/

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

type PulishParam struct {
	Token string `json:"token"`
	//Data  byte   `json:"data"`
	Title string `json:"title"`
}

/*

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}


type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	Response
	Comment Comment `json:"comment,omitempty"`
}

type UserListResponse struct {
	Response
	UserList []User `json:"user_list"`
}
*/
