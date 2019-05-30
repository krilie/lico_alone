package user

import (
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/module/account/model"
	"github.com/krilie/lico_alone/module/account/pojo"
	"github.com/krilie/lico_alone/module/account/service"
	"github.com/shopspring/decimal"
	"time"
)

var account service.Account

func (a AppUser) DeleteBill(ctx *context.Context, billId string, userId string) error {
	panic("implement me")
}

func (a AppUser) GetAccountHistory(ctx *context.Context, start, end time.Time, userId, AccountId, note string) ([]model.BillDetail, error) {
	panic("implement me")
}

func (a AppUser) GetAccountInfo(ctx *context.Context, userId string) ([]*pojo.AccountInfo, error) {
	panic("implement me")
}

func (a AppUser) AddBill(ctx *context.Context, userId, note, image string, amount decimal.Decimal, detail []pojo.BillDetail) (string, error) {
	panic("implement me")
}

func (a AppUser) AddAccount(ctx *context.Context, userId, name, num, description, image string, balance decimal.Decimal) (string, error) {
	panic("implement me")
}

func (a AppUser) DeleteAccount(ctx *context.Context, accountId string, userId string) error {
	panic("implement me")
}

func (a AppUser) GetMonthSummary(ctx *context.Context, userId string, time time.Time) (*pojo.AccountSummary, error) {
	panic("implement me")
}

func (a AppUser) GetTimeZoneSummary(ctx *context.Context, userId string, timeStart, timeEnd time.Time) (*pojo.AccountSummary, error) {
	panic("implement me")
}
