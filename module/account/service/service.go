package service

import (
	"github.com/krilie/lico_alone/common/clog"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/module/account/model"
	"github.com/shopspring/decimal"
	"time"
)

var log = clog.NewLog(context.NewContext(), "alone.module.account.service", "init")

type Account struct{}

type Accounter interface {
	DeleteBill(ctx context.Context, billId string, userId string) error
	GetAccountHistory(ctx context.Context, start, end time.Time, userId, AccountId, note string) ([]model.BillDetail, error)
	GetAccountInfo(ctx context.Context, userId string) ([]*model.AccountInfo, error)
	AddBill(ctx context.Context, userId, note, image string, amount decimal.Decimal, detail []model.BillDetail) (string, error)
	AddAccount(ctx context.Context, userId, name, num, description, image string, balance decimal.Decimal) (string, error)
	DeleteAccount(ctx context.Context, accountId string, userId string) error
	GetMonthSummary(ctx context.Context, userId string, time time.Time) (*model.AccountSummary, error)
	GetTimeZoneSummary(ctx context.Context, userId string, timeStart, timeEnd time.Time) (*model.AccountSummary, error)
}
