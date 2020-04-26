package third_api

import (
	"context"
	"io"
)

type OssAliyun struct{}

func (o *OssAliyun) UploadFile(ctx context.Context, userId, name string, file io.ReadSeeker, size int64) (content, bucket, key string, err error) {
	panic("implement me")
}

func (o *OssAliyun) DeleteFile(ctx context.Context, userId, key string) error {
	panic("implement me")
}

func (o *OssAliyun) GetFullUrl(ctx context.Context, isPub bool, key string) (url string) {
	panic("implement me")
}

func (o *OssAliyun) GetBaseUrl(ctx context.Context) string {
	panic("implement me")
}

func (o *OssAliyun) GetBucketName(ctx context.Context) string {
	panic("implement me")
}
