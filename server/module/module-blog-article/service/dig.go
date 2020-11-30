package service

import (
	"github.com/krilie/lico_alone/common/appdig"
	"github.com/krilie/lico_alone/component"
	"github.com/krilie/lico_alone/module/module-blog-article/dao"
)

var DigModuleBlogArticleProviderAll = []interface{}{
	dao.NewBlogArticleDao,
	NewBlogArticleModule,
}

// 测试用
func BuildTestContainer() *appdig.AppContainer {
	var container = appdig.NewAppDig()
	container.
		MustProvides(component.DigComponentProviderAllForTest).
		MustProvides(DigModuleBlogArticleProviderAll)
	return container
}
