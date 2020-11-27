package service

import (
	"github.com/krilie/lico_alone/module/module-dynamic-share/dao"
)

var DigModuleDynamicShareProviderAll = []interface{}{
	dao.NewDynamicShareDao,
	NewDynamicShareModule,
}
