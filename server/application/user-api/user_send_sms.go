package user_api

import (
	"context"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/common/utils/random"
	"github.com/krilie/lico_alone/component/nlog"
)

func (a *AppUser) SendRegisterSms(ctx context.Context, phoneNum string) error {
	if phoneNum == "" {
		return errs.NewBadRequest().WithMsg("手机号格式不正确")
	}
	master, err := a.UserService.Dao.GetUserMasterByPhoneNum(ctx, phoneNum)
	if err != nil {
		nlog.With(ctx).Error(err)
		return err
	}
	if master != nil {
		return errs.NewBadRequest().WithMsg("已经注册")
	}
	err = a.Message.SendRegisterSms(ctx, phoneNum, random.GetRandomNum(5))
	return err
}
