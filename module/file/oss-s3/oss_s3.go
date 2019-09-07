package oss_s3

import (
	"context"
	"mime/multipart"
)

type FileOperator interface {
	UploadFile(ctx context.Context, userId, fileName, contentType string, file multipart.File, size int64) (objName string, err error)
	DeleteFile(ctx context.Context, userId, objKey string) error
	GetBucketName() string
}
