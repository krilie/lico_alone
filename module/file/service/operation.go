package service

import (
	"github.com/krilie/lico_alone/common/common_struct/errs"
	"github.com/krilie/lico_alone/common/context_util"
	"github.com/krilie/lico_alone/common/uuid_util"
	"github.com/minio/minio-go"
	"io"
	"mime/multipart"
	"net/http"
)

func UploadFile(ctx *context_util.Context, userId, fileName string, file multipart.File, size int64) (string, *errs.Error) {
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
	n, err := minioClient.PutObject(ossBucket, objName, file, size, minio.PutObjectOptions{ContentType: contentType, UserMetadata: userMate})
	if err != nil {
		_ = minioClient.RemoveIncompleteUpload(ossBucket, objName) // 删除可能存在的不完整文件
		return "", errs.ErrInternal.NewWithMsg(err.Error())
	} else if n != size {
		_ = minioClient.RemoveIncompleteUpload(ossBucket, objName) // 删除可能存在的不完整文件
		return "", errs.ErrInternal.NewWithMsg("un completed upload please check")
	} else {
		return objName, nil
	}
}

func DeleteFile(ctx *context_util.Context, userId, filePath string) *errs.Error {
	minioClient, err := minio.New(ossEndPoint, ossKey, ossSecret, true)
	if err != nil {
		return errs.ErrInternal.NewWithMsg(err.Error())
	}
	err = minioClient.RemoveObject(ossBucket, filePath)
	if err != nil {
		return errs.ErrInternal.NewWithMsg(err.Error())
	}
	return nil
}
