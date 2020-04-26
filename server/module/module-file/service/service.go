package service

import (
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/krilie/lico_alone/module/module-file/dao"
	oss_s3 "github.com/krilie/lico_alone/module/module-file/third-api"
)

type FileService struct {
	dao     *dao.FileDao
	log     *nlog.NLog
	fileApi oss_s3.FileOperator
}
