package service

import (
	"context"
	"github.com/krilie/lico_alone/common/com-model"
	context_enum "github.com/krilie/lico_alone/common/com-model/context-enum"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/common/utils/file_util"
	"github.com/krilie/lico_alone/common/utils/id_util"
	"github.com/krilie/lico_alone/module/module-file/model"
	"gorm.io/gorm"
	"io"
	"mime"
	"time"
)

// 内部有事务的存在
func (a *FileModule) UploadFile(ctx context.Context, userId, fileName string, file io.Reader, size int) (url, bucket, key string, err error) {
	log := a.log.Get(ctx).WithField(context_enum.Function.Str(), "UploadFile")
	err = a.dao.Transaction(ctx, func(ctx context.Context) error {
		var content string
		extension := file_util.GetFileExtension(fileName)
		content1 := mime.TypeByExtension(extension)
		content2, newReader, err := file_util.GetContentType2(file)
		if err != nil {
			panic(err)
		}
		if content1 != "" {
			content = content1
		} else {
			content = content2
		}
		url, key, err = a.fileApi.UploadFile(ctx, "static/"+id_util.NextSnowflake()+fileName, newReader, int64(size))
		if err != nil {
			return err
		}
		item := model.FileMaster{
			Model: com_model.Model{
				Id:        id_util.GetUuid(),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: gorm.DeletedAt{},
			},
			KeyName:     key,
			BucketName:  bucket,
			Url:         url,
			UserId:      userId,
			ContentType: content,
			BizType:     "",
			Size:        size,
		}
		err = a.dao.CreateFile(ctx, &item)
		if err != nil {
			log.Error(err.Error())
			return errs.NewInternal().WithError(err)
		}
		return nil
	})
	return url, bucket, key, err
}

// 内部有事务的存在
func (a *FileModule) DeleteFile(ctx context.Context, bucket, key string) (err error) {
	err = a.dao.Transaction(ctx, func(ctx context.Context) error {
		if bucket == "" {
			bucket = a.fileApi.GetBucketName(ctx)
		}
		err := a.dao.DeleteFileByBucketKey(ctx, bucket, key)
		if err != nil {
			return errs.NewInternal().WithError(err)
		}
		err = a.fileApi.DeleteFile(ctx, key)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (a *FileModule) DeleteFileById(ctx context.Context, fileId string) (err error) {
	log := a.log.Get(ctx).
		WithField(context_enum.Function.Str(), "DeleteFileById").
		WithField("file_id", fileId)
	err = a.dao.Transaction(ctx, func(ctx context.Context) error {
		file, err2 := a.dao.GetFileById(ctx, fileId)
		if err2 != nil {
			log.WithField("err", err2).Error("get file item by id err")
			return err2
		}
		if file == nil {
			log.Error("file item not found")
			return errs.NewNotExistsError().WithMsg("file not found")
		}
		err := a.dao.DeleteFileByBucketKey(ctx, file.BucketName, file.KeyName)
		if err != nil {
			log.WithField("err", err).Error("delete file item on db error")
			return errs.NewInternal().WithError(err)
		}
		err = a.fileApi.DeleteFile(ctx, file.KeyName)
		if err != nil {
			log.WithField("err", err).Error("delete file from oss error")
			return err
		}
		return nil
	})
	return err
}

func (a *FileModule) GetBaseUrl(ctx context.Context) string {
	return a.fileApi.GetBaseUrl(ctx)
}
