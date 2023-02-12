package handlers

import (
	"context"
	"fmt"
	"log"
	"path/filepath"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/gitgou/simple_douyin/cmd/api/rpc"
	"github.com/gitgou/simple_douyin/kitex_gen/videodemo"
	"github.com/gitgou/simple_douyin/pkg/constants"
	"github.com/gitgou/simple_douyin/pkg/errno"
	"github.com/hertz-contrib/jwt"
)

func Publish(ctx context.Context, c *app.RequestContext) {
	var publishVar PulishParam
	if err := c.Bind(&publishVar); err != nil {
		SendErrResponse(c, errno.ConvertErr(err))
		return
	}

	if len(publishVar.Token) == 0 || len(publishVar.Title) == 0 {
		SendErrResponse(c, errno.ParamErr)
	}
	claims := jwt.ExtractClaims(ctx, c)
	userId := int64(claims[constants.IdentityKey].(float64))
	fmt.Println("get userID, ", userId);
	//TODO video byte to url
	data, err := c.FormFile("data")
	if err != nil {
		SendErrResponse(c, errno.ConvertErr(err))
		return
	}

	fileName := filepath.Base(data.Filename)
	finalName := fmt.Sprintf("%s_%s_%s",  strconv.FormatInt(userId, 16), publishVar.Title, fileName)
	savePath := filepath.Join("../../../static/", finalName)

	// save to local
	if err := c.SaveUploadedFile(data, savePath); err != nil {
		SendErrResponse(c, errno.ConvertErr(err))
		return
	}

	// save to minio
	bucketName := "dousheng"
	if err := rpc.FileUploader(ctx, bucketName, finalName, savePath); err != nil {
		fmt.Println("save to minio failed")
		SendErrResponse(c, errno.ConvertErr(err))
		return
	}

	// get URL from minio
	url, err := rpc.GetFileUrl(bucketName, finalName, 0)
	if err != nil {
		log.Printf("get url failed")
	} else {
		log.Printf("User uploaded a file, %s", url)
	}

	if err := rpc.Publish(ctx, &videodemo.PublishRequest{ Url: url.String(), Title:   publishVar.Title,UserId:  userId}); err != nil {
		SendErrResponse(c, errno.ConvertErr(err))
		return
	}

	SendResponse(c, map[string]interface{}{constants.StatusCode: 0, constants.StatusMsg:"upload successful."})
	return ;
}

func PublishList(ctx context.Context, c *app.RequestContext) {

}
