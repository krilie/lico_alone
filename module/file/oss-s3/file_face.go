package oss_s3

import (
	"context"
	"io"
)

type FileOperator interface {
	UploadFile(ctx context.Context, userId, name string, file io.ReadSeeker, size int64) (content, bucket, key string, err error)
	DeleteFile(ctx context.Context, userId, key string) error
	GetFullUrl(ctx context.Context, isPub bool, key string) (url string)
	GetBucketName() string
}
