package service

import (
	"github.com/krilie/lico_alone/module/module-zunion/dao"
)

var DigModuleZUnionProviderAll = []interface{}{
	NewZUnionModule,
	dao.NewZUnionDao,
}
