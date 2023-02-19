package rpc

import (
	"context"
	"net"
	"net/url"
	"strings"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gitgou/simple_douyin/pkg/constants"
	"github.com/gitgou/simple_douyin/pkg/errno"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var minioClient *minio.Client

func initMinio() {
	// Initialize minio client 
	client, err := minio.New(constants.MinioEndPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(constants.MinioAccessID, constants.MinioAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		klog.Fatal("minio connect error", err)
		return 
	}
	//log.Printf("%#v\n", minioClient) // minioClient is now set up
	klog.Info("minioClient is now set up")
	minioClient = client

	err = CreateBucket(constants.MinioBucketName)
	if err != nil {
		klog.Fatalf("Create minio bucket error, %s", err)
		return
	}
}

func CreateBucket(bucketName string) error {
	if len(bucketName) <= 0 {
		return errno.ParamErr 
	}
	ctx := context.Background()
	err := minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: constants.Location})
	if err != nil {
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			klog.Fatal("bucket %s already exists", bucketName)
			return nil
		} 
			return err
	}

	klog.Info("bucket %s create successfully", bucketName)
	return nil
}

// FileUploader upload local file
func FileUploader(ctx context.Context, bucketName string, objectName string, filePath string) error {
	object, err := minioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{})
	if err != nil {
		klog.Error("upload failed：", err)
		return err
	}
	klog.Info("Successfully uploaded %s of size %d\n", objectName, object.Size)
	return nil
}

// GetFileUrl get shareUrl of file from minio
func GetFileUrl(bucketName string, fileName string, expires time.Duration) (*url.URL, error) {
	ctx := context.Background()

	// get ip
	ip, err := GetOutBoundIP()
	if err != nil {
		klog.Errorf("Get FileUrl Fail,%s", err)
		return nil, errno.ConvertErr(err)
	}
	
	klog.Infof("ip: %s ", ip)
	minioEndPointAddr := ip + ":9000" // "172.26.41.217:9000"
	// Initialize minio client object.
	client, err := minio.New(minioEndPointAddr, &minio.Options{
		Creds:  credentials.NewStaticV4(constants.MinioAccessID, constants.MinioAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		klog.Error("minio connect error", err)
		return nil ,err
	}

	reqParams := make(url.Values)
	if expires <= 0 {
		expires = time.Second * 60 * 60 * 24
	}
	fileUrl, err := client.PresignedGetObject(ctx, bucketName, fileName, expires, reqParams)
	if err != nil {
		klog.Errorf("get url of file %s from bucket %s failed, %s", fileName, bucketName, err)
		return nil, err
	}
	// TODO: url is quite long, or need to shorten
	return fileUrl, nil
}

func RemoveFile(bucketName string, objectName string) error {
	opts := minio.RemoveObjectOptions{
		GovernanceBypass: true,
	}
	err := minioClient.RemoveObject(context.Background(), bucketName, objectName, opts)
	if err != nil {
		klog.Errorf("RemoveFile, %s", err)
		return err
	}
	klog.Info("Success")
	return nil
}

func GetOutBoundIP() (ip string, err error) {
	// 使用udp发起网络连接, 这样不需要关注连接是否可通, 随便填一个即可
	conn, err := net.Dial("udp", "8.8.8.8:5000")
	if err != nil {
		klog.Errorf("GetOutBoundIP: %s", err)
		return
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	// fmt.Println(localAddr.String())
	ip = strings.Split(localAddr.String(), ":")[0]
	return
}
