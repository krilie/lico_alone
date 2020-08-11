package service

import (
	"context"
	context_enum "github.com/krilie/lico_alone/common/com-model/context-enum"
	"github.com/krilie/lico_alone/component/ncfg"
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/krilie/lico_alone/module/module-file/dao"
	"github.com/krilie/lico_alone/module/module-file/file-api"
	"io"
)

type IFileService interface {
	UploadFile(ctx context.Context, userId, fileName string, file io.Reader, size int) (url, bucket, key string, err error)
	DeleteFile(ctx context.Context, bucket, key string) (err error)
	GetBaseUrl(ctx context.Context) string
}

type FileModule struct {
	dao     *dao.FileDao
	log     *nlog.NLog
	fileApi file_api.FileOperator
}

func NewFileModule(dao *dao.FileDao, log *nlog.NLog, cfgs *ncfg.NConfig) *FileModule {
	log = log.WithField(context_enum.Module.Str(), "module file service")
	var fileApi file_api.FileOperator
	cfg := &cfgs.Cfg.FileSave
	if cfg.Channel == "local" {
		fileApi = file_api.NewLocalFileSave(cfg.OssBucket, cfg.OssEndPoint)
	} else if cfg.Channel == "qiniu" {
		fileApi = file_api.NewOssQiNiu(cfg)
	} else if cfg.Channel == "minio" {
		fileApi = file_api.NewOssMinioClient(cfg)
	} else {
		panic("config error on file save " + cfg.Channel)
	}
	return &FileModule{
		dao:     dao,
		log:     log,
		fileApi: fileApi,
	}
}
