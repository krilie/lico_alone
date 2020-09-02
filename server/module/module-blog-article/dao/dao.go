package dao

import (
	context_enum "github.com/krilie/lico_alone/common/com-model/context-enum"
	"github.com/krilie/lico_alone/component/ndb"
	"github.com/krilie/lico_alone/component/nlog"
)

type BlogArticleDao struct {
	*ndb.NDb
	log *nlog.NLog
}

func NewBlogArticleDao(ndb *ndb.NDb, log *nlog.NLog) *BlogArticleDao {
	log = log.WithField(context_enum.Module.Str(), "blog article dao")
	//err := ndb.GetDb(context.NewContext()).AutoMigrate(&model.Article{})
	//if err != nil {
	//	panic(err)
	//}
	return &BlogArticleDao{
		NDb: ndb,
		log: log,
	}
}
