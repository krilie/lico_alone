package service

import (
	"github.com/krilie/lico_alone/common/appdig"
	"github.com/krilie/lico_alone/component"
	"github.com/krilie/lico_alone/module/module-test/dao"
)

var DigModuleTestArticleProviderAll = []interface{}{
	dao.NewTestArticleDao,
	NewTestArticleModule,
}

// 测试用
func BuildTestContainer() *appdig.AppContainer {
	var container = appdig.NewAppDig()
	container.
		MustProvides(component.DigComponentProviderAll).
		MustProvides(DigModuleTestArticleProviderAll)
	return container
}
