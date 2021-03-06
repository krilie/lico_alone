package service_init_data

import (
	"context"
	context_enum "github.com/krilie/lico_alone/common/com-model/context-enum"
	"github.com/krilie/lico_alone/component/ndb"
	"github.com/krilie/lico_alone/component/nlog"
	ConfigService "github.com/krilie/lico_alone/module/module-config/service"
	"github.com/krilie/lico_alone/module/module-user/service"
)

type InitDataService struct {
	log          *nlog.NLog
	moduleConfig *ConfigService.ConfigModule
	moduleUser   *service.UserModule
}

func (initData *InitDataService) GetNDb(ctx context.Context) *ndb.NDb {
	return initData.moduleConfig.Dao.NDb
}

func NewInitDataService(log *nlog.NLog, moduleConfig *ConfigService.ConfigModule, moduleUser *service.UserModule) *InitDataService {
	log = log.WithField(context_enum.Module.Str(), "service init")
	return &InitDataService{
		log:          log,
		moduleConfig: moduleConfig,
		moduleUser:   moduleUser,
	}
}
