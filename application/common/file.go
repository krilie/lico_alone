package common

import (
	"github.com/krilie/lico_alone/common/common_struct/errs"
	"github.com/krilie/lico_alone/common/context_util"
	"github.com/krilie/lico_alone/module/file/service"
	user2 "github.com/krilie/lico_alone/module/userbase/user"
	"mime/multipart"
)

var fileOp = service.FileOp{}
var user = user2.User{}

func (AppCommon) UploadFile(ctx *context_util.Context, userId, fileName string, file multipart.File, size int64) (string, error) {
	// 用户是否存在,没有e即取到了
	_, e := user.GetInfo(ctx, userId)
	if e != nil {
		return "", e
	}
	return fileOp.UploadFile(ctx, userId, fileName, file, size)
}

func (AppCommon) DeleteFile(ctx *context_util.Context, userId, filePath string) error {
	// 用户是否存在,没有e即取到了
	_, e := user.GetInfo(ctx, userId)
	if e != nil {
		return e
	}
	file, e := fileOp.QueryByFilePath(ctx, userId, filePath)
	if e != nil {
		return e
	}
	if file.UserId != userId {
		return errs.ErrNoPermission.NewWithMsg("can not find file you gaved")
	}
	return fileOp.DeleteFile(ctx, userId, filePath)
}
