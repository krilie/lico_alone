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

// SendRegisterSms 发送注册短信
func (s *Service) SendRegisterSms(ctx context.Context, phone, code string) error {
	log := clog.NewLog(ctx, "module/message/service/service_sms.go:9", "SetTx")
	err := s.sms.SendRegisterSms(ctx, phone, code)
	if err != nil {
		log.Error(err)
		return errs.NewInternal().WithMsg("短信发送失败").WithError(err)
	}
	err = s.Dao.Db.Create(model.MessageSms{
		Model: cmodel.Model{
			Id:         id_util.GetUuid(),
			CreateTime: time.Now(),
		},
		SendTime:  time.Now(),
		Name:      "",
		To:        phone,
		Message:   code,
		IsSuccess: true,
		Other:     "注册短信",
	}).Error
	if err != nil {
		log.Error(err)
		return errs.NewErrDbCreate().WithError(err)
	}
	return nil
}
