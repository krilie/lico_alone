package service

import (
	"context"
	"errors"
	"github.com/krilie/lico_alone/common/com-model"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/common/utils/id_util"
	"github.com/krilie/lico_alone/module/module-message/model"
	"github.com/prometheus/common/log"
	"time"
)

// SendRegisterSms 发送注册短信
func (s *MessageModule) SendRegisterSms(ctx context.Context, phone, code string) error {
	sendErr := s.sms.SendRegisterSms(ctx, phone, code)
	if sendErr != nil {
		s.log.Get(ctx).Error(sendErr)
		// 记录发送记录
		err := s.Dao.CreateMessageSms(ctx, &model.MessageSms{
			Model: com_model.Model{
				Id:        id_util.GetUuid(),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: nil,
			},
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
			Model: com_model.Model{
				Id:        id_util.GetUuid(),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: nil,
			},
			SendTime:  time.Now(),
			Name:      "",
			To:        phone,
			Message:   code,
			IsSuccess: true,
			Other:     "注册短信",
		})
		if err != nil {
			s.log.Get(ctx).Error(err)
			return err
		}
		// 记录注册短信
		err = s.Dao.CreateMessageValidCode(ctx, &model.MessageValidCode{
			Model: com_model.Model{
				Id:        id_util.GetUuid(),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: nil,
			},
			SendTime: time.Now(),
			PhoneNum: phone,
			Code:     code,
			Type:     model.MessageValidCodeTypeRegister.ToInt(),
		})
		if err != nil {
			s.log.Get(ctx).Error(err)
			return err
		}
		return nil
	}
}
