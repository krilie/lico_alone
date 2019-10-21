package oss_s3

import (
	"context"
	"fmt"
	"github.com/krilie/lico_alone/common/config"
	"github.com/krilie/s-file"
	"io"
)

type OssLocal struct {
	File *s_file.SFile
	Url  string
}

func (o *OssLocal) GetFullUrl(ctx context.Context, isPub bool, key string) string {
	return fmt.Sprintf("%v/%v", o.Url, key)
}

func NewOssLocal(cfg config.Config) *OssLocal {
	return &OssLocal{File: s_file.NewSFile(cfg.FileSave.LocalFileSaveDir)}
}

func (o *OssLocal) UploadFile(ctx context.Context, userId, name string, file io.ReadSeeker, size int64) (content string, bucket string, key string, err error) {
	content, key, err = o.File.SaveFile(ctx, name, file)
	return content, "", key, err
}

func (o *OssLocal) DeleteFile(ctx context.Context, userId, key string) error {
	return o.File.DeleteFile(ctx, key)
}

func (o *OssLocal) GetBucketName() string {
	return ""
}

func (o *OssLocal) GetBaseUrl(ctx context.Context) string {
	return o.Url
}
