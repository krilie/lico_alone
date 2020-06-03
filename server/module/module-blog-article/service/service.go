package service

import (
	context_enum "github.com/krilie/lico_alone/common/com-model/context-enum"
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/krilie/lico_alone/module/module-blog-article/dao"
)

// 系统配置服务
type BlogArticleModule struct {
	Dao *dao.BlogArticleDao
	log *nlog.NLog
}

func NewBlogArticleModule(log *nlog.NLog, dao *dao.BlogArticleDao) *BlogArticleModule {
	log = log.WithField(context_enum.Module.Str(), "blog article service")
	return &BlogArticleModule{
		Dao: dao,
		log: log,
	}
}
