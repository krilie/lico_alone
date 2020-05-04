package service

import (
	"context"
	"github.com/krilie/lico_alone/common/config"
	all_service "github.com/krilie/lico_alone/service/all-service"
	"github.com/krilie/lico_alone/service/cron-job"
	"github.com/krilie/lico_alone/service/init-data"
	"github.com/krilie/lico_alone/service/user-api"
)

type App struct {
	User      *user_api.AppUser
	Init      *init_data.Init
	CronJob   *cron_job.CronJob
	All       *all_service.AllService
	Version   string
	GitCommit string
	BuildTime string
	GoVersion string
	Cfg       config.Config
}

func NewApp(ctx context.Context, cfg config.Config, version string, buildTime string, gitCommit, goVersion string) *App {
	allSrv := all_service.NewAllService(cfg)
	return &App{
		User:      user_api.NewAppUser(allSrv),
		Init:      init_data.NewInit(allSrv),
		CronJob:   cron_job.NewCronJob(allSrv),
		All:       allSrv,
		Version:   version,
		GitCommit: gitCommit,
		BuildTime: buildTime,
		GoVersion: goVersion,
		Cfg:       cfg,
	}
}
