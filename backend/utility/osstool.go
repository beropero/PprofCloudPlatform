package utility

import (
	"context"
	"path/filepath"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gogf/gf/v2/frame/g"
)

// GetConfig 获取OSS的配置信息
func GetConfig(ctx context.Context) (endpoint, accessKeyId, accessKeySecret, bucketName string, err error) {
	endpointVar, err := g.Cfg().Get(ctx, "oss.endpoint")
	if err != nil {
		return
	}
	accessKeyIdVar, err := g.Cfg().Get(ctx, "oss.accessKeyId")
	if err != nil {
		return
	}
	accessKeySecretVar, err := g.Cfg().Get(ctx, "oss.accessKeySecret")
	if err != nil {
		return
	}
	bucketNameVar, err := g.Cfg().Get(ctx, "oss.bucketName")
	if err != nil {
		return
	}
	return endpointVar.String(), accessKeyIdVar.String(), accessKeySecretVar.String(), bucketNameVar.String(), nil
}

// GetOssBucket 获取Oss存储空间
func GetOssBucket(ctx context.Context) (*oss.Bucket, error) {
	// 获取OSS的配置信息。
	endpoint, accessKeyId, accessKeySecret, bucketName, err := GetConfig(ctx)
	if err != nil {
		return nil, err
	}
	// 创建OSSClient实例。
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		return nil, err
	}
	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return nil, err
	}
	return bucket, nil
}

// UploadFileOss 上传文件到OSS
func UploadFileOss(ctx context.Context, FileName string) (err error) {
	// 拼装文件路径。
	tempDir := g.Cfg().MustGet(ctx, "server.tempDir").String()
	localFileName := filepath.Join(tempDir, FileName)
	// 上传文件。
	bucket, err := GetOssBucket(ctx)
	if err != nil {
		return
	}
	err = bucket.PutObjectFromFile(FileName, localFileName)
	return
}

// DownloadFileOss 从OSS下载文件
func DownloadFileOss(ctx context.Context, objectName string, localFileName string) (err error) {
	// 下载文件。
	bucket, err := GetOssBucket(ctx)
	if err != nil {
		return
	}
	err = bucket.GetObjectToFile(objectName, localFileName)
	return
}
