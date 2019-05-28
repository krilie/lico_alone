package service

import (
	"github.com/jinzhu/gorm"
	"github.com/krilie/lico_alone/common/comstruct/errs"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/id_util"
	"github.com/krilie/lico_alone/module/account/model"
	"github.com/krilie/lico_alone/module/account/pojo"
	"github.com/shopspring/decimal"
	"time"
)

func (Account) AddBill(ctx *context.Context, userId, note, image string, amount decimal.Decimal, detail []pojo.BillDetail) (string, error) {
	// 贷  借 借贷平衡
	credit, debit := decimal.Zero, decimal.Zero
	for _, v := range detail {
		credit = credit.Add(v.Credit)
		debit = debit.Add(v.Debit)
	}
	if !credit.Equal(debit) {
		return "", errs.ErrParam.NewWithMsg("credit and debit not equal.")
	}
	if !debit.Equal(amount) {
		return "", errs.ErrParam.NewWithMsg("debit and amount not equal.")
	}
	var bill model.Bill
	bill.Id = id_util.GetUuid()
	bill.Image = image
	bill.CreateTime = time.Now()
	bill.UserId = userId
	bill.Amount = amount
	bill.Note = note
	bill.IsValid = true
	// 插入到数据库中 事务开始
	tx := model.Db.Begin()
	if e := tx.Create(&bill).Error; e != nil {
		tx.Rollback()
		return "", errs.ErrInternal.NewWithMsg(e.Error())
	}
	for _, v := range detail {
		var detail model.BillDetail
		detail.Id = id_util.GetUuid()
		detail.BillId = bill.Id
		detail.UserId = userId
		detail.Debit = v.Debit
		detail.Credit = v.Credit
		detail.Note = v.Note
		// 查询account值
		var account model.Account
		if e := tx.Find(&account, "id=?", v.AccountId).Error; e != nil {
			tx.Rollback()
			if e == gorm.ErrRecordNotFound {
				return "", errs.ErrParam.NewWithMsg("no such account id:" + v.AccountId)
			} else {
				return "", errs.ErrInternal.NewWithMsg(e.Error())
			}
		}
		account.Credit = account.Credit.Add(v.Credit)
		account.Debit = account.Debit.Add(v.Debit)
		account.Balance = account.Balance.Add(v.Credit).Sub(v.Debit)
		if e := tx.Update(&account).Error; e != nil {
			tx.Rollback()
			return "", errs.ErrInternal.NewWithMsg(e.Error())
		}
		//  更新account
		detail.Balance = account.Balance
		detail.AccountId = v.AccountId
		detail.AccountNum = account.Num
		detail.AccountName = account.Name
		if e := tx.Create(&detail).Error; e != nil {
			tx.Rollback()
			return "", errs.ErrInternal.NewWithMsg(e.Error())
		}
	}
	tx.Commit()
	return bill.Id, nil
}
