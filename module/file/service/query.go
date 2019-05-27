package service

import (
	"github.com/jinzhu/gorm"
	"github.com/krilie/lico_alone/common/comstruct"
	"github.com/krilie/lico_alone/common/comstruct/errs"
	"github.com/krilie/lico_alone/common/context_util"
	"github.com/krilie/lico_alone/module/file/model"
)

func (FileOp) QueryByFilePath(ctx *context_util.Context, userId, filePath string) (*model.File, error) {
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

func (FileOp) QueryById(ctx *context_util.Context, userId, fileId string) (*model.File, error) {
	var file = model.File{}
	e := model.Db.Find(&file, model.File{DbHandler: comstruct.DbHandler{ID: fileId}}).Error
	if e == nil {
		return &file, nil
	} else if e == gorm.ErrRecordNotFound {
		return nil, nil
	} else {
		return nil, errs.ErrInternal.NewWithMsg(e.Error())
	}
}
