package oss_s3

import (
	"context"
	"github.com/krilie/lico_alone/common/config"
	"github.com/krilie/s-file"
	"mime/multipart"
)

type OssLocal struct {
	File *s_file.SFile
}

func NewOssLocal(cfg config.Config) *OssLocal {
	return &OssLocal{File: s_file.NewSFile(cfg.FileSavePath)}
}

func (o *OssLocal) UploadFile(ctx context.Context, userId, fileName, contentType string, file multipart.File, size int64) (objName string, err error) {
	_, key, err := o.File.SaveFile(ctx, fileName, file)
	return key, err
}

func (o *OssLocal) DeleteFile(ctx context.Context, userId, objKey string) error {
	return o.File.DeleteFile(ctx, objKey)
}

func (o *OssLocal) GetBucketName() string {
	return ""
}
