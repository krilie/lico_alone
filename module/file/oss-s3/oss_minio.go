package oss_s3

import (
	"context"
	"github.com/krilie/lico_alone/common/config"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/common/utils/id_util"
	"github.com/minio/minio-go"
	"mime/multipart"
)

type OssClient struct {
	Client     *minio.Client
	BucketName string
}

func NewOssClient(cfg config.Config) *OssClient {
	minioClient, err := minio.New(cfg.OssS3.OssEndPoint, cfg.OssS3.OssKey, cfg.OssS3.OssSecret, true) //endpoint, accessKeyID, secretAccessKey string, secure bool
	if err != nil {
		panic(errs.NewInternal().WithError(err))
	}
	return &OssClient{Client: minioClient, BucketName: cfg.OssS3.OssBucket}
}

func (f *OssClient) GetBucketName() string {
	return f.BucketName
}

func (f *OssClient) UploadFile(ctx context.Context, userId, fileName, contentType string, file multipart.File, size int64) (objName string, err error) {
	objName = id_util.GetUuid() + fileName
	userMate := make(map[string]string)
	userMate["user_id"] = userId
	n, err := f.Client.PutObject(f.BucketName, objName, file, size, minio.PutObjectOptions{ContentType: contentType, UserMetadata: userMate})
	if err != nil {
		_ = f.Client.RemoveIncompleteUpload(f.BucketName, objName) // 删除可能存在的不完整文件
		return "", errs.NewInternal().WithError(err)
	} else if n != size {
		_ = f.Client.RemoveIncompleteUpload(f.BucketName, objName) // 删除可能存在的不完整文件
		return "", errs.NewInternal().WithMsg("un completed upload please check")
	} else {
		return objName, nil
	}
}

func (f *OssClient) DeleteFile(ctx context.Context, userId, objKey string) error {
	err := f.Client.RemoveObject(f.BucketName, objKey)
	if err != nil {
		return errs.NewInternal().WithError(err)
	}
	return nil
}
