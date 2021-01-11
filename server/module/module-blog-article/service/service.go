package service

import (
	context2 "context"
	context_enum "github.com/krilie/lico_alone/common/com-model/context-enum"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/component/broker"
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/krilie/lico_alone/module/module-blog-article/dao"
	module_like_dislike "github.com/krilie/lico_alone/module/module-like-dislike"
)

// 系统配置服务
type BlogArticleModule struct {
	Dao               *dao.BlogArticleDao
	log               *nlog.NLog
	broker            *broker.Broker
	likeDislikeModule *module_like_dislike.LikeDisLikeModule
}

func NewBlogArticleModule(log *nlog.NLog, dao *dao.BlogArticleDao, broker *broker.Broker, likeDisLike *module_like_dislike.LikeDisLikeModule) *BlogArticleModule {
	log = log.WithField(context_enum.Module.Str(), "blog article service")
	var module = &BlogArticleModule{
		Dao:               dao,
		log:               log,
		broker:            broker,
		likeDislikeModule: likeDisLike,
	}
	ctx := context.NewAppCtx(context2.Background())
	module.broker.MustRegister(ctx, module.HandleBrokerBlogArticleVisitedMessage)
	return module
}
