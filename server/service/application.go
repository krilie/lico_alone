package service

import (
	"github.com/krilie/lico_alone/common/config"
	"github.com/krilie/lico_alone/run_env"
	common_service "github.com/krilie/lico_alone/service/common-service"
	"github.com/krilie/lico_alone/service/cron-job-service"
	"github.com/krilie/lico_alone/service/init-data-service"
	"github.com/krilie/lico_alone/service/notification-email-service"
	"github.com/krilie/lico_alone/service/user-service"
)

type App struct {
	UserService              *user_service.UserService
	InitService              *init_data_service.InitDataService
	CronJobService           *cron_job_service.CronJobService
	NotificationEmailService *notification_email_service.NotificationEmailService
	CommonService            *common_service.CommonService
	RunEnv                   *run_env.RunEnv
	Cfg                      *config.Config
}

func NewApp(cfg *config.Config, runEnv *run_env.RunEnv,
	UserService *user_service.UserService,
	InitService *init_data_service.InitDataService,
	CronJobService *cron_job_service.CronJobService,
	NotificationEmailService *notification_email_service.NotificationEmailService,
	CommonService *common_service.CommonService,
) *App {
	return &App{
		UserService:              UserService,
		InitService:              InitService,
		CronJobService:           CronJobService,
		NotificationEmailService: NotificationEmailService,
		CommonService:            CommonService,
		RunEnv:                   runEnv,
		Cfg:                      cfg,
	}
}
