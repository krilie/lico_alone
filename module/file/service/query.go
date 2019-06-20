package service

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/krilie/lico_alone/common/comstruct/errs"
	"github.com/krilie/lico_alone/module/file/model"
)

func (FileOp) QueryByFilePath(ctx context.Context, userId, filePath string) (*model.File, error) {
	var file = model.File{}
	e := model.Db.Find(&file, model.File{ObjKey: filePath}).Error
	if e == nil {
		return &file, nil
	} else if e == gorm.ErrRecordNotFound {
		return nil, nil
	} else {
		return nil, errs.ErrInternal.NewWithMsg(e.Error())
	}
}

func (FileOp) QueryById(ctx context.Context, userId, fileId string) (*model.File, error) {
	var file = model.File{}
	e := model.Db.Find(&file, model.File{ID: fileId}).Error
	if e == nil {
		return &file, nil
	} else if e == gorm.ErrRecordNotFound {
		return nil, nil
	} else {
		return nil, errs.ErrInternal.NewWithMsg(e.Error())
	}
}
