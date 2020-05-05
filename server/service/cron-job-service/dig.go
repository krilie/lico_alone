package cron_job_service

import (
	"github.com/krilie/lico_alone/common/dig"
	union_service "github.com/krilie/lico_alone/service/union-service"
)

func init() {
	dig.Container.MustProvide(func(unionService *union_service.UnionService) *CronJobService {
		return NewCronJobService(unionService)
	})
}
