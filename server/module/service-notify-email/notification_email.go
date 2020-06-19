package service_notify_email

import (
	"context"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/common/utils/time_util"
	"github.com/krilie/lico_alone/module/module-config/model"
	ConfigService "github.com/krilie/lico_alone/module/module-config/service"
	MessageService "github.com/krilie/lico_alone/module/module-message/service"
	"github.com/prometheus/common/log"
	"time"
)

type NotificationEmailService struct {
	ModuleConfig  *ConfigService.ConfigModule
	ModuleMessage *MessageService.MessageModule
}

func NewNotificationEmailService(moduleConfig *ConfigService.ConfigModule, moduleMessage *MessageService.MessageModule) *NotificationEmailService {
	return &NotificationEmailService{
		ModuleConfig:  moduleConfig,
		ModuleMessage: moduleMessage,
	}
}

// SendRunUpEmail 发送服务启动消息
func (a *NotificationEmailService) SendServiceUpEmail(ctx context.Context) error {
	emailAddr, err := a.ModuleConfig.GetValueStr(ctx, model.ConfigItemsNotificationEmail.Val())
	if err != nil {
		log.Error(err)
		return err
	}
	if emailAddr == nil || *emailAddr == "" {
		return errs.NewInternal().WithMsg("no email config")
	}
	// 发送上线邮件
	err = a.ModuleMessage.SendEmail(ctx, *emailAddr, "app-server", "启动成功: "+time.Now().Format(time_util.DefaultFormat))
	if err != nil {
		log.Error(err)
	}
	return err
}

// SendGoodMorningEmail 发送早上好消息
func (a *NotificationEmailService) SendGoodMorningEmail(ctx context.Context) error {
	emailAddr, err := a.ModuleConfig.GetValueStr(ctx, model.ConfigItemsNotificationEmail.Val())
	if err != nil {
		log.Error(err)
		return err
	}
	if emailAddr == nil || *emailAddr == "" {
		return errs.NewInternal().WithMsg("no email config")
	}
	err = a.ModuleMessage.SendEmail(ctx, *emailAddr, "早上好", "早上好: "+time.Now().Format(time_util.DefaultFormat))
	if err != nil {
		log.Error(err)
	}
	return err
}

// SendGoodMorningEmail 发送服务关闭消息
func (a *NotificationEmailService) SendServiceEndEmail(ctx context.Context) error {
	emailAddr, err := a.ModuleConfig.GetValueStr(ctx, model.ConfigItemsNotificationEmail.Val())
	if err != nil {
		log.Error(err)
		return err
	}
	if emailAddr == nil || *emailAddr == "" {
		return errs.NewInternal().WithMsg("no email config")
	}
	err = a.ModuleMessage.SendEmail(ctx, *emailAddr, "app-server", "服务关闭: "+time.Now().Format(time_util.DefaultFormat))
	if err != nil {
		log.Error(err)
	}
	return err
}
