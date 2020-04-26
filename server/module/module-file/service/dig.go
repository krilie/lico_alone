package service

import (
	"github.com/krilie/lico_alone/common/dig"
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/krilie/lico_alone/module/module-file/dao"
)

func init() {
	dig.Container.MustProvide(func(dao *dao.FileDao, log *nlog.NLog) *FileService {
		return NewFileService(dao, log)
	})
}
