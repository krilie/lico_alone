package service

import (
	"github.com/jinzhu/gorm"
	"github.com/krilie/lico_alone/common/comstruct/errs"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/module/account/model"
)

func (Account) DeleteBill(ctx *context.Context, billId string) error {
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
	bill.IsValid = false
	if e := tx.Update(&bill).Error; e != nil {
		tx.Rollback()
		return errs.ErrInternal.NewWithMsg(e.Error())
	}
	if e := tx.Model(&model.BillDetail{}).Where("bill_id=?", bill.Id).Update("is_valid", false).Error; e != nil {
		tx.Rollback()
		return errs.ErrInternal.NewWithMsg(e.Error())
	}
	tx.Commit()
	return nil
}
