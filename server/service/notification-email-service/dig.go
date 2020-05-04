package notification_email_service

import (
	"github.com/krilie/lico_alone/common/dig"
	union_service "github.com/krilie/lico_alone/service/union-service"
)

func init() {
	dig.Container.MustProvide(func(unionService *union_service.UnionService) *NotificationEmailService {
		return NewNotificationEmailService(unionService)
	})
}
