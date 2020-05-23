package service

import (
	"github.com/krilie/lico_alone/common/dig"
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/krilie/lico_alone/module/module-blog-article/dao"
)

func init() {
	dig.Container.MustProvide(func(log *nlog.NLog, dao *dao.BlogArticleDao) *BlogArticleService {
		return NewService(log, dao)
	})
}
