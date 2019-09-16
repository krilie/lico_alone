package oss_s3

import (
	"context"
	"github.com/krilie/lico_alone/common/config"
	"github.com/krilie/s-file"
	"io"
)

type OssLocal struct {
	File *s_file.SFile
}

func NewOssLocal(cfg config.Config) *OssLocal {
	return &OssLocal{File: s_file.NewSFile(cfg.FileSavePath)}
}

func (o *OssLocal) UploadFile(ctx context.Context, userId, name string, file io.ReadSeeker, size int64) (content string, bucket string, key string, err error) {
	content, key, err = o.File.SaveFile(ctx, name, file)
	return content, "", key, err
}

func (o *OssLocal) DeleteFile(ctx context.Context, userId, bucket, key string) error {
	return o.File.DeleteFile(ctx, bucket)
}

func (o *OssLocal) GetBucketName() string {
	return ""
}
