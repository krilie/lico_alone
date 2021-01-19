package service

import (
	context_enum "github.com/krilie/lico_alone/common/com-model/context-enum"
	"github.com/krilie/lico_alone/component/broker"
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/krilie/lico_alone/module/module-catchword/dao"
)

// 系统配置服务
type CatchwordModule struct {
	Dao    *dao.CatchwordDao
	log    *nlog.NLog
	broker *broker.Broker
}

func NewCatchwordModule(log *nlog.NLog, dao *dao.CatchwordDao, broker *broker.Broker) *CatchwordModule {
	log = log.WithField(context_enum.Module.Str(), "catchword service")
	var module = &CatchwordModule{
		Dao:    dao,
		log:    log,
		broker: broker,
	}
	return module
}
