package user_service

import (
	"github.com/krilie/lico_alone/component/nlog"
	service2 "github.com/krilie/lico_alone/module/module-blog-article/service"
	ConfigService "github.com/krilie/lico_alone/module/module-config/service"
	MessageService "github.com/krilie/lico_alone/module/module-message/service"
	"github.com/krilie/lico_alone/module/module-user/service"
)

type UserService struct {
	log           *nlog.NLog
	moduleUser    *service.UserService
	moduleMsg     *MessageService.MessageService
	moduleConfig  *ConfigService.ConfigService
	moduleArticle *service2.BlogArticleService
}

func (u *UserService) GetAuthFace() *service.UserService {
	return u.moduleUser
}

func NewUserService(log *nlog.NLog, moduleConfig *ConfigService.ConfigService, moduleUser *service.UserService, moduleMsg *MessageService.MessageService) *UserService {
	return &UserService{
		log:          log,
		moduleUser:   moduleUser,
		moduleMsg:    moduleMsg,
		moduleConfig: moduleConfig,
	}
}
