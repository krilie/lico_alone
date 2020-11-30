package service

import (
	"github.com/krilie/lico_alone/module/module-customer/dao"
)

var DigModuleCustomerProviderAll = []interface{}{
	dao.NewCustomerDao,
	NewCustomerModule,
}
