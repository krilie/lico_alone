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

	var fileCfg = ncfg.FileSave{}
	cfgs.GetConfigItem("file_save", &fileCfg)

	log = log.WithField(context_enum.Module.Str(), "module file service")
	var fileApi file_api.FileOperator

	if fileCfg.Channel == "local" {
		fileApi = file_api.NewLocalFileSave(fileCfg.OssBucket, fileCfg.OssEndPoint)
	} else if fileCfg.Channel == "qiniu" {
		fileApi = file_api.NewOssQiNiu(&fileCfg)
	} else if fileCfg.Channel == "minio" {
		fileApi = file_api.NewOssMinioClientByCfg(&fileCfg)
	} else {
		panic("config error on file save " + fileCfg.Channel)
	}
	return &FileModule{
		dao:     dao,
		log:     log,
		fileApi: fileApi,
	}
}
