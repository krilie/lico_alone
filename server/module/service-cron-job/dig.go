package service_cron_job

import "github.com/krilie/lico_alone/common/dig"

// DigProvider provider
func DigProvider() {
	dig.Container.MustProvide(NewCronJobService)
}
