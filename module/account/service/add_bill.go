package service

import (
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/module/account/model"
	"github.com/shopspring/decimal"
)

func (Account) AddBill(ctx *context.Context, userId, note, image string, amount decimal.Decimal, detail []model.BillDetail) {

}
