// Package cron_job 定时任务
package cron_job

import (
	all_service "github.com/krilie/lico_alone/application/all-service"
	"github.com/krilie/lico_alone/module/module-user/service"
)

type CronJob struct {
	UserService *service.UserService
}

func NewCronJob(allSrv *all_service.AllService) *CronJob {
	return &CronJob{UserService: allSrv.UserService}
}
