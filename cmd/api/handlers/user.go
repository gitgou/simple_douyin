package handlers

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"path/filepath"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gitgou/simple_douyin/cmd/api/rpc"
	"github.com/gitgou/simple_douyin/kitex_gen/userdemo"
	"github.com/gitgou/simple_douyin/kitex_gen/videodemo"
	"github.com/gitgou/simple_douyin/pkg/constants"
	"github.com/gitgou/simple_douyin/pkg/errno"
	"github.com/hertz-contrib/jwt"
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
				//"user_id":     8,
				"status_msg": "success",
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
				"status_code": code,
				"status_msg":  message,
			})
		},
	})
	if err != nil {
		panic(err)
	}
}

// TODO 登录之后需要调用 Login
func GetUser(ctx context.Context, c *app.RequestContext) {
	var userParam GetUserParam
	if err := c.Bind(&userParam); err != nil {
		klog.Errorf("Get User Bind Param Err, %s", err.Error())
		SendErrResponse(c, errno.ConvertErr(err))
		return
	}
	claims, _ := JwtMiddleware.GetClaimsFromJWT(ctx, c)
	userId := int64(claims[constants.IdentityKey].(float64))
	klog.Infof("Get User, %d.", userId)
	user, err := rpc.GetUser(context.Background(), &userdemo.GetUserRequest{
		UserId: userId,
	})
	if err != nil {
		klog.Errorf("Get User Bind Param Err, %s", err.Error())
		SendErrResponse(c, errno.ConvertErr(err))
		return
	}
	SendResponse(c, map[string]interface{}{
		constants.StatusCode: 0, constants.User: user})
}

func Register(ctx context.Context, c *app.RequestContext) {
	var userParam UserParam
	if err := c.Bind(&userParam); err != nil {
		klog.Error("Get Param ERR.", err)
		SendErrResponse(c, errno.ConvertErr(err))
		return
	}
	if len(userParam.UserName) == 0 || len(userParam.Password) == 0 {
		klog.Error("Get Param ERR.", errno.ParamErr, userParam.UserName, userParam.Password)
		SendErrResponse(c, errno.ParamErr)
		return
	}
	userId, err := rpc.CreateUser(context.Background(), &userdemo.CreateUserRequest{
		Name:     userParam.UserName,
		Password: userParam.Password,
	})
	if err != nil {
		SendErrResponse(c, errno.ConvertErr(err))
		return
	}
	//注册之后需要调用登录接口
	JwtMiddleware.LoginHandler(ctx, c)

	//测试代码
	//因为发布视频时，总是超时，创建用户时，新建两条视频，便于测试
	for index := 0; index < 2; index++ {
		fileName := rand.Intn(5) + 1
		finalName := fmt.Sprintf("%d.mp4", fileName)
		savePath := filepath.Join("../../data/", finalName)
		klog.Infof("savePath: %s", savePath)
		// save to minio
		bucketName := "dousheng"
		if err := rpc.FileUploader(ctx, bucketName, finalName, savePath); err != nil {
			klog.Errorf("save to minio failed| Id:%d, err: %s", userId, err.Error())
			continue
		}

		// get URL from minio
		url, err := rpc.GetFileUrl(bucketName, finalName, 0)
		if err != nil {
			klog.Errorf("get url fail| Id:%d, err: %s", userId, err.Error())
		} else {
			klog.Infof("User uploaded a file, %s", url)
		}

		if err := rpc.Publish(ctx, &videodemo.PublishRequest{Url: url.String(), Title: "test video", UserId: userId}); err != nil {
			klog.Infof("User uploaded a file, %s", url)
			continue
		}
	}
}
