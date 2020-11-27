package service

import (
	service_dynamic_share "github.com/krilie/lico_alone/module/module-dynamic-share/service"
	service_common "github.com/krilie/lico_alone/module/service-common"
	service_cronjob "github.com/krilie/lico_alone/module/service-cron-job"
	service_init_data "github.com/krilie/lico_alone/module/service-init-data"
	service_notify_email "github.com/krilie/lico_alone/module/service-notify-email"
	service_user "github.com/krilie/lico_alone/module/service-user"
)

func DigProviderService() {
	service_common.DigProvider()
	service_init_data.DigProvider()
	service_notify_email.DigProvider()
	service_user.DigProvider()
	service_cronjob.DigProvider()
	service_dynamic_share.DigProvider()
}
