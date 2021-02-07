package service

import (
	"github.com/krilie/lico_alone/module/module-user/dao"
)

var DigModuleUserProviderAll = []interface{}{
	NewUserModule,
	dao.NewUserDao,
}
