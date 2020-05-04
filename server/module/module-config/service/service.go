package service

import (
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/krilie/lico_alone/module/config/dao"
)

// 系统配置服务
type ConfigService struct {
	Dao *dao.ConfigDao
	log *nlog.NLog
}

func NewService(log *nlog.NLog, dao *dao.ConfigDao) *ConfigService {
	return &ConfigService{
		Dao: dao,
		log: log,
	}
}
