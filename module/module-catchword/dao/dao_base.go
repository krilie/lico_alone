package dao

import (
	"context"
	"errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/krilie/lico_alone/common/utils/id_util"
	"github.com/krilie/lico_alone/module/module-catchword/model"
	"gorm.io/gorm"
)

func (t *CatchwordDao) GetCatchwordById(ctx context.Context, id string) (*model.Catchword, error) {
	catchword := new(model.Catchword)
	err := t.GetDb(ctx).First(catchword, "id=?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return catchword, err
}

func (t *CatchwordDao) DeleteCatchwordById(ctx context.Context, id string) error {
	return t.GetDb(ctx).Where("id=?", id).Delete(&model.Catchword{}).Error
}

func (t *CatchwordDao) UpdateCatchwordById(ctx context.Context, catchword *model.Catchword) error {
	result := t.GetDb(ctx).Model(new(model.Catchword)).Select("*").Omit("create_at", "delete_at").Where("id=?", catchword.Id).Updates(catchword)
	return result.Error
}

func (t *CatchwordDao) CreateCatchword(ctx context.Context, catchword *model.Catchword) error {
	if catchword.Id == "" {
		catchword.Id = id_util.GetUuid()
	}
	err := t.GetDb(ctx).Model(catchword).Create(catchword).Error
	return err
}

func (t *CatchwordDao) DeleteCatchwordById2(ctx context.Context, id string) error {
	sql, args, err := sq.Delete("tb_catchword").Where("id=?", id).ToSql()
	if err != nil {
		return err
	}
	t.log.Get(ctx).WithField("sql", sql).WithField("params", args).Debug("sql to exec")
	err = t.GetDb(ctx).Exec(sql, args...).Error
	return err
}

func (t *CatchwordDao) DeleteCatchwordById3(ctx context.Context, id string) error {
	sql := "delete from tb_catchword where id=?"
	t.log.Get(ctx).WithField("sql", sql).WithField("params", []interface{}{id}).Debug("sql to exec")
	err := t.GetDb(ctx).Exec(sql, id).Error
	return err
}
