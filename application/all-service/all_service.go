package all_service

import (
	"github.com/krilie/lico_alone/common/config"
	AccountService "github.com/krilie/lico_alone/module/account/service"
	ConfigService "github.com/krilie/lico_alone/module/config/service"
	FileService "github.com/krilie/lico_alone/module/file/service"
	MessageService "github.com/krilie/lico_alone/module/message/service"
	UserService "github.com/krilie/lico_alone/module/user/service"
)

type AllService struct {
	UserService    *UserService.Service
	ConfigService  *ConfigService.Service
	FileService    *FileService.Service
	AccountService *AccountService.Service
	Message        *MessageService.Service
}

func NewAllService(cfg config.Config) *AllService {
	return &AllService{
		UserService:    UserService.NewService(cfg.DB),
		ConfigService:  ConfigService.NewService(cfg.DB),
		FileService:    FileService.NewService(cfg),
		AccountService: AccountService.NewService(cfg.DB),
		Message:        MessageService.NewService(cfg),
	}
}
