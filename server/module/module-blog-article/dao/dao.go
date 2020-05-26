package dao

import (
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/component/ndb"
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/krilie/lico_alone/module/module-blog-article/model"
)

type BlogArticleDao struct {
	*ndb.NDb
	log *nlog.NLog
}

func NewBlogArticleDao(ndb *ndb.NDb, log *nlog.NLog) *BlogArticleDao {
	err := ndb.GetDb(context.NewContext()).AutoMigrate(&model.Article{}).Error
	if err != nil {
		panic(err)
	}
	return &BlogArticleDao{
		NDb: ndb,
		log: log,
	}
}
