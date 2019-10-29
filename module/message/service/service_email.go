package service

import (
	"context"
	"github.com/krilie/lico_alone/common/clog"
	"github.com/krilie/lico_alone/common/cmodel"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/common/utils/id_util"
	"github.com/krilie/lico_alone/module/message/model"
	"time"
)

func (s *Service) SendEmail(ctx context.Context, to, subject, content string) error {
	log := clog.NewLog(ctx, "module/message/service/service_email.go:8", "SetTx")
	err := s.email.SendEmail(ctx, to, subject, content)
	if err != nil {
		log.Error(err)
		return errs.NewInternal().WithError(err)
	}
	err = s.Dao.Db.Create(model.MessageSms{
		Model: cmodel.Model{
			Id:         id_util.GetUuid(),
			CreateTime: time.Now(),
		},
		SendTime:  time.Now(),
		Name:      "",
		To:        to,
		Message:   subject + ":" + content,
		IsSuccess: true,
		Other:     "自由邮件",
	}).Error
	if err != nil {
		log.Error(err)
		return errs.ErrDbCreate.WithError(err)
	}
	return nil
}
