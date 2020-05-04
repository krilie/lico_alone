package file_api

import (
	"context"
	"io"
)

type OssAliyun struct{}

func (o *OssAliyun) DeleteFileByUrl(ctx context.Context, url string) error {
	panic("implement me")
}

func (o *OssAliyun) UploadFile(ctx context.Context, fileName string, fileStream io.ReadSeeker, fileSize int64) (url, key string, err error) {
	panic("implement me")
}

func (o *OssAliyun) DeleteFile(ctx context.Context, fileKey string) error {
	panic("implement me")
}

func (o *OssAliyun) GetUrl(ctx context.Context, isPub bool, fileKey string) (url string, err error) {
	panic("implement me")
}

func (o *OssAliyun) GetBaseUrl(ctx context.Context) string {
	panic("implement me")
}

func (o *OssAliyun) GetBucketName(ctx context.Context) string {
	panic("implement me")
}
