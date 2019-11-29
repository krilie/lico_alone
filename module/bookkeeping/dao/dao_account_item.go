package dao

import (
	"context"
	"github.com/krilie/lico_alone/module/bookkeeping/model"
)

type IAccountItem interface {
	GetAccountItemById(ctx context.Context, id string) *model.AccountItem
	GetAccountItemListByUserId(ctx context.CancelFunc, userId string) []*model.AccountItem
	CreateAccountItem(ctx context.CancelFunc, item *model.AccountItem) error
	UpdateAccountItem(ctx context.CancelFunc, item *model.AccountItem) error
}
