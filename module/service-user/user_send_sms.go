package service_user

import (
	"context"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/common/utils/random"
)

func (a *UserService) SendRegisterSms(ctx context.Context, phoneNum string) error {
	if phoneNum == "" {
		return errs.NewParamError().WithMsg("手机号格式不正确")
	}
	master, err := a.moduleUser.Dao.GetUserMasterByPhoneNum(ctx, phoneNum)
	if err != nil {
		a.log.Get(ctx).Error(err)
		return err
	}
	if master != nil {
		return errs.NewNormal().WithMsg("已经注册")
	}
	err = a.moduleMsg.SendRegisterSms(ctx, phoneNum, random.GetRandomNum(5))
	return err
}
