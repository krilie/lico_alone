package service_init_data

import (
	"context"
	"github.com/krilie/lico_alone/component/ndb"
	"github.com/krilie/lico_alone/component/nlog"
	ConfigService "github.com/krilie/lico_alone/module/module-config/service"
	"github.com/krilie/lico_alone/module/module-user/service"
)

type InitDataService struct {
	log          *nlog.NLog
	moduleConfig *ConfigService.ConfigService
	moduleUser   *service.UserService
}

func (initData *InitDataService) GetNDb(ctx context.Context) *ndb.NDb {
	return initData.moduleConfig.Dao.NDb
}

func NewInitDataService(log *nlog.NLog, moduleConfig *ConfigService.ConfigService, moduleUser *service.UserService) *InitDataService {
	return &InitDataService{
		log:          log,
		moduleConfig: moduleConfig,
		moduleUser:   moduleUser,
	}
}
