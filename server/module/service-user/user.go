package service_user

import (
	context_enum "github.com/krilie/lico_alone/common/com-model/context-enum"
	"github.com/krilie/lico_alone/component/nlog"
	service2 "github.com/krilie/lico_alone/module/module-blog-article/service"
	ConfigService "github.com/krilie/lico_alone/module/module-config/service"
	service3 "github.com/krilie/lico_alone/module/module-file/service"
	MessageService "github.com/krilie/lico_alone/module/module-message/service"
	"github.com/krilie/lico_alone/module/module-user/service"
)

type UserService struct {
	log           *nlog.NLog
	moduleUser    *service.UserModule
	moduleMsg     *MessageService.MessageModule
	moduleConfig  *ConfigService.ConfigModule
	moduleArticle *service2.BlogArticleModule
	moduleFile    *service3.FileModule
}

func (u *UserService) GetAuthFace() *service.UserModule {
	return u.moduleUser
}

func NewUserService(moduleArticle *service2.BlogArticleModule, moduleFile *service3.FileModule, log *nlog.NLog, moduleConfig *ConfigService.ConfigModule, moduleUser *service.UserModule, moduleMsg *MessageService.MessageModule) *UserService {
	log = log.WithField(context_enum.Module.Str(), "service user")
	return &UserService{
		log:           log,
		moduleUser:    moduleUser,
		moduleMsg:     moduleMsg,
		moduleConfig:  moduleConfig,
		moduleArticle: moduleArticle,
		moduleFile:    moduleFile,
	}
}
