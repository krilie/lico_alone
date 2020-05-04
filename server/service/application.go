package service

import (
	"context"
	"github.com/krilie/lico_alone/common/config"
	"github.com/krilie/lico_alone/service/cron-job-service"
	"github.com/krilie/lico_alone/service/init-data-service"
	"github.com/krilie/lico_alone/service/notification-email-service"
	"github.com/krilie/lico_alone/service/user-service"
)

type App struct {
	UserService              *user_service.AppUser
	InitService              *init_data_service.init
	CronJobService           *cron_job_service.CronJobService
	NotificationEmailService *notification_email_service.NotificationEmailService
	Version                  string
	GitCommit                string
	BuildTime                string
	GoVersion                string
	Cfg                      config.Config
}

func NewApp(ctx context.Context, cfg config.Config, version string, buildTime string, gitCommit, goVersion string) *App {
	allSrv := all_service.NewAllService(cfg)
	return &App{
		User:      user_service.NewAppUser(allSrv),
		Init:      init_data_service.NewInit(allSrv),
		CronJob:   cron_job_service.NewCronJob(allSrv),
		All:       allSrv,
		Version:   version,
		GitCommit: gitCommit,
		BuildTime: buildTime,
		GoVersion: goVersion,
		Cfg:       cfg,
	}
}
