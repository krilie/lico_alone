package file_api

import (
	"context"
	"fmt"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/common/utils/file_util"
	"github.com/krilie/lico_alone/common/utils/id_util"
	"github.com/krilie/lico_alone/component/ncfg"
	"github.com/minio/minio-go"
	"io"
)

type OssMinio struct {
	Client     *minio.Client
	BucketName string
	Url        string
}

func (o *OssMinio) DeleteFileByUrl(ctx context.Context, url string) error {
	panic("implement me")
}

func (f *OssMinio) GetUrl(ctx context.Context, isPub bool, fileKey string) (url string, err error) {
	return fmt.Sprintf("%v/%v/%v", f.Url, f.BucketName, fileKey), nil
}

func (f *OssMinio) GetBucketName(ctx context.Context) string {
	return f.BucketName
}

func (f *OssMinio) UploadFile(ctx context.Context, fileName string, fileStream io.Reader, fileSize int64) (url, key string, err error) {
	content, err := file_util.GetContentType(file)
	if err != nil {
		return "", "", err
	}
	key = id_util.GetUuid() + name
	userMate := make(map[string]string)
	userMate["user_id"] = "userId"
	n, err := f.Client.PutObject(f.BucketName, key, file, size, minio.PutObjectOptions{ContentType: content, UserMetadata: userMate})
	if err != nil {
		_ = f.Client.RemoveIncompleteUpload(f.BucketName, key) // 删除可能存在的不完整文件
		return f.BucketName, key, errs.NewInternal().WithError(err)
	} else if n != size {
		_ = f.Client.RemoveIncompleteUpload(f.BucketName, key) // 删除可能存在的不完整文件
		return f.BucketName, key, errs.NewInternal().WithMsg("un completed upload please check")
	} else {
		return f.BucketName, key, nil
	}
}

func (f *OssMinio) DeleteFile(ctx context.Context, key string) error {
	err := f.Client.RemoveObject(f.BucketName, key)
	if err != nil {
		return errs.NewInternal().WithError(err)
	}
	return nil
}

func (o *OssMinio) GetBaseUrl(ctx context.Context) string {
	return o.Url
}

func NewOssMinioClient(cfg *ncfg.FileSave) *OssMinio {
	minioClient, err := minio.New(cfg.OssEndPoint, cfg.OssKey, cfg.OssSecret, true) //endpoint, accessKeyID, secretAccessKey string, secure bool
	if err != nil {
		panic(errs.NewInternal().WithError(err))
	}
	url := fmt.Sprintf("%v%v%v", cfg.OssEndPoint, "/", cfg.OssBucket)
	return &OssMinio{Client: minioClient, BucketName: cfg.OssBucket, Url: url}
}
