package service_user

import (
	context_enum "github.com/krilie/lico_alone/common/com-model/context-enum"
	"github.com/krilie/lico_alone/component/nlog"
	service2 "github.com/krilie/lico_alone/module/module-blog-article/service"
	service4 "github.com/krilie/lico_alone/module/module-carousel/service"
	ConfigService "github.com/krilie/lico_alone/module/module-config/service"
	service3 "github.com/krilie/lico_alone/module/module-file/service"
	MessageService "github.com/krilie/lico_alone/module/module-message/service"
	"github.com/krilie/lico_alone/module/module-user/service"
)

type UserService struct {
	log            *nlog.NLog
	moduleUser     *service.UserModule
	moduleMsg      *MessageService.MessageModule
	moduleConfig   *ConfigService.ConfigModule
	ModuleArticle  *service2.BlogArticleModule
	ModuleFile     *service3.FileModule
	ModuleCarousel *service4.CarouselModule
}

func (a *UserService) GetAuthFace() *service.UserModule {
	return a.moduleUser
}

func NewUserService(
	moduleArticle *service2.BlogArticleModule,
	moduleFile *service3.FileModule,
	log *nlog.NLog,
	moduleConfig *ConfigService.ConfigModule,
	moduleUser *service.UserModule,
	moduleMsg *MessageService.MessageModule,
	moduleCarousel *service4.CarouselModule,
) *UserService {
	log = log.WithField(context_enum.Module.Str(), "service user")
	return &UserService{
		log:            log,
		moduleUser:     moduleUser,
		moduleMsg:      moduleMsg,
		moduleConfig:   moduleConfig,
		ModuleArticle:  moduleArticle,
		ModuleFile:     moduleFile,
		ModuleCarousel: moduleCarousel,
	}
}
