package service

import (
	context_enum "github.com/krilie/lico_alone/common/com-model/context-enum"
	"github.com/krilie/lico_alone/component/broker"
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/krilie/lico_alone/module/module-test/dao"
)

// 系统配置服务
type TestArticleModule struct {
	Dao    *dao.TestArticleDao
	log    *nlog.NLog
	broker *broker.Broker
}

func NewTestArticleModule(log *nlog.NLog, dao *dao.TestArticleDao, broker *broker.Broker) *TestArticleModule {
	log = log.WithField(context_enum.Module.Str(), "test_article service")
	var module = &TestArticleModule{
		Dao:    dao,
		log:    log,
		broker: broker,
	}
	return module
}
