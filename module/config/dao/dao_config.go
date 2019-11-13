package dao

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/module/config/model"
)

func (a *Dao) GetConfigByName(ctx context.Context, name string) (*model.Config, error) {
	config := new(model.Config)
	err := a.Db.Where("name=?", name).Find(config).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, errs.NewErrDbQuery().WithError(err)
	}
	return config, nil
}

func (a *Dao) CreateConfig(ctx context.Context, config *model.Config) error {
	err := a.Db.Create(config).Error
	if err != nil {
		return errs.NewErrDbCreate().WithError(err)
	}
	return nil
}

func (a *Dao) DeleteConfig(ctx context.Context, name string) error {
	err := a.Db.Where("name=?", name).Delete(new(model.Config)).Error
	if err != nil {
		return errs.NewErrDbDelete().WithError(err)
	}
	return nil
}
func (a *Dao) UpdateConfig(ctx context.Context, config *model.Config) error {
	err := a.Db.Model(config).Where("name=?", config.Name).Omit("create_time").Update(config).Error
	if err != nil {
		return errs.NewErrDbUpdate().WithError(err)
	}
	return nil
}
