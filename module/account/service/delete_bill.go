package service

import (
	"github.com/jinzhu/gorm"
	"github.com/krilie/lico_alone/common/comstruct/errs"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/module/account/model"
	"time"
)

func (Account) DeleteBill(ctx *context.Context, billId string, userId string) error {
	// 标记删除
	tx := model.Db.Begin()
	var bill model.Bill
	if e := tx.Find(&bill, "id=?", billId).Error; e != nil {
		tx.Rollback()
		if e == gorm.ErrRecordNotFound {
			return errs.ErrNotFound.NewWithMsg("not find this bill")
		} else {
			return errs.ErrInternal.NewWithMsg(e.Error())
		}
	}
	if !bill.IsValid {
		tx.Rollback()
		return errs.ErrParam.NewWithMsg("bill is deleted")
	}
	// 更新 bill
	bill.IsValid = false
	if e := tx.Update(&bill).Error; e != nil {
		tx.Rollback()
		return errs.ErrInternal.NewWithMsg(e.Error())
	}
	// 更新bill detail
	if e := tx.Model(&model.BillDetail{}).Where("bill_id=?", bill.Id).Update("is_valid", false).Error; e != nil {
		tx.Rollback()
		return errs.ErrInternal.NewWithMsg(e.Error())
	}
	// 撤销对账户的操作
	billDetails := make([]model.BillDetail, 0, 4)
	if e := tx.Model(&model.BillDetail{}).Find(billDetails, "bill_id=?", bill.Id).Error; e != nil {
		tx.Rollback()
		return errs.ErrInternal.NewAppendMsg(e.Error())
	}
	for _, v := range billDetails {
		// 查询account
		var account model.Account
		if e := tx.Find(&account, "id=?", v.AccountId).Error; e != nil {
			tx.Rollback()
			if gorm.IsRecordNotFoundError(e) {
				return errs.ErrParam.NewWithMsg(e.Error() + " no this account: " + v.AccountId)
			} else {
				return errs.ErrInternal.NewWithMsg(e.Error())
			}
		}
		// 更新account的值
		account.Credit = account.Credit.Sub(v.Credit)
		account.Debit = account.Debit.Sub(v.Debit)
		account.Balance = account.Balance.Sub(v.Credit).Add(v.Debit)
		account.UpdateTime = time.Now()
		if e := tx.Save(account).Error; e != nil {
			tx.Rollback()
			return errs.ErrInternal.NewWithMsg(e.Error())
		}
	}
	tx.Commit()
	return nil
}
