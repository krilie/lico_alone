package service

import (
	"github.com/krilie/lico_alone/common/dig"
	"github.com/krilie/lico_alone/module/module-blog-article/dao"
)

// DigProvider provider
func DigProvider() {
	dig.Container.MustProvide(NewBlogArticleModule)
}

func DigProviderWithDao() {
	dao.DigProvider()
	DigProvider()
}
