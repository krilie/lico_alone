package dao

import (
	"context"
	"errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/krilie/lico_alone/common/utils/id_util"
	"github.com/krilie/lico_alone/module/module-test/model"
	"gorm.io/gorm"
)

func (t *TestArticleDao) GetBaseOneById(ctx context.Context, id string) (*model.BaseOne, error) {
	baseOne := new(model.BaseOne)
	err := t.GetDb(ctx).First(baseOne, "id=?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return baseOne, err
}

func (t *TestArticleDao) DeleteBaseOneById(ctx context.Context, id string) error {
	return t.GetDb(ctx).Delete(&model.BaseOne{}).Where("id=?", id).Error
}

func (t *TestArticleDao) UpdateBaseOneById(ctx context.Context, baseOne *model.BaseOne) error {
	result := t.GetDb(ctx).Model(new(model.BaseOne)).Select("*").Updates(baseOne)
	return result.Error
}

func (t *TestArticleDao) CreateBaseOne(ctx context.Context, baseOne *model.BaseOne) error {
	if baseOne.Id == "" {
		baseOne.Id = id_util.GetUuid()
	}
	err := t.GetDb(ctx).Model(baseOne).Create(baseOne).Error
	return err
}

func (t *TestArticleDao) DeleteBaseOneById2(ctx context.Context, id string) error {
	sql, args, err := sq.Delete("tb_base_one").Where("id=?", id).ToSql()
	if err != nil {
		return err
	}
	t.log.Get(ctx).WithField("sql", sql).WithField("params", args).Debug("sql to exec")
	err = t.GetDb(ctx).Exec(sql, args...).Error
	return err
}

func (t *TestArticleDao) DeleteBaseOneById3(ctx context.Context, id string) error {
	sql := "delete from tb_article_master where id=?"
	t.log.Get(ctx).WithField("sql", sql).WithField("params", []interface{}{id}).Debug("sql to exec")
	err := t.GetDb(ctx).Exec(sql, id).Error
	return err
}
