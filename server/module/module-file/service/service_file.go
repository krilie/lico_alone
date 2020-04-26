package service

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/common/utils/id_util"
	"github.com/krilie/lico_alone/module/module-file/model"
	"github.com/prometheus/common/log"
	"mime/multipart"
	"time"
)

// 内部有事务的存在
func (a *FileService) UploadFile(ctx context.Context, tx *gorm.DB, userId, fileName string, file multipart.File, size int) (url, bucket, key string, err error) {
	err = a.dao.Transaction(ctx, func(ctx context.Context) error {
		var content string
		content, bucket, key, err = a.fileApi.UploadFile(ctx, userId, fileName, file, int64(size))
		if err != nil {
			return err
		}
		item := model.FileMaster{
			Id:          id_util.GetUuid(),
			CreateTime:  time.Now(),
			KeyName:     key,
			BucketName:  bucket,
			UserId:      userId,
			ContentType: content,
			BizType:     "",
			Size:        size,
		}
		err = a.dao.CreateFile(ctx, &item)
		if err != nil {
			log.Error(err.Error())
			return errs.NewErrDbCreate().WithError(err)
		}
		url = a.fileApi.GetFullUrl(ctx, true, key)
		return nil
	})
	return url, bucket, key, err
}

// 内部有事务的存在
func (a *FileService) DeleteFile(ctx context.Context, bucket, key string) (err error) {

	err = a.dao.Transaction(ctx, func(ctx context.Context) error {
		err := a.dao.DeleteFileByBucketKey(ctx, bucket, key)
		if err != nil {
			return errs.NewErrDbDelete().WithError(err)
		}
		err = a.fileApi.DeleteFile(ctx, "", key)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (a *FileService) GetBaseUrl(ctx context.Context) string {
	return a.fileApi.GetBaseUrl(ctx)
}
