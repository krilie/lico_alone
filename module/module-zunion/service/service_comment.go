package service

import (
	"context"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/module/module-zunion/model"
	"time"
)

func (a *ZUnionModule) AddComment(ctx context.Context, comment *model.TbComment) error {
	err := a.Dao.GetDb(ctx).Model(&model.TbComment{}).Create(comment).Error
	if err != nil {
		a.log.Errorf("%v", err.Error())
		return errs.NewInternal().WithMsg(err.Error())
	}
	return nil
}

func (a *ZUnionModule) DeleteComment(ctx context.Context, id string) error {
	exec := a.Dao.GetDb(ctx).Exec("update tb_comment set deleted_at=? where id=?", time.Now(), id)
	if exec.Error != nil {
		return errs.NewInternal().WithMsg(exec.Error.Error())
	}
	return nil
}

func (a *ZUnionModule) QueryComment(ctx context.Context, targetId string) ([]*model.TbComment, error) {
	panic("not implement")
}
