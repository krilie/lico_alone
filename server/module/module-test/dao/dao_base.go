package dao

import (
	"context"
	"errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/krilie/lico_alone/common/utils/id_util"
	"github.com/krilie/lico_alone/module/module-test/model"
	"gorm.io/gorm"
)

func (t *TestArticleDao) GetOneById(ctx context.Context, id string) (*model.One, error) {
	one := new(model.One)
	err := t.GetDb(ctx).First(one, "id=?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return one, err
}

func (t *TestArticleDao) DeleteOneById(ctx context.Context, id string) error {
	return t.GetDb(ctx).Delete(&model.One{}).Where("id=?", id).Error
}

func (t *TestArticleDao) UpdateOneById(ctx context.Context, one *model.One) error {
	result := t.GetDb(ctx).Model(new(model.One)).Select("*").Updates(one)
	return result.Error
}

func (t *TestArticleDao) CreateOne(ctx context.Context, one *model.One) error {
	if one.Id == "" {
		one.Id = id_util.GetUuid()
	}
	err := t.GetDb(ctx).Model(one).Create(one).Error
	return err
}

func (t *TestArticleDao) DeleteOneById2(ctx context.Context, id string) error {
	sql, args, err := sq.Delete("tb_one").Where("id=?", id).ToSql()
	if err != nil {
		return err
	}
	t.log.Get(ctx).WithField("sql", sql).WithField("params", args).Debug("sql to exec")
	err = t.GetDb(ctx).Exec(sql, args...).Error
	return err
}

func (t *TestArticleDao) DeleteOneById3(ctx context.Context, id string) error {
	sql := "delete from tb_one where id=?"
	t.log.Get(ctx).WithField("sql", sql).WithField("params", []interface{}{id}).Debug("sql to exec")
	err := t.GetDb(ctx).Exec(sql, id).Error
	return err
}

func (t *TestArticleDao) GetTwoById(ctx context.Context, id string) (*model.Two, error) {
	two := new(model.Two)
	err := t.GetDb(ctx).First(two, "id=?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return two, err
}

func (t *TestArticleDao) DeleteTwoById(ctx context.Context, id string) error {
	return t.GetDb(ctx).Delete(&model.Two{}).Where("id=?", id).Error
}

func (t *TestArticleDao) UpdateTwoById(ctx context.Context, two *model.Two) error {
	result := t.GetDb(ctx).Model(new(model.Two)).Select("*").Updates(two)
	return result.Error
}

func (t *TestArticleDao) CreateTwo(ctx context.Context, two *model.Two) error {
	if two.Id == "" {
		two.Id = id_util.GetUuid()
	}
	err := t.GetDb(ctx).Model(two).Create(two).Error
	return err
}

func (t *TestArticleDao) DeleteTwoById2(ctx context.Context, id string) error {
	sql, args, err := sq.Delete("tb_two").Where("id=?", id).ToSql()
	if err != nil {
		return err
	}
	t.log.Get(ctx).WithField("sql", sql).WithField("params", args).Debug("sql to exec")
	err = t.GetDb(ctx).Exec(sql, args...).Error
	return err
}

func (t *TestArticleDao) DeleteTwoById3(ctx context.Context, id string) error {
	sql := "delete from tb_two where id=?"
	t.log.Get(ctx).WithField("sql", sql).WithField("params", []interface{}{id}).Debug("sql to exec")
	err := t.GetDb(ctx).Exec(sql, id).Error
	return err
}
