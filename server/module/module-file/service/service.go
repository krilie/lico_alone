package service

import (
	"context"
	"github.com/krilie/lico_alone/common/config"
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/krilie/lico_alone/module/module-file/dao"
	"github.com/krilie/lico_alone/module/module-file/file-api"
	"io"
)

type IFileService interface {
	UploadFile(ctx context.Context, userId, fileName string, file io.ReadSeeker, size int) (url, bucket, key string, err error)
	DeleteFile(ctx context.Context, bucket, key string) (err error)
	GetBaseUrl(ctx context.Context) string
}

type FileService struct {
	dao     *dao.FileDao
	log     *nlog.NLog
	fileApi file_api.FileOperator
}

func NewFileService(dao *dao.FileDao, log *nlog.NLog, cfgs *config.Config) *FileService {
	var fileApi file_api.FileOperator
	cfg := &cfgs.FileSave
	if cfg.SaveType == "local" {
		fileApi = file_api.NewLocalFileSave(cfg.LocalFileSaveDir, cfg.LocalFileSaveUrl)
	} else if cfg.SaveType == "qiniuOss" {
		fileApi = file_api.NewOssQiNiu(cfgs)
	} else {
		panic("config error on file save " + cfg.SaveType)
	}
	return &FileService{
		dao:     dao,
		log:     log,
		fileApi: fileApi,
	}
}
