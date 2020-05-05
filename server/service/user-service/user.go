package user_service

import (
	"github.com/krilie/lico_alone/component/nlog"
	MessageService "github.com/krilie/lico_alone/module/module-message/service"
	"github.com/krilie/lico_alone/module/module-user/service"
)

type UserService struct {
	log        *nlog.NLog
	moduleUser *service.UserService
	moduleMsg  *MessageService.MessageService
}

func NewUserService(log *nlog.NLog, moduleUser *service.UserService, moduleMsg *MessageService.MessageService) *UserService {
	return &UserService{
		log:        log,
		moduleUser: moduleUser,
		moduleMsg:  moduleMsg,
	}
}
