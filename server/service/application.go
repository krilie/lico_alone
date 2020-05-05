package service

import (
	"github.com/krilie/lico_alone/common/config"
	"github.com/krilie/lico_alone/run_env"
	"github.com/krilie/lico_alone/service/cron-job-service"
	"github.com/krilie/lico_alone/service/init-data-service"
	"github.com/krilie/lico_alone/service/notification-email-service"
	union_service "github.com/krilie/lico_alone/service/union-service"
	"github.com/krilie/lico_alone/service/user-service"
)

type App struct {
	UserService              *user_service.UserService
	InitService              *init_data_service.InitDataService
	CronJobService           *cron_job_service.CronJobService
	NotificationEmailService *notification_email_service.NotificationEmailService
	UnionService             *union_service.UnionService
	Version                  string
	GitCommit                string
	BuildTime                string
	GoVersion                string
	Cfg                      config.Config
}

func NewApp(cfg *config.Config, runEnv *run_env.RunEnv,
	UserService *user_service.UserService,
	InitService *init_data_service.InitDataService,
	CronJobService *cron_job_service.CronJobService,
	NotificationEmailService *notification_email_service.NotificationEmailService,
	UnionService *union_service.UnionService,
) *App {
	return &App{
		UserService:              UserService,
		InitService:              InitService,
		CronJobService:           CronJobService,
		NotificationEmailService: NotificationEmailService,
		UnionService:             UnionService,
		Version:                  runEnv.Version,
		GitCommit:                runEnv.GitCommit,
		BuildTime:                runEnv.BuildTime,
		GoVersion:                runEnv.GoVersion,
		Cfg:                      *cfg,
	}
}
