package service

import (
	"context"
	"errors"
	"github.com/krilie/lico_alone/common/clog"
	"github.com/krilie/lico_alone/common/cmodel"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/common/utils/id_util"
	"github.com/krilie/lico_alone/module/message/model"
	"time"
)

// SendRegisterSms 发送注册短信
func (s *Service) SendRegisterSms(ctx context.Context, phone, code string) error {
	log := clog.NewLog(ctx, "module/message/service/service_sms.go:9", "NewWithTxOrFromCtx")
	sendErr := s.sms.SendRegisterSms(ctx, phone, code)
	if sendErr != nil {
		log.Error(sendErr)
		// 记录发送记录
		err := s.Dao.CreateMessageSms(ctx, &model.MessageSms{
			Model:     cmodel.Model{Id: id_util.GetUuid(), CreateTime: time.Now()},
			SendTime:  time.Now(),
			Name:      "",
			To:        phone,
			Message:   code,
			IsSuccess: false,
			Other:     sendErr.Error(),
		})
		if err != nil {
			log.Error(err)
			return errs.NewInternal().WithMsg("发送失败并保存失败").WithError(errors.New(err.Error() + sendErr.Error()))
		}
		return errs.NewInternal().WithMsg("短信发送失败").WithError(sendErr)
	} else {
		// 记录发送记录
		err := s.Dao.CreateMessageSms(ctx, &model.MessageSms{
			Model:     cmodel.Model{Id: id_util.GetUuid(), CreateTime: time.Now()},
			SendTime:  time.Now(),
			Name:      "",
			To:        phone,
			Message:   code,
			IsSuccess: true,
			Other:     "注册短信",
		})
		if err != nil {
			log.Error(err)
			return err
		}
		// 记录注册短信
		err = s.Dao.CreateMessageValidCode(ctx, &model.MessageValidCode{
			Model:    cmodel.Model{Id: id_util.GetUuid(), CreateTime: time.Now()},
			SendTime: time.Now(),
			PhoneNum: phone,
			Code:     code,
			Type:     model.MessageValidCodeTypeRegister,
		})
		if err != nil {
			log.Error(err)
			return err
		}
		return nil
	}
}
