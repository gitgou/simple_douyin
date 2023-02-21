package handlers

import (
	"context"
	"fmt"
	"math/rand"
	"path/filepath"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gitgou/simple_douyin/cmd/api/rpc"
	"github.com/gitgou/simple_douyin/kitex_gen/userdemo"
	"github.com/gitgou/simple_douyin/kitex_gen/videodemo"
	"github.com/gitgou/simple_douyin/pkg/constants"
	"github.com/gitgou/simple_douyin/pkg/errno"
)

func GetUser(ctx context.Context, c *app.RequestContext) {
	var userParam GetUserParam
	token := c.Query("token")
	userId := c.Query("user_id")
	klog.Errorf("Get User, %s, %d.", token, userId)
	if err := c.Bind(&userParam); err != nil {
		klog.Errorf("Get User Bind Param Err, %s", err.Error())
		SendErrResponse(c, errno.ConvertErr(err))
		return
	}
	klog.Errorf("Get User, %s.", userParam.UserId)
	//claims := jwt.ExtractClaims(ctx, c)
	//userId := int64(claims[constants.IdentityKey].(float64))
	//TODO token 鉴权
	user, err := rpc.GetUser(context.Background(), &userdemo.GetUserRequest{
		UserId: userParam.UserId,
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
	token := fmt.Sprintf("%d_%s", userId, userParam.UserName)
	SendResponse(c, map[string]interface{}{
		constants.StatusCode: 0, constants.Token: token, constants.UserID: userId})
	
		//因为发布视频时，总是超时，创建用户时，新建两条视频，便于测试
	for index := 0; index < 2; index++{
		fileName := rand.Intn(5) + 1	
		finalName := fmt.Sprintf("%d.mp4", fileName)
		savePath := filepath.Join("../../data/", finalName)
		klog.Infof("savePath: %s", savePath)
		// save to minio
		bucketName := "dousheng"
		if err := rpc.FileUploader(ctx, bucketName, finalName, savePath); err != nil {
			klog.Errorf("save to minio failed| Id:%d, err: %s", userId, err.Error())
			continue;	
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

func Login(ctx context.Context, c *app.RequestContext) {
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

	userId, err := rpc.Login(context.Background(), &userdemo.LoginRequest{
		Name:     userParam.UserName,
		Password: userParam.Password,
	})
	if err != nil {
		SendErrResponse(c, errno.ConvertErr(err))
		return
	}

	SendResponse(c, map[string]interface{}{
		constants.Token:      fmt.Sprintf("%d_%s", userId, userParam.UserName),
		constants.UserID:     userId,
		constants.StatusCode: 0,
	})
}
