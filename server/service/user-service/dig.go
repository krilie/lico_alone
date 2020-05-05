package user_service

import (
	"github.com/krilie/lico_alone/common/dig"
	"github.com/krilie/lico_alone/component/nlog"
	MessageService "github.com/krilie/lico_alone/module/module-message/service"
	"github.com/krilie/lico_alone/module/module-user/service"
)

func init() {
	dig.Container.MustProvide(func(log *nlog.NLog, moduleUser *service.UserService, moduleMsg *MessageService.MessageService) *UserService {
		return NewUserService(log, moduleUser, moduleMsg)
	})
}
