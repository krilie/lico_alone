package service

import (
	context_enum "github.com/krilie/lico_alone/common/com-model/context-enum"
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/krilie/lico_alone/module/module-config/dao"
)

// ConfigModule 系统配置服务
type ConfigModule struct {
	Dao *dao.ConfigDao
	log *nlog.NLog
}

func NewConfigModule(log *nlog.NLog, dao *dao.ConfigDao) *ConfigModule {
	log = log.WithField(context_enum.Module.Str(), "config service")
	return &ConfigModule{
		Dao: dao,
		log: log,
	}
}
