package service_notify_email

import (
	"github.com/krilie/lico_alone/common/appdig"
)

// DigProvider provider
func DigProvider() {
	appdig.Container.MustProvide(NewNotificationEmailService)
}
