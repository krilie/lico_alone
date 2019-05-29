package service

import (
	"github.com/krilie/lico_alone/common/comstruct/errs"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/time_util"
	"github.com/krilie/lico_alone/common/validator"
	"github.com/krilie/lico_alone/module/account/model"
	"github.com/krilie/lico_alone/module/account/pojo"
	"github.com/shopspring/decimal"
	"time"
)

// 总结 统计

// 统计此用户月份的信息
func (a Account) GetMonthSummary(ctx *context.Context, userId string, time time.Time) (*pojo.AccountSummary, error) {
	// 月份的开始与结束
	monStart := time_util.GetBeijingMonthStartTime(time)
	monEnd := time_util.GetBeijingLastDateOfMonth(time)
	return a.GetTimeZoneSummary(ctx, userId, monStart, monEnd)
}

// 统计一个时区的信息
func (Account) GetTimeZoneSummary(ctx *context.Context, userId string, timeStart, timeEnd time.Time) (*pojo.AccountSummary, error) {
	// 参数格式检查
	if !validator.IsIdStr(userId) {
		return nil, errs.ErrParam.NewWithMsg("user id in err format")
	}
	summary := pojo.AccountSummary{}
	summary.BeginTime = timeStart
	summary.EndTime = timeEnd
	summary.Accounts = make([]pojo.AccountItem, 0, 4)
	summary.Bills = make([]pojo.BillItem, 0, 4)
	// 查询账户
	accounts := make([]model.Account, 0, 4)
	if e := model.Db.Find(accounts, "user_id=?", userId).Error; e != nil {
		return nil, errs.ErrInternal.NewWithMsg(e.Error())
	}
	// 帐户统计 accounts
	for _, v := range accounts {
		details := make([]model.BillDetail, 0, 4)
		if e := model.Db.Where("create_time between ? and ? and is_valid=true and account_id=?", timeStart, timeEnd, v.Id).Find(details).Error; e != nil {
			return nil, errs.ErrInternal.NewWithMsg(e.Error())
		}
		item := pojo.AccountItem{Id: v.Id, Name: v.Name, Num: v.Num, Credit: v.Credit, Debit: v.Debit, Amount: decimal.Zero, NowBalance: v.Balance}
		for _, v := range details {
			item.Amount = item.Amount.Add(v.Credit)
		}
		summary.Accounts = append(summary.Accounts, item)
	}
	// 帐户统计 bills
	bills := make([]model.Bill, 0, 4)
	if e := model.Db.Where("create_time between ? and ? and user_id=?", timeStart, timeEnd, userId).Find(bills).Error; e != nil {
		return nil, errs.ErrInternal.NewWithMsg(e.Error())
	}
	for _, v := range bills {
		tempBill := pojo.BillItem{Id: v.Id, CreateTime: v.CreateTime, Note: v.Note, Amount: v.Amount, Image: v.Image, IsValid: v.IsValid}
		summary.Bills = append(summary.Bills, tempBill)
	}
	return &summary, nil
}

// 获取一个账户的account信息
func (Account) GetAccountHistory(ctx *context.Context, start, end time.Time, userId, AccountId, note string) ([]model.BillDetail, error) {
	billDetails := make([]model.BillDetail, 0, 4)
	if e := model.Db.Where("create_time between ? and ? and account_id=?", start, end, AccountId).Find(billDetails).Error; e != nil {
		return nil, errs.ErrInternal.NewWithMsg(e.Error())
	}
	return billDetails, nil
}
