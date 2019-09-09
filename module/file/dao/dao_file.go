package dao

import (
	"context"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/module/file/model"
)

func (a *Dao) CreateFile(ctx context.Context, file *model.FileMaster) error {
	err := a.Db.Create(file).Error
	if err != nil {
		return errs.ErrDbCreate.WithError(err)
	}
	return nil
}

func (a *Dao) SaveFile(ctx context.Context, file *model.FileMaster) error {
	err := a.Db.Save(file).Error
	if err != nil {
		return errs.ErrDbUpdate.WithError(err)
	}
	return nil
}

func (a *Dao) UpdateFile(ctx context.Context, file *model.FileMaster) error {
	err := a.Db.Update(file).Error
	if err != nil {
		return errs.ErrDbUpdate.WithError(err)
	}
	return nil
}

func (a *Dao) DeleteFile(ctx context.Context, id string) error {
	err := a.Db.Where("id=?", id).Delete(&model.FileMaster{}).Error
	if err != nil {
		return errs.ErrDbDelete.WithError(err)
	}
	return nil
}
