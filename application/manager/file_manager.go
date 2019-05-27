package manager

import (
	"github.com/krilie/lico_alone/common/context_util"
	"mime/multipart"
)

func (AppManager) UploadFile(ctx *context_util.Context, userId, fileName string, file multipart.File, size int64) (string, error) {
	panic("implement me")
}

func (AppManager) DeleteFile(ctx *context_util.Context, userId, filePath string) error {
	panic("implement me")
}
