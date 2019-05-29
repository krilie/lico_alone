package service

import (
	"github.com/krilie/lico_alone/common/comstruct/errs"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/time_util"
	"github.com/krilie/lico_alone/common/validator"
	"github.com/krilie/lico_alone/module/account/model"
	"time"
)

// 总结 统计

// 统计此用户月份的信息
func (a Account) GetMonthSummary(ctx *context.Context, userId string, time time.Time) (string, error) {
	// 月份的开始与结束
	monStart := time_util.GetBeijingMonthStartTime(time)
	monEnd := time_util.GetBeijingLastDateOfMonth(time)
	return a.GetTimeZoneSummary(ctx, userId, monStart, monEnd)
}

func (Account) GetTimeZoneSummary(ctx *context.Context, userId string, timeStart, timeEnd time.Time) (string, error) {
	// 参数格式检查
	if !validator.IsIdStr(userId) {
		return "", errs.ErrParam.NewWithMsg("user id in err format")
	}
	// 查询账户
	var accounts []Account
	if e := model.Db.Find(accounts, "user_id=?", userId).Error; e != nil {
		return "", errs.ErrInternal.NewWithMsg(e.Error())
	}
	// 帐户统计

}
