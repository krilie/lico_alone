package application

import (
	all_service "github.com/krilie/lico_alone/application/all-service"
	"github.com/krilie/lico_alone/application/cron-job"
	"github.com/krilie/lico_alone/application/init-data"
	"github.com/krilie/lico_alone/application/user-api"
	"github.com/krilie/lico_alone/common/config"
)

type App struct {
	User    *user_api.AppUser
	Init    *init_data.Init
	CronJob *cron_job.CronJob
	All     *all_service.AllService
}

func NewApp(cfg config.Config) *App {
	allSrv := all_service.NewAllService(cfg)
	return &App{
		User:    user_api.NewAppUser(allSrv),
		Init:    init_data.NewInit(allSrv),
		CronJob: cron_job.NewCronJob(allSrv),
		All:     allSrv,
	}
}
