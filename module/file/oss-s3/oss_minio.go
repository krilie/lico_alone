package oss_s3

import (
	"context"
	"github.com/krilie/lico_alone/common/config"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/common/utils/file_util"
	"github.com/krilie/lico_alone/common/utils/id_util"
	"github.com/minio/minio-go"
	"io"
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

func (f *OssClient) UploadFile(ctx context.Context, userId, name string, file io.ReadSeeker, size int64) (content string, bucket string, key string, err error) {
	content, err = file_util.GetContentType(file)
	if err != nil {
		return "", "", "", err
	}
	key = id_util.GetUuid() + name
	userMate := make(map[string]string)
	userMate["user_id"] = userId
	n, err := f.Client.PutObject(f.BucketName, key, file, size, minio.PutObjectOptions{ContentType: content, UserMetadata: userMate})
	if err != nil {
		_ = f.Client.RemoveIncompleteUpload(f.BucketName, key) // 删除可能存在的不完整文件
		return content, f.BucketName, key, errs.NewInternal().WithError(err)
	} else if n != size {
		_ = f.Client.RemoveIncompleteUpload(f.BucketName, key) // 删除可能存在的不完整文件
		return content, f.BucketName, key, errs.NewInternal().WithMsg("un completed upload please check")
	} else {
		return content, f.BucketName, key, nil
	}
}

func (f *OssClient) DeleteFile(ctx context.Context, userId, bucket, key string) error {
	err := f.Client.RemoveObject(bucket, key)
	if err != nil {
		return errs.NewInternal().WithError(err)
	}
	return nil
}
