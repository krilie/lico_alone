package user

import (
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/module/account/model"
	"github.com/krilie/lico_alone/module/account/pojo"
	"github.com/krilie/lico_alone/module/account/service"
	"github.com/shopspring/decimal"
	"time"
)

type Account struct{}

var account service.Account

func (a Account) DeleteBill(ctx context.Context, billId string, userId string) error {
	return account.DeleteBill(ctx, billId, userId)
}

func (a Account) GetAccountHistory(ctx context.Context, start, end time.Time, userId, AccountId, note string) ([]model.BillDetail, error) {
	return account.GetAccountHistory(ctx, start, end, userId, AccountId, note)
}

func (a Account) GetAccountInfo(ctx context.Context, userId string) ([]*pojo.AccountInfo, error) {
	return account.GetAccountInfo(ctx, userId)
}

func (a Account) AddBill(ctx context.Context, userId, note, image string, amount decimal.Decimal, detail []pojo.BillDetail) (string, error) {
	return account.AddBill(ctx, userId, note, image, amount, detail)
}

func (a Account) AddAccount(ctx context.Context, userId, name, num, description, image string, balance decimal.Decimal) (string, error) {
	return account.AddAccount(ctx, userId, name, num, description, image, balance)
}

func (a Account) DeleteAccount(ctx context.Context, accountId string, userId string) error {
	return account.DeleteAccount(ctx, accountId, userId)
}

func (a Account) GetMonthSummary(ctx context.Context, userId string, time time.Time) (*pojo.AccountSummary, error) {
	return account.GetMonthSummary(ctx, userId, time)
}

func (a Account) GetTimeZoneSummary(ctx context.Context, userId string, timeStart, timeEnd time.Time) (*pojo.AccountSummary, error) {
	return account.GetTimeZoneSummary(ctx, userId, timeStart, timeEnd)
}
