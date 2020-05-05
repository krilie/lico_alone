package service

import (
	"github.com/krilie/lico_alone/common/com-model/run-env"
	"github.com/krilie/lico_alone/common/config"
	"github.com/krilie/lico_alone/common/dig"
	cron_job_service "github.com/krilie/lico_alone/service/cron-job-service"
	init_data_service "github.com/krilie/lico_alone/service/init-data-service"
	notification_email_service "github.com/krilie/lico_alone/service/notification-email-service"
	union_service "github.com/krilie/lico_alone/service/union-service"
	user_service "github.com/krilie/lico_alone/service/user-service"
)

func init() {
	dig.Container.MustProvide(func(cfg *config.Config, runEnv *run_env.RunEnv,
		UserService *user_service.UserService,
		InitService *init_data_service.InitDataService,
		CronJobService *cron_job_service.CronJobService,
		NotificationEmailService *notification_email_service.NotificationEmailService,
		UnionService *union_service.UnionService) *App {
		return NewApp(cfg, runEnv, UserService, InitService, CronJobService, NotificationEmailService, UnionService)
	})
}
