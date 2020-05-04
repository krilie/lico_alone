// Package cron_job 定时任务
package cron_job_service

import (
	"github.com/krilie/lico_alone/module/module-user/service"
	all_service "github.com/krilie/lico_alone/service/notification-email-service"
)

type CronJob struct {
	UserService *service.UserService
}

func NewCronJob(allSrv *all_service.NotificationEmailService) *CronJob {
	return &CronJob{UserService: allSrv.UserService}
}
