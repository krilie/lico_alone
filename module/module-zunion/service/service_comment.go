package service

import (
	"context"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/module/module-zunion/model"
)

func (a *ZUnionModule) AddComment(ctx context.Context, comment *model.TbComment) error {
	err := a.Dao.GetDb(ctx).Model(&model.TbComment{}).Create(comment).Error
	if err != nil {
		a.log.Errorf("%v", err.Error())
		return errs.NewInternal().WithMsg(err.Error())
	}
	return nil
}
