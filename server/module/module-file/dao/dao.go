package dao

import (
	"context"
	context_enum "github.com/krilie/lico_alone/common/com-model/context-enum"
	context2 "github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/component/ndb"
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/krilie/lico_alone/module/module-file/model"
)

type FileDao struct {
	*ndb.NDb
	log *nlog.NLog
}

func NewFileDao(db *ndb.NDb, log *nlog.NLog) *FileDao {
	log = log.WithField(context_enum.Module.Str(), "module file dao")
	err := db.GetDb(context2.NewContext()).
		AutoMigrate(&model.FileMaster{}).Error
	if err != nil {
		panic(err)
	}
	return &FileDao{
		NDb: db,
		log: log,
	}
}

func (a *FileDao) CreateFile(ctx context.Context, file *model.FileMaster) error {
	err := a.GetDb(ctx).Create(file).Error
	if err != nil {
		return errs.NewInternal().WithError(err)
	}
	return nil
}

func (a *FileDao) SaveFile(ctx context.Context, file *model.FileMaster) error {
	err := a.GetDb(ctx).Save(file).Error
	if err != nil {
		return errs.NewInternal().WithError(err)
	}
	return nil
}

func (a *FileDao) UpdateFile(ctx context.Context, file *model.FileMaster) error {
	err := a.GetDb(ctx).Update(file).Error
	if err != nil {
		return errs.NewInternal().WithError(err)
	}
	return nil
}

func (a *FileDao) DeleteFile(ctx context.Context, id string) error {
	err := a.GetDb(ctx).Where("id=?", id).Delete(&model.FileMaster{}).Error
	if err != nil {
		return errs.NewInternal().WithError(err)
	}
	return nil
}

func (a *FileDao) DeleteFileByBucketKey(ctx context.Context, bucket, key string) error {
	result := a.GetDb(ctx).Where("bucket_name=? and key_name=?", bucket, key).Delete(&model.FileMaster{})
	err := result.Error
	if err != nil {
		return errs.NewInternal().WithError(err)
	}
	if result.RowsAffected == 0 {
		return errs.NewNormal().WithMsg("删除没有成功")
	}
	return nil
}
