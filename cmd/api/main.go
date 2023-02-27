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

package main

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/gitgou/simple_douyin/cmd/api/handlers"
	"github.com/gitgou/simple_douyin/cmd/api/rpc"
	"github.com/gitgou/simple_douyin/pkg/constants"
	"github.com/gitgou/simple_douyin/pkg/errno"
	"github.com/gitgou/simple_douyin/pkg/tracer"
)


func Init() {
	tracer.InitJaeger(constants.ApiServiceName)
	rpc.InitRPC()
	handlers.InitJwt()
}

func main() {
	r := server.Default(
		server.WithHostPorts("0.0.0.0:8080"),
		//server.WithHandleMethodNotAllowed(true),
	)
	Init()
	r.Use(recovery.Recovery(recovery.WithRecoveryHandler(
		func(ctx context.Context, c *app.RequestContext, err interface{}, stack []byte) {
			hlog.SystemLogger().CtxErrorf(ctx, "[Recovery] err=%v\nstack=%s", err, stack)
			c.JSON(consts.StatusInternalServerError, map[string]interface{}{
				"status_code":    errno.ServiceErrCode,
				"status_msg": fmt.Sprintf("[Recovery] err=%v\nstack=%s", err, stack),
			})
		})))
	apiRoute := r.Group("/douyin")

	apiRoute.GET("/feed/", handlers.Feed)
	//user
	apiRoute.POST("/user/login/", handlers.JwtMiddleware.LoginHandler)
	apiRoute.POST("/user/register/", handlers.Register)
	apiRoute.GET("/user/", /*JwtMiddleware.MiddlewareFunc(),*/ handlers.GetUser)

	pubRoute := apiRoute.Group("/publish", /*JwtMiddleware.MiddlewareFunc()*/)
	pubRoute.POST("/action/", handlers.Publish)
	pubRoute.GET("/list/", handlers.PublishList)

	chatRoute := apiRoute.Group("/message", /*JwtMiddleware.MiddlewareFunc()*/)
	chatRoute.GET("/chat/", handlers.GetChat)
	chatRoute.POST("/action/", handlers.ChatAction)

	//interaction
	apiRoute.POST("/favorite/action/", handlers.FavoriteAction)
	apiRoute.GET("/favorite/list/", handlers.FavoriteList)
	apiRoute.POST("/comment/action/", handlers.CommentAction)
	apiRoute.GET("/comment/list/", handlers.CommentList)

	//relation
	relationRoute := apiRoute.Group("/relation", /*JwtMiddleware.MiddlewareFunc()*/)
	relationRoute.POST("/action/", handlers.Relation)
	//关注列表
	relationRoute.GET("/follow/list/", handlers.FollowList) //TODO
	relationRoute.GET("/follower/list/", handlers.FollowerList)
	relationRoute.GET("/friend/list/", handlers.FriendList)

	r.NoRoute(func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "no route")
	})
	r.NoMethod(func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "no method")
	})

	r.Spin()
}
