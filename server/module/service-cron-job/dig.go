package service_cron_job

import "github.com/krilie/lico_alone/common/appdig"

// DigProvider provider
func DigProvider() {
	appdig.Container.MustProvide(NewCronJobService)
}
