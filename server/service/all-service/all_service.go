package all_service

import (
	"context"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/common/utils/time_util"
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/krilie/lico_alone/module/module-config/model"
	ConfigService "github.com/krilie/lico_alone/module/module-config/service"
	FileService "github.com/krilie/lico_alone/module/module-file/service"
	MessageService "github.com/krilie/lico_alone/module/module-message/service"
	UserService "github.com/krilie/lico_alone/module/module-user/service"
	"github.com/prometheus/common/log"
	"time"
)

type AllService struct {
	UserService    *UserService.UserService
	ConfigService  *ConfigService.ConfigService
	FileService    *FileService.FileService
	MessageService *MessageService.MessageService
	log            *nlog.NLog
}

func NewAllService(UserService *UserService.UserService,
	ConfigService *ConfigService.ConfigService,
	FileService *FileService.FileService,
	MessageService *MessageService.MessageService,
	log *nlog.NLog) *AllService {
	return &AllService{
		UserService:    UserService,
		ConfigService:  ConfigService,
		FileService:    FileService,
		MessageService: MessageService,
		log:            log,
	}
}

// SendRunUpEmail 发送服务启动消息
func (a *AllService) SendServiceUpEmail(ctx context.Context) error {
	emailAddr, err := a.ConfigService.GetValueStr(ctx, model.CommonNotificationEmail)
	if err != nil {
		log.Error(err)
		return err
	}
	if emailAddr == nil || *emailAddr == "" {
		return errs.NewInternal().WithMsg("no email config")
	}
	// 发送上线邮件
	err = a.MessageService.SendEmail(ctx, *emailAddr, "app-server", "启动成功: "+time.Now().Format(time_util.DefaultFormat))
	if err != nil {
		log.Error(err)
	}
	return err
}

// SendGoodMorningEmail 发送早上好消息
func (a *AllService) SendGoodMorningEmail(ctx context.Context) error {
	emailAddr, err := a.ConfigService.GetValueStr(ctx, model.CommonNotificationEmail)
	if err != nil {
		log.Error(err)
		return err
	}
	if emailAddr == nil || *emailAddr == "" {
		return errs.NewInternal().WithMsg("no email config")
	}
	err = a.MessageService.SendEmail(ctx, *emailAddr, "早上好", "早上好: "+time.Now().Format(time_util.DefaultFormat))
	if err != nil {
		log.Error(err)
	}
	return err
}

// SendGoodMorningEmail 发送服务关闭消息
func (a *AllService) SendServiceEndEmail(ctx context.Context) error {
	emailAddr, err := a.ConfigService.GetValueStr(ctx, model.CommonNotificationEmail)
	if err != nil {
		log.Error(err)
		return err
	}
	if emailAddr == nil || *emailAddr == "" {
		return errs.NewInternal().WithMsg("no email config")
	}
	err = a.MessageService.SendEmail(ctx, *emailAddr, "app-server", "服务关闭: "+time.Now().Format(time_util.DefaultFormat))
	if err != nil {
		log.Error(err)
	}
	return err
}
