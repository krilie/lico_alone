package user_service

import (
	"context"
	"github.com/krilie/lico_alone/common/errs"
	model2 "github.com/krilie/lico_alone/module/module-message/model"
	"time"
)

func (a *UserService) UserLogin(ctx context.Context, phone, password, clientId string) (jwt string, err error) {
	return a.moduleUser.UserLogin(ctx, phone, password, clientId)
}

func (a *UserService) UserRegister(ctx context.Context, phone, password, validCode, clientId string) error {
	code, err := a.moduleMsg.Dao.GetLastMessageValidCodeByPhoneNum(ctx, phone, model2.MessageValidCodeTypeRegister)
	if err != nil {
		return err
	}
	if code == nil {
		return errs.NewNormal().WithMsg("验证码无效")
	}
	if code.Code != validCode {
		return errs.NewNormal().WithMsg("验证码错误")
	}
	if time.Now().Sub(code.SendTime) > 5*time.Minute {
		return errs.NewNormal().WithMsg("验证码过期")
	}
	return a.moduleUser.RegisterNewUser(ctx, phone, password)
}
