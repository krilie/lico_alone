package service

import (
	"context"
	"github.com/krilie/lico_alone/common/comstruct/errs"
	"github.com/krilie/lico_alone/common/utils/id_util"
	"github.com/krilie/lico_alone/module/account/model"
	"github.com/shopspring/decimal"
	"time"
)

func (Account) AddAccount(ctx context.Context, userId, name, num, description, image string, balance decimal.Decimal) (string, error) {
	var account model.Account
	account.Id = id_util.NextSnowflakeId().String()
	account.UserId = userId
	account.Name = name
	account.CreateTime = time.Now()
	account.UpdateTime = account.CreateTime
	account.Num = num
	account.Credit = decimal.Zero
	account.Debit = decimal.Zero
	account.Balance = balance
	account.Description = description
	account.Image = image
	if e := model.Db.Create(&account).Error; e != nil {
		log.Error("add account ", e)
		return "", errs.ErrInternal.NewWithMsg(e.Error())
	}
	return account.Id, nil
}
