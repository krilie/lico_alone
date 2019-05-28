package service

import (
	"github.com/jinzhu/gorm"
	"github.com/krilie/lico_alone/common/comstruct/errs"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/module/account/model"
)

func (Account) DeleteBill(ctx *context.Context, billId string) error {
	// 标记删除
	var bill model.Bill
	if e := model.Db.Find(&bill, "id=?", billId).Error; e != nil {
		if e == gorm.ErrRecordNotFound {
			return errs.ErrNotFound.NewWithMsg("not find this bill")
		} else {
			return errs.ErrInternal.NewWithMsg(e.Error())
		}
	}
	if !bill.IsValid {
		return errs.ErrParam.NewWithMsg("bill is deleted")
	}
	bill.IsValid = false
	if e := model.Db.Update(bill).Error; e != nil {
		return errs.ErrInternal.NewWithMsg(e.Error())
	}
	return nil
}
