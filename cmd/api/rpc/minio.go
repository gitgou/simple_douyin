package rpc

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/url"
	"strings"
	"time"

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
		log.Fatalln("minio connect error", err)
		return 
	}
	//log.Printf("%#v\n", minioClient) // minioClient is now set up
	fmt.Println("minioClient is now set up")
	minioClient = client

	err = CreateBucket(constants.MinioBucketName)
	if err != nil {
		log.Panic("Create minio bucket error")
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
			log.Printf("bucket %s already exists", bucketName)
			return nil
		} 
			return err
	}

	log.Printf("bucket %s create successfully", bucketName)
	return nil
}

// FileUploader upload local file
func FileUploader(ctx context.Context, bucketName string, objectName string, filePath string) error {
	object, err := minioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{})
	if err != nil {
		log.Println("upload failed：", err)
		return err
	}
	log.Printf("Successfully uploaded %s of size %d\n", objectName, object.Size)
	return nil
}

// GetFileUrl get shareUrl of file from minio
func GetFileUrl(bucketName string, fileName string, expires time.Duration) (*url.URL, error) {
	ctx := context.Background()

	// get ip
	ip, err := GetOutBoundIP()
	if err != nil {
		fmt.Println(err)
		return nil, errno.ConvertErr(err)
	}
	
	fmt.Println(ip)
	minioEndPointAddr := ip + ":9000" // "172.26.41.217:9000"
	// Initialize minio client object.
	client, err := minio.New(minioEndPointAddr, &minio.Options{
		Creds:  credentials.NewStaticV4(constants.MinioAccessID, constants.MinioAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		log.Fatalln("minio connect error", err)
		return nil ,err
	}

	reqParams := make(url.Values)
	if expires <= 0 {
		expires = time.Second * 60 * 60 * 24
	}
	fileUrl, err := client.PresignedGetObject(ctx, bucketName, fileName, expires, reqParams)
	if err != nil {
		log.Printf("get url of file %s from bucket %s failed, %s", fileName, bucketName, err)
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
		log.Fatalln(err)
		return err
	}
	log.Println("Success")
	return nil
}

func GetOutBoundIP() (ip string, err error) {
	// 使用udp发起网络连接, 这样不需要关注连接是否可通, 随便填一个即可
	conn, err := net.Dial("udp", "8.8.8.8:5000")
	if err != nil {
		fmt.Println(err)
		return
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	// fmt.Println(localAddr.String())
	ip = strings.Split(localAddr.String(), ":")[0]
	return
}
