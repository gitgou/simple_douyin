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
	"net/http"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gitgou/simple_douyin/kitex_gen/userdemo"
	"github.com/hertz-contrib/jwt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/gitgou/simple_douyin/cmd/api/handlers"
	"github.com/gitgou/simple_douyin/cmd/api/rpc"
	"github.com/gitgou/simple_douyin/pkg/constants"
	"github.com/gitgou/simple_douyin/pkg/errno"
	"github.com/gitgou/simple_douyin/pkg/tracer"
)

var (
	JwtMiddleware *jwt.HertzJWTMiddleware
)

func InitJwt() {
	var err error
	JwtMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Realm:         "test zone",
		Key:           []byte("secret key"),
		Timeout:       time.Hour * 12,
		MaxRefresh:    time.Hour,
		TokenLookup:   "header: Authorization, query: token, cookie: jwt, form: token",
		TokenHeadName: "Bearer",
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			klog.Errorf("Login Res: token:%s, code:%d", token, code)
			c.JSON(http.StatusOK, utils.H{
				"status_code": 0,
				"token":       token,
				"expire":      expire.Format(time.RFC3339),
				"user_id":     8,
				"status_msg":  "success",
			})
		},
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var loginStruct struct {
				Account  string `form:"username" json:"username" query:"username" vd:"(len($) > 0 && len($) < 30); msg:'Illegal format'"`
				Password string `form:"password" json:"password" query:"password" vd:"(len($) > 0 && len($) < 30); msg:'Illegal format'"`
			}
			if err := c.BindAndValidate(&loginStruct); err != nil {
				return nil, err
			}
			userId, err := rpc.Login(context.Background(), &userdemo.LoginRequest{Password: loginStruct.Password, Name: loginStruct.Account})
			if err != nil {
				return userId, err
			}

			return userId, nil
		},
		IdentityKey: constants.IdentityKey,
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return int64(claims[constants.IdentityKey].(float64))

		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					constants.IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			hlog.CtxErrorf(ctx, "jwt biz err = %+v", e.Error())
			return e.Error()
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(http.StatusOK, utils.H{
				"code":    code,
				"message": message,
			})
		},
	})
	if err != nil {
		panic(err)
	}
}

func Init() {
	tracer.InitJaeger(constants.ApiServiceName)
	rpc.InitRPC()
	InitJwt()
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
	apiRoute.POST("/user/login/", /*JwtMiddleware.LoginHandler*/ handlers.Login)
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
