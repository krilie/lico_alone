package service

import (
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/id_util"
	"github.com/krilie/lico_alone/module/account/model"
	"github.com/shopspring/decimal"
	"time"
)

func (Account) AddBill(ctx *context.Context, userId, note, image string, amount decimal.Decimal, detail []model.BillDetail) (string, error) {
	// 借贷平衡

	var bill model.Bill
	bill.Id = id_util.GetUuid()
	bill.Image = image
	bill.CreateTime = time.Now()
	bill.UserId = userId
	bill.Amount = amount
	bill.Note = note

	for k, v := range detail {

	}

	tx := model.Db.Begin()

	return "", nil
}
