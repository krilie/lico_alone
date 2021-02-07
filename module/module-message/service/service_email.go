package service

import (
	"context"
	"github.com/krilie/lico_alone/common/com-model"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/common/utils/id_util"
	"github.com/krilie/lico_alone/module/module-message/model"
	"gorm.io/gorm"
	"time"
)

func (s *MessageModule) SendEmail(ctx context.Context, to, subject, content string) error {
	email := &model.MessageEmail{
		Model: com_model.Model{
			Id:        id_util.GetUuid(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: gorm.DeletedAt{},
		},
		SendTime:  time.Now(),
		From:      "sys",
		To:        to,
		Subject:   subject,
		Content:   content,
		IsSuccess: true,
		Other:     "自由邮件",
	}
	err := s.email.SendEmail(ctx, to, subject, content)
	if err != nil {
		s.log.Get(ctx).Error(err)
		email.IsSuccess = false
		email.Other = err.Error()
		err = s.Dao.CreateMessageEmail(ctx, email)
		if err != nil {
			s.log.Get(ctx).Error(err)
			return err
		}
		return errs.NewInternal().WithError(err)
	}
	err = s.Dao.CreateMessageEmail(ctx, email)
	if err != nil {
		s.log.Get(ctx).Error(err)
		return err
	}
	return nil
}
