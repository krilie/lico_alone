package service

import (
	"github.com/krilie/lico_alone/module/module-config/dao"
)

// DigModuleConfigProviderAll provider
var DigModuleConfigProviderAll = []interface{}{
	dao.NewConfigDao,
	NewConfigModule,
}
