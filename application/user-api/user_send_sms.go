package user_api

import (
	"context"
	"github.com/krilie/lico_alone/common/clog"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/common/utils/id_util"
)

func (a *AppUser) SendRegisterSms(ctx context.Context, phoneNum string) error {
	master, err := a.UserService.Dao.GetUserMasterByPhoneNum(ctx, phoneNum)
	if err != nil {
		clog.With(ctx).Error(err)
		return err
	}
	if master != nil {
		return errs.NewBadRequest().WithMsg("已经注册")
	}
	err = a.Message.SendRegisterSms(ctx, phoneNum, id_util.NextSnowflake())
	return err
}
