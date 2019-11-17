package user_api

import (
	"context"
	"github.com/krilie/lico_alone/common/clog"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/common/utils/random"
)

func (a *AppUser) SendRegisterSms(ctx context.Context, phoneNum string) error {
	if phoneNum == "" {
		return errs.NewBadRequest().WithMsg("手机号格式不正确")
	}
	master, err := a.UserService.Dao.GetUserMasterByPhoneNum(ctx, phoneNum)
	if err != nil {
		clog.With(ctx).Error(err)
		return err
	}
	if master != nil {
		return errs.NewBadRequest().WithMsg("已经注册")
	}
	err = a.Message.SendRegisterSms(ctx, phoneNum, random.GetRandomNum(5))
	return err
}
