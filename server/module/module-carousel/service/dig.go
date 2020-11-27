package service

import (
	"github.com/krilie/lico_alone/common/appdig"
	"github.com/krilie/lico_alone/module/module-carousel/dao"
)

// DigProvider provider
func DigModuleCarouselProviderAll(c *appdig.AppContainer) {
	c.MustProvide(dao.NewCarouseDao)
	c.MustProvide(NewCarouselModule)
}
