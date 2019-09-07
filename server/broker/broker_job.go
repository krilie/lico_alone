package broker

import (
	all_service "github.com/krilie/lico_alone/application/all-service"
	"github.com/krilie/lico_alone/common/broker"
	"github.com/krilie/lico_alone/common/ccontext"
	service4 "github.com/krilie/lico_alone/module/account/service"
	service2 "github.com/krilie/lico_alone/module/config/service"
	service3 "github.com/krilie/lico_alone/module/file/service"
	"github.com/krilie/lico_alone/module/user/service"
)

// 消息处理工作
type BrokerJob struct {
	UserService    *service.Service
	ConfigService  *service2.Service
	FileService    *service3.Service
	AccountService *service4.Service
}

func NewBrokerJob(all all_service.AllService) *BrokerJob {
	return &BrokerJob{
		UserService:    all.UserService,
		ConfigService:  all.ConfigService,
		FileService:    all.FileService,
		AccountService: all.AccountService,
	}
}

func (a *BrokerJob) InitBrokerJob(ctx ccontext.Context) (close func()) {
	close2 := broker.InitBrokers(ctx)
	a.UserService.RegisterBroker(ctx)
	a.FileService.RegisterBroker(ctx)
	a.AccountService.RegisterBroker(ctx)
	return close2
}
