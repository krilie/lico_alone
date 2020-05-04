package file_api

import (
	"context"
	"io"
)

type FileOperator interface {
	UploadFile(ctx context.Context, fileName string, fileStream io.ReadSeeker, fileSize int64) (url, key string, err error)
	DeleteFile(ctx context.Context, fileKey string) error
	DeleteFileByUrl(ctx context.Context, url string) error
	GetUrl(ctx context.Context, isPub bool, fileKey string) (url string, err error)
	GetBaseUrl(ctx context.Context) string
	GetBucketName(ctx context.Context) string
}
