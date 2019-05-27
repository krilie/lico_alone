package service

import (
	"github.com/krilie/lico_alone/common/comstruct/errs"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/uuid_util"
	"github.com/krilie/lico_alone/module/file/model"
	"github.com/minio/minio-go"
	"io"
	"mime/multipart"
	"net/http"
	"time"
)

func (FileOp) UploadFile(ctx *context.Context, userId, fileName string, file multipart.File, size int64) (string, error) {
	minioClient, err := minio.New(ossEndPoint, ossKey, ossSecret, true)
	if err != nil {
		return "", errs.ErrInternal.NewWithMsg(err.Error())
	}
	decByte := make([]byte, 512)
	if _, e := file.Read(decByte); e != nil {
		return "", errs.ErrInternal.NewWithMsg(e.Error())
	}
	if _, e := file.Seek(0, io.SeekStart); e != nil {
		return "", errs.ErrInternal.NewWithMsg(e.Error())
	}
	contentType := http.DetectContentType(decByte)
	objName := uuid_util.GetUuid() + fileName
	userMate := make(map[string]string)
	userMate["user_id"] = userId
	tx := model.Db.Begin()
	fileS := model.File{}
	fileS.ID = uuid_util.GetUuid()
	fileS.CreateTime = time.Now()
	fileS.ObjKey = objName
	fileS.UserId = userId
	fileS.ContentType = contentType
	//fileS.BizType = ""
	fileS.Size = int(size)
	err = tx.Create(fileS).Error
	if err != nil {
		tx.Rollback()
		return "", errs.ErrInternal.NewWithMsg(err.Error())
	}
	n, err := minioClient.PutObject(ossBucket, objName, file, size, minio.PutObjectOptions{ContentType: contentType, UserMetadata: userMate})
	if err != nil {
		tx.Rollback()
		_ = minioClient.RemoveIncompleteUpload(ossBucket, objName) // 删除可能存在的不完整文件
		return "", errs.ErrInternal.NewWithMsg(err.Error())
	} else if n != size {
		tx.Rollback()
		_ = minioClient.RemoveIncompleteUpload(ossBucket, objName) // 删除可能存在的不完整文件
		return "", errs.ErrInternal.NewWithMsg("un completed upload please check")
	} else {
		tx.Commit()
		return objName, nil
	}
}

func (FileOp) DeleteFile(ctx *context.Context, userId, filePath string) error {
	tx := model.Db.Begin()
	e := tx.Delete(model.File{}, model.File{ObjKey: filePath}).Error
	if e != nil {
		tx.Rollback()
		return errs.ErrInternal.NewWithMsg(e.Error())
	}
	minioClient, err := minio.New(ossEndPoint, ossKey, ossSecret, true)
	if err != nil {
		tx.Rollback()
		return errs.ErrInternal.NewWithMsg(err.Error())
	}
	err = minioClient.RemoveObject(ossBucket, filePath)
	if err != nil {
		tx.Rollback()
		return errs.ErrInternal.NewWithMsg(err.Error())
	}
	tx.Commit()
	return nil
}
