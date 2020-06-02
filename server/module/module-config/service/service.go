package service

import (
	context_enum "github.com/krilie/lico_alone/common/com-model/context-enum"
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/krilie/lico_alone/module/module-config/dao"
)

// 系统配置服务
type ConfigService struct {
	Dao *dao.ConfigDao
	log *nlog.NLog
}

func NewService(log *nlog.NLog, dao *dao.ConfigDao) *ConfigService {
	log = log.WithField(context_enum.Module.Str(), "config service")
	return &ConfigService{
		Dao: dao,
		log: log,
	}
}
