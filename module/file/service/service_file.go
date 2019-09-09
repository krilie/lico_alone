package service

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/krilie/lico_alone/common/cdb"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/common/utils/file_util"
	"github.com/krilie/lico_alone/common/utils/id_util"
	"github.com/krilie/lico_alone/module/file/model"
	"mime/multipart"
	"time"
)

// 内部有事务的存在
func (a *Service) UploadFile(ctx context.Context, tx *gorm.DB, userId, fileName string, file multipart.File, size int) (objName, contentType string, err error) {
	err = cdb.WithTrans(ctx, a, func(s cdb.Service) error {
		fileService := s.(*Service)
		// 生成文件记录
		content, err := file_util.GetContentType(file)
		if err != nil {
			return errs.NewInternal().WithError(err)
		}
		item := model.FileMaster{
			Id:          id_util.GetUuid(),
			CreateTime:  time.Now(),
			KeyName:     id_util.GetUuid(),
			BucketName:  fileService.Oss.GetBucketName(),
			UserId:      userId,
			ContentType: content,
			BizType:     "",
			Size:        size,
		}
		err = fileService.Dao.CreateFile(ctx, &item)
		if err != nil {
			return errs.NewErrDbCreate().WithError(err)
		}
		name, err := fileService.Oss.UploadFile(ctx, userId, fileName, content, file, int64(size))
		if err != nil {
			return err
		}
		item.KeyName = name
		err = fileService.Dao.SaveFile(ctx, &item)
		if err != nil {
			return err
		}
		contentType = content
		objName = name
		return nil
	})
	if err != nil {
		return "", "", err
	} else {
		return objName, contentType, nil
	}
}

// 内部有事务的存在
func (a *Service) DeleteFile(ctx context.Context, bucket, key string) (err error) {
	err = cdb.WithTrans(ctx, a, func(s cdb.Service) error {
		srv := s.(*Service)
		srv.Oss.del
		srv.Dao.DeleteFile()
	})
	return err
}
