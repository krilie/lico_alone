package third_api

import (
	"context"
	"github.com/krilie/lico_alone/common/config"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"io"
)

type OssQiNiu struct {
	AccessKey  string
	SecretKey  string
	BucketName string
	qboxMac    *qbox.Mac
}

func (o *OssQiNiu) UploadFile(ctx context.Context, fileName string, fileStream io.ReadSeeker, fileSize int64) (url, key string, err error) {
	panic("implement me")
}

func (o *OssQiNiu) DeleteFile(ctx context.Context, fileKey string) error {
	panic("implement me")
}

func (o *OssQiNiu) GetUrl(ctx context.Context, isPub bool, fileKey string) (url string, err error) {
	panic("implement me")
}

func (o *OssQiNiu) GetBaseUrl(ctx context.Context) string {
	panic("implement me")
}

func (o *OssQiNiu) GetBucketName(ctx context.Context) string {
	panic("implement me")
}

func NewOssQiNiu(cfg *config.FileSave) *OssQiNiu {
	return &OssQiNiu{
		AccessKey:  cfg.OssKey,
		SecretKey:  cfg.OssSecret,
		BucketName: cfg.OssBucket,
		qboxMac:    nil,
	}
}
