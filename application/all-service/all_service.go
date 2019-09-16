package all_service

import (
	"github.com/krilie/lico_alone/common/config"
	service4 "github.com/krilie/lico_alone/module/account/service"
	service2 "github.com/krilie/lico_alone/module/config/service"
	service3 "github.com/krilie/lico_alone/module/file/service"
	"github.com/krilie/lico_alone/module/user/service"
)

type AllService struct {
	UserService    *service.Service
	ConfigService  *service2.Service
	FileService    *service3.Service
	AccountService *service4.Service
}

func NewAllService(cfg config.Config) *AllService {
	return &AllService{
		UserService:    service.NewService(cfg),
		ConfigService:  service2.NewService(cfg),
		FileService:    service3.NewService(cfg),
		AccountService: service4.NewService(cfg),
	}
}
