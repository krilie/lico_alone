package third_api

import (
	"context"
	"io"
)

type OssQiNiu struct{}

func (o *OssQiNiu) UploadFile(ctx context.Context, name string, file io.ReadSeeker, size int64) (content, bucket, key string, err error) {
	panic("implement me")
}

func (o *OssQiNiu) DeleteFile(ctx context.Context, userId, key string) error {
	panic("implement me")
}

func (o *OssQiNiu) GetFullUrl(ctx context.Context, isPub bool, key string) (url string) {
	panic("implement me")
}

func (o *OssQiNiu) GetBaseUrl(ctx context.Context) string {
	panic("implement me")
}

func (o *OssQiNiu) GetBucketName(ctx context.Context) string {
	panic("implement me")
}
