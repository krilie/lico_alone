package service

import (
	"context"
	"github.com/krilie/lico_alone/common/comstruct/errs"
	"github.com/krilie/lico_alone/module/account/model"
)

func (Account) DeleteAccount(ctx context.Context, accountId string, userId string) error {
	// 查看是否已经使用
	var count int
	model.Db.Table("tb_bill_detail").Where("account_id = ?", accountId).Count(&count)
	if count == 0 {
		e := model.Db.Where("id=?", accountId).Delete(&model.Account{}).Error
		if e != nil {
			return errs.ErrInternal.NewWithMsg(e.Error())
		} else {
			return nil
		}
	} else {
		return errs.ErrParam.NewWithMsg("in use can not delete id:" + accountId)
	}
}
