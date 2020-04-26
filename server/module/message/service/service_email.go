package service

import (
	"context"
	"github.com/krilie/lico_alone/common/com-model"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/common/utils/id_util"
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/krilie/lico_alone/module/message/model"
	"time"
)

func (s *Service) SendEmail(ctx context.Context, to, subject, content string) error {
	log := nlog.NewLog(ctx, "module/message/service/service_email.go:8", "NewWithTxOrFromCtx")
	email := &com_model.MessageEmail{
		Model:     com_model.Model{Id: id_util.GetUuid(), CreateTime: time.Now()},
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
		log.Error(err)
		email.IsSuccess = false
		email.Other = err.Error()
		err = s.Dao.CreateMessageEmail(ctx, email)
		if err != nil {
			log.Error(err)
			return err
		}
		return errs.NewInternal().WithError(err)
	}
	err = s.Dao.CreateMessageEmail(ctx, email)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
