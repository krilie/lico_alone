package service

import (
	"github.com/krilie/lico_alone/common/appdig"
	"github.com/krilie/lico_alone/component"
	"github.com/krilie/lico_alone/module/module-blog-article/dao"
	module_like_dislike "github.com/krilie/lico_alone/module/module-like-dislike"
)

var DigModuleBlogArticleProviderAll = []interface{}{
	dao.NewBlogArticleDao,
	NewBlogArticleModule,
}

// BuildTestContainer 测试用
func BuildTestContainer() *appdig.AppContainer {
	var container = appdig.NewAppDig()
	container.
		MustProvides(component.DigComponentProviderAll).
		MustProvides(DigModuleBlogArticleProviderAll).
		MustProvides(module_like_dislike.DigModuleLikeDisLikeProviderAll)
	return container
}
