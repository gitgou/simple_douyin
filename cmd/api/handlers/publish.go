package handlers

import (
	"context"
	"fmt"
	"path/filepath"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gitgou/simple_douyin/cmd/api/rpc"
	"github.com/gitgou/simple_douyin/kitex_gen/videodemo"
	"github.com/gitgou/simple_douyin/pkg/constants"
	"github.com/gitgou/simple_douyin/pkg/errno"
	"github.com/gitgou/simple_douyin/pkg/utils"
)

func Publish(ctx context.Context, c *app.RequestContext) {

	klog.Infof("Publish | test, %s.")
	var publishVar PulishParam
	if err := c.Bind(&publishVar); err != nil {
		SendErrResponse(c, errno.ConvertErr(err))
		return
	}
	klog.Infof("Publish | test, %s.", publishVar)
	if len(publishVar.Token) == 0 || len(publishVar.Title) == 0 {
		SendErrResponse(c, errno.ParamErr)
	}

	userId := utils.GetUserIdInToken(publishVar.Token)
	klog.Infof("test| Publish userId: %d", userId)
	//claims := jwt.ExtractClaims(ctx, c)
	//userId := int64(claims[constants.IdentityKey].(float64))
	//TODO video byte to url
	data, err := c.FormFile("data")
	if err != nil {
		klog.Errorf("Get Data Fail. ID:%d, err:%s", userId, err.Error())
		SendErrResponse(c, errno.ConvertErr(err))
		return
	}

	fileName := filepath.Base(data.Filename)
	finalName := fmt.Sprintf("%s_%s_%s", strconv.FormatInt(userId, 16), publishVar.Title, fileName)
	savePath := filepath.Join("../../../data/", finalName)

	// save to local
	if err := c.SaveUploadedFile(data, savePath); err != nil {
		klog.Errorf("Save File Fail. ID:%d, err:%s", userId, err.Error())
		SendErrResponse(c, errno.ConvertErr(err))
		return
	}

	// save to minio
	bucketName := "dousheng"
	if err := rpc.FileUploader(ctx, bucketName, finalName, savePath); err != nil {
		klog.Errorf("save to minio failed| Id:%d, err: %s", userId, err.Error())
		SendErrResponse(c, errno.ConvertErr(err))
		return
	}

	// get URL from minio
	url, err := rpc.GetFileUrl(bucketName, finalName, 0)
	if err != nil {
		klog.Errorf("get url fail| Id:%d, err: %s", userId, err.Error())
	} else {
		klog.Infof("User uploaded a file, %s", url)
	}

	if err := rpc.Publish(ctx, &videodemo.PublishRequest{Url: url.String(), Title: publishVar.Title, UserId: userId}); err != nil {
		klog.Infof("User uploaded a file, %s", url)
		SendErrResponse(c, errno.ConvertErr(err))
		return
	}

	SendResponse(c, map[string]interface{}{constants.StatusCode: 0, constants.StatusMsg: "upload successful."})
	return
}

func PublishList(ctx context.Context, c *app.RequestContext) {
	var publishListVar PublishListParam
	if err := c.Bind(&publishListVar); err != nil {
		SendErrResponse(c, errno.ConvertErr(err))
		return
	}
	//claims := jwt.ExtractClaims(ctx, c)
	//userID := int64(claims[constants.IdentityKey].(float64))
	klog.Infof("test publishList, %d", publishListVar.UserId)
	videoList, err := rpc.GetPublishList(context.Background(), &videodemo.PublishListRequest{
		UserId: publishListVar.UserId,
	})
	if err != nil {
		klog.Errorf(" publishList, userId:%d, Err:%s ", publishListVar.UserId, err.Error())
		SendErrResponse(c, errno.ConvertErr(err))
		return
	}
	SendResponse(c, map[string]interface{}{
		constants.StatusCode: 0,
		constants.VideoList:  videoList,
	})

}
