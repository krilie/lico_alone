package dao

import (
	"context"
	"errors"
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
	err := t.GetDb(ctx).Model(new(model.BaseOne)).Create(baseOne).Error
	return err
}
