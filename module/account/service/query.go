package service

import (
	"github.com/krilie/lico_alone/common/comstruct/errs"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/validator"
	"github.com/krilie/lico_alone/module/account/model"
	"github.com/krilie/lico_alone/module/account/pojo"
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

func (Account) GetAccountInfo(ctx *context.Context, userId string) ([]*pojo.AccountInfo, error) {
	if !validator.IsIdStr(userId) {
		return nil, errs.ErrParam.NewWithMsg("user_id 格式不正确")
	}
	infos := make([]model.Account, 0, 4)
	if e := model.Db.Where("user_id=?", userId).Find(infos).Error; e != nil {
		return nil, errs.ErrInternal.NewWithMsg(e.Error())
	}
	sumInfos := make([]*pojo.AccountInfo, 0, len(infos))
	for _, v := range infos {
		info := pojo.AccountInfo{Id: v.Id, Name: v.Name, Num: v.Num, Balance: v.Balance, Image: v.Image}
		sumInfos = append(sumInfos, &info)
	}
	return sumInfos, nil
}
