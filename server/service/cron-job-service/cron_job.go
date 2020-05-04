// Package cron_job 定时任务
package cron_job_service

import (
	union_service "github.com/krilie/lico_alone/service/union-service"
)

type CronJobService struct {
	*union_service.UnionService
}

func NewCronJobService(unionService *union_service.UnionService) *CronJobService {
	return &CronJobService{UnionService: unionService}
}
