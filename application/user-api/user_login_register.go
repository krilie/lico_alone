package user_api

import (
	"context"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/module/message/model"
	"time"
)

func (a *AppUser) UserLogin(ctx context.Context, phone, password, clientId string) (jwt string, err error) {
	return a.UserService.UserLogin(ctx, phone, password, clientId)
}

func (a *AppUser) UserRegister(ctx context.Context, phone, password, validCode, clientId string) error {
	code, err := a.Message.Dao.GetLastMessageValidCodeByPhoneNum(ctx, phone, model.MessageValidCodeTypeRegister)
	if err != nil {
		return err
	}
	if code == nil {
		return errs.ErrBadRequest.WithMsg("验证码无效")
	}
	if code.Code != validCode {
		return errs.ErrBadRequest.WithMsg("验证码错误")
	}
	if time.Now().Sub(code.SendTime) > 5*time.Minute {
		return errs.ErrBadRequest.WithMsg("验证码过期")
	}
	return a.UserService.RegisterNewUser(ctx, phone, password)
}
