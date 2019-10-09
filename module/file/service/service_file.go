package service

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/krilie/lico_alone/common/cdb"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/common/utils/id_util"
	"github.com/krilie/lico_alone/module/file/model"
	"mime/multipart"
	"time"
)

// 内部有事务的存在
func (a *Service) UploadFile(ctx context.Context, tx *gorm.DB, userId, fileName string, file multipart.File, size int) (url, bucket, key string, err error) {
	err = cdb.WithTrans(ctx, a, func(s cdb.Service) error {
		fileService := s.(*Service)
		item := model.FileMaster{
			Id:          id_util.GetUuid(),
			CreateTime:  time.Now(),
			KeyName:     "",
			BucketName:  "",
			UserId:      userId,
			ContentType: "",
			BizType:     "",
			Size:        size,
		}
		err = fileService.Dao.CreateFile(ctx, &item)
		if err != nil {
			return errs.NewErrDbCreate().WithError(err)
		}
		var content string
		content, bucket, key, err = fileService.FileSaver.UploadFile(ctx, userId, fileName, file, int64(size))
		if err != nil {
			return err
		}
		item.KeyName = key
		item.BucketName = bucket
		item.ContentType = content
		err = fileService.Dao.SaveFile(ctx, &item)
		if err != nil {
			return err
		}
		url = fileService.FileSaver.GetFullUrl(ctx, true, key)
		return nil
	})
	return url, bucket, key, err
}

// 内部有事务的存在
func (a *Service) DeleteFile(ctx context.Context, bucket, key string) (err error) {
	err = cdb.WithTrans(ctx, a, func(s cdb.Service) error {
		srv := s.(*Service)
		err := srv.Dao.DeleteFileByBucketKey(ctx, bucket, key)
		if err != nil {
			return errs.NewErrDbDelete().WithError(err)
		}
		err = srv.FileSaver.DeleteFile(ctx, "", "")
		if err != nil {
			return err
		}
		return nil
	})
	return err
}
