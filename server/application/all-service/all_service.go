package all_service

import (
	"context"
	"github.com/krilie/lico_alone/common/config"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/common/utils/time_util"
	"github.com/krilie/lico_alone/component/nlog"
	AccountService "github.com/krilie/lico_alone/module/bookkeeping/service"
	"github.com/krilie/lico_alone/module/config/model"
	ConfigService "github.com/krilie/lico_alone/module/config/service"
	FileService "github.com/krilie/lico_alone/module/file/service"
	MessageService "github.com/krilie/lico_alone/module/message/service"
	UserService "github.com/krilie/lico_alone/module/module-user/service"
	"time"
)

type AllService struct {
	UserService    *UserService.UserService
	ConfigService  *ConfigService.Service
	FileService    *FileService.Service
	AccountService *AccountService.Service
	Message        *MessageService.Service
}

func NewAllService(cfg config.Config) *AllService {
	return &AllService{
		UserService:    UserService.NewUserService(cfg.DB),
		ConfigService:  ConfigService.NewService(cfg.DB),
		FileService:    FileService.NewService(cfg),
		AccountService: AccountService.NewService(cfg.DB),
		Message:        MessageService.NewService(cfg),
	}
}

// SendRunUpEmail 发送服务启动消息
func (a *AllService) SendServiceUpEmail(ctx context.Context) error {
	var log = nlog.NewLog(ctx, "application/all-service/all_service.go:34", "SendServiceUpEmail")
	emailAddr, err := a.ConfigService.GetValueStr(ctx, model.CommonNotificationEmail)
	if err != nil {
		log.Error(err)
		return err
	}
	if emailAddr == nil || *emailAddr == "" {
		return errs.NewBadRequest().WithMsg("no email config")
	}
	// 发送上线邮件
	err = a.Message.SendEmail(ctx, *emailAddr, "app-server", "启动成功: "+time.Now().Format(time_util.DefaultFormat))
	if err != nil {
		log.Error(err)
	}
	return err
}

// SendGoodMorningEmail 发送早上好消息
func (a *AllService) SendGoodMorningEmail(ctx context.Context) error {
	log := nlog.NewLog(ctx, "application/all-service/all_service.go:45", "SendGoodMorningEmail")
	emailAddr, err := a.ConfigService.GetValueStr(ctx, model.CommonNotificationEmail)
	if err != nil {
		log.Error(err)
		return err
	}
	if emailAddr == nil || *emailAddr == "" {
		return errs.NewBadRequest().WithMsg("no email config")
	}
	err = a.Message.SendEmail(ctx, *emailAddr, "早上好", "早上好: "+time.Now().Format(time_util.DefaultFormat))
	if err != nil {
		log.Error(err)
	}
	return err
}

// SendGoodMorningEmail 发送服务关闭消息
func (a *AllService) SendServiceEndEmail(ctx context.Context) error {
	var log = nlog.NewLog(ctx, "application/all-service/all_service.go:50", "SendServiceEndEmail")
	emailAddr, err := a.ConfigService.GetValueStr(ctx, model.CommonNotificationEmail)
	if err != nil {
		log.Error(err)
		return err
	}
	if emailAddr == nil || *emailAddr == "" {
		return errs.NewBadRequest().WithMsg("no email config")
	}
	err = a.Message.SendEmail(ctx, *emailAddr, "app-server", "服务关闭: "+time.Now().Format(time_util.DefaultFormat))
	if err != nil {
		log.Error(err)
	}
	return err
}
