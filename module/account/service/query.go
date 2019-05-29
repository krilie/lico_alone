package service

import (
	"github.com/krilie/lico_alone/common/comstruct/errs"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/module/account/model"
	"time"
)

// 获取一个账户的account信息
func (Account) GetAccountHistory(ctx *context.Context, start, end time.Time, userId, AccountId, note string) ([]model.BillDetail, error) {
	billDetails := make([]model.BillDetail, 0, 4)
	if e := model.Db.Where("create_time between ? and ? and account_id=?", start, end, AccountId).Find(billDetails).Error; e != nil {
		return nil, errs.ErrInternal.NewWithMsg(e.Error())
	}
	return billDetails, nil
}
