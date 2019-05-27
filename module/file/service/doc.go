package service

import (
	"github.com/krilie/lico_alone/common/config"
	"github.com/krilie/lico_alone/common/context"
	"mime/multipart"
)

var ossKey = config.GetString("oss.key")
var ossSecret = config.GetString("oss.secret")
var ossEndPoint = config.GetString("oss.endpoint")
var ossBucket = config.GetString("oss.bucket")

type FileOp struct{}
type FileUploadDeleter interface {
	UploadFile(ctx *context.Context, userId, fileName string, file multipart.File, size int64) (string, error)
	DeleteFile(ctx *context.Context, userId, filePath string) error
}
