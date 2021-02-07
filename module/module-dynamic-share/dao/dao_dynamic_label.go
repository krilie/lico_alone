package dao

import (
	"context"
	"errors"
	"github.com/ahmetb/go-linq/v3"
	com_model "github.com/krilie/lico_alone/common/com-model"
	"github.com/krilie/lico_alone/module/module-dynamic-share/model"
	"time"
)

func (a *DynamicShareDao) GetAllLabels(ctx context.Context) (*[]model.DynamicShareLabel, error) {
	data := make([]model.DynamicShareLabel, 0)
	err := a.GetDb(ctx).Model(&model.DynamicShareLabel{}).Find(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (a *DynamicShareDao) GetLabelsByShareIds(ctx context.Context, ids []string) (*[]model.DynamicShareLabel, error) {
	data := make([]model.DynamicShareLabel, 0)
	err := a.GetDb(ctx).Model(&model.DynamicShareLabel{}).Where("id in (?)", ids).Find(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (a *DynamicShareDao) AddLabels(ctx context.Context, labels []model.CreateDynamicShareLabelModel) error {
	adds := linq.From(labels).SelectT(func(o model.CreateDynamicShareLabelModel) interface{} {
		return model.DynamicShareLabel{
			Model:   com_model.NewModel(),
			ShareId: o.ShareId,
			Label:   o.Label,
		}
	}).Results()
	err := a.GetDb(ctx).Create(adds).Error
	return err
}

func (a *DynamicShareDao) DeleteLabels(ctx context.Context, ids []string) error {
	var sql = "update tb_dynamic_share_label set deleted_at=? where id in (?)"
	result := a.GetDb(ctx).Exec(sql, time.Now(), ids)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected <= 0 {
		return errors.New("no affected rows")
	}
	return nil
}
