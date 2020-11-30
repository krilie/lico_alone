package service

import (
	"github.com/krilie/lico_alone/module/module-carousel/dao"
)

var DigModuleCarouselProviderAll = []interface{}{
	dao.NewCarouseDao,
	NewCarouselModule,
}
