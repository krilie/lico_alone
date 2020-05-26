package dao

import (
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/dig"
	"github.com/krilie/lico_alone/component/ndb"
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/krilie/lico_alone/module/module-blog-article/model"
)

func init() {
	dig.Container.MustProvide(func(ndb *ndb.NDb, log *nlog.NLog) *BlogArticleDao {
		ndb.GetDb(context.NewContext()).AutoMigrate(new(model.Article))
		return NewBlogArticleDao(ndb, log)
	})
}
