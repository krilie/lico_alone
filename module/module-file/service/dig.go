package service

import (
	"github.com/krilie/lico_alone/module/module-file/dao"
)

var DigModuleFileProviderAll = []interface{}{
	dao.NewFileDao,
	NewFileModule,
}
