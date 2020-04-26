package service

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/common/utils/id_util"
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/krilie/lico_alone/module/file/model"
	"mime/multipart"
	"time"
)

// 内部有事务的存在
func (a *Service) UploadFile(ctx context.Context, tx *gorm.DB, userId, fileName string, file multipart.File, size int) (url, bucket, key string, err error) {
	log := nlog.NewLog(ctx, "module/file/service/service_broker.go:5", "RegisterBroker")
	err = cdb.WithTrans(ctx, a, func(ctx context.Context, s cdb.Service) error {
		fileService := s.(*Service)
		var content string
		content, bucket, key, err = fileService.FileSaver.UploadFile(ctx, userId, fileName, file, int64(size))
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
		err = fileService.Dao.CreateFile(ctx, &item)
		if err != nil {
			log.Error(err.Error())
			return errs.NewErrDbCreate().WithError(err)
		}
		url = fileService.FileSaver.GetFullUrl(ctx, true, key)
		return nil
	})
	return url, bucket, key, err
}

// 内部有事务的存在
func (a *Service) DeleteFile(ctx context.Context, bucket, key string) (err error) {
	err = cdb.WithTrans(ctx, a, func(ctx context.Context, s cdb.Service) error {
		srv := s.(*Service)
		err := srv.Dao.DeleteFileByBucketKey(ctx, bucket, key)
		if err != nil {
			return errs.NewErrDbDelete().WithError(err)
		}
		err = srv.FileSaver.DeleteFile(ctx, "", key)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (a *Service) GetBaseUrl(ctx context.Context) string {
	return a.FileSaver.GetBaseUrl(ctx)
}
