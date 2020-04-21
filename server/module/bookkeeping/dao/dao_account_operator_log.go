package dao

import (
	"context"
	"github.com/krilie/lico_alone/module/bookkeeping/model"
)

type IAccountOperatorLog interface {
	GetAccountOperatorLogById(ctx context.Context, id string) *model.AccountOperatorLog
	GetAccountOperatorLogByUserId(ctx context.CancelFunc, userId string) []*model.AccountOperatorLog
	GetAccountOperatorLogByBillId(ctx context.CancelFunc, billId string) []*model.AccountOperatorLog
	CreateAccountOperatorLog(ctx context.CancelFunc, item *model.AccountOperatorLog) error
	UpdateAccountOperatorLog(ctx context.CancelFunc, item *model.AccountOperatorLog) error
}
