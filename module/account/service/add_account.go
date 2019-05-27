package service

import (
	"github.com/krilie/lico_alone/common/comstruct/errs"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/id_util"
	"github.com/krilie/lico_alone/common/log"
	"github.com/krilie/lico_alone/module/account/model"
	"github.com/shopspring/decimal"
	"time"
)

func (Account) AddAccount(ctx *context.Context, userId, name, description, image string, balance decimal.Decimal) (string, error) {
	var account model.Account
	account.Id = id_util.NextSnowflakeId().String()
	account.UserId = userId
	account.Name = name
	account.CreateTime = time.Now()
	account.UpdateTime = account.CreateTime
	account.Num = account.Id
	account.Balance = balance
	account.Description = description
	account.Image = image
	if e := model.Db.Create(&account).Error; e != nil {
		log.Error("add account ", e)
		return "", errs.ErrInternal.NewWithMsg(e.Error())
	}
	return account.Id, nil
}
