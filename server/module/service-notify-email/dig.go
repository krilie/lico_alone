package service_notify_email

import (
	"github.com/krilie/lico_alone/common/dig"
)

func init() {
	dig.Container.MustProvide(NewNotificationEmailService)
}
