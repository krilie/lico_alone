package user_service

import (
	"github.com/krilie/lico_alone/common/dig"
	"github.com/krilie/lico_alone/component/nlog"
	ConfigService "github.com/krilie/lico_alone/module/module-config/service"
	MessageService "github.com/krilie/lico_alone/module/module-message/service"
	"github.com/krilie/lico_alone/module/module-user/service"
)

func init() {
	dig.Container.MustProvide(func(log *nlog.NLog, moduleConfig *ConfigService.ConfigService, moduleUser *service.UserService, moduleMsg *MessageService.MessageService) *UserService {
		return NewUserService(log, moduleConfig, moduleUser, moduleMsg)
	})
}
