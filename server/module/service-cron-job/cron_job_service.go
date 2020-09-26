package service_cron_job

import (
	"context"
	context_enum "github.com/krilie/lico_alone/common/com-model/context-enum"
	"github.com/krilie/lico_alone/component/cron"
	"github.com/krilie/lico_alone/component/ndb"
	"github.com/krilie/lico_alone/component/nlog"
	ConfigService "github.com/krilie/lico_alone/module/module-config/service"
	MessageService "github.com/krilie/lico_alone/module/module-message/service"
	"github.com/krilie/lico_alone/module/module-user/service"
)

type CronJobService struct {
	log           *nlog.NLog  // log是全局资源 不由此service释放或关闭
	cron          *cron.NCron // 同上
	moduleConfig  *ConfigService.ConfigModule
	moduleUser    *service.UserModule
	moduleMessage *MessageService.MessageModule
}

func (cron *CronJobService) GetNDb(ctx context.Context) *ndb.NDb {
	return cron.moduleConfig.Dao.NDb
}

func NewCronJobService(log *nlog.NLog, cron *cron.NCron, moduleMessage *MessageService.MessageModule, moduleConfig *ConfigService.ConfigModule, moduleUser *service.UserModule) *CronJobService {
	log = log.WithField(context_enum.Module.Str(), "CronJobService")
	return &CronJobService{
		log:           log,
		cron:          cron,
		moduleConfig:  moduleConfig,
		moduleUser:    moduleUser,
		moduleMessage: moduleMessage,
	}
}
