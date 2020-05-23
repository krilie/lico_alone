package service

import (
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/krilie/lico_alone/module/module-blog-article/dao"
)

// 系统配置服务
type BlogArticleService struct {
	Dao *dao.BlogArticleDao
	log *nlog.NLog
}

func NewService(log *nlog.NLog, dao *dao.BlogArticleDao) *BlogArticleService {
	return &BlogArticleService{
		Dao: dao,
		log: log,
	}
}
