package file_api

import (
	"context"
	"fmt"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/common/utils/fileutil"
	"github.com/krilie/lico_alone/component/ncfg"
	"github.com/minio/minio-go"
	"io"
	"strings"
	"time"
)

type OssMinio struct {
	Client     *minio.Client
	BucketName string
	Url        string
}

func (f *OssMinio) DeleteFileByUrl(ctx context.Context, url string) error {
	bucket, key := f.GetBucketAndKeyByUrl(ctx, url)
	if bucket != f.BucketName {
		return errs.NewParamError().WithMsg("minio err bucket name on delete file")
	}
	if err := f.DeleteFile(ctx, key); err != nil {
		return err
	}
	return nil
}

func (f *OssMinio) GetUrl(ctx context.Context, isPub bool, fileKey string) (url string, err error) {
	if isPub {
		return fmt.Sprintf("%v/%v/%v", f.Url, f.BucketName, fileKey), nil
	} else {
		url, err := f.Client.PresignedGetObject(f.BucketName, fileKey, time.Hour*5, nil)
		if err != nil {
			return "", err
		}
		return url.String(), nil
	}
}

func (f *OssMinio) GetBucketName(ctx context.Context) string {
	return f.BucketName
}

func (f *OssMinio) UploadFile(ctx context.Context, fileName string, fileStream io.Reader, fileSize int64) (url, key string, err error) {
	content, reader, err := fileutil.GetContentType2(fileStream)
	if err != nil {
		return "", "", err
	}
	key = fileName
	userMate := make(map[string]string)
	userMate["user_id"] = "userId"
	n, err := f.Client.PutObject(f.BucketName, key, reader, fileSize, minio.PutObjectOptions{ContentType: content, UserMetadata: userMate})
	if err != nil {
		_ = f.Client.RemoveIncompleteUpload(f.BucketName, key) // 删除可能存在的不完整文件
		return f.BucketName, key, errs.NewInternal().WithError(err)
	} else if fileSize != -1 && n != fileSize && fileSize > 1024*1024*64 {
		_ = f.Client.RemoveIncompleteUpload(f.BucketName, key) // 删除可能存在的不完整文件
		return f.BucketName, key, errs.NewInternal().WithMsg("un completed upload please check")
	} else {
		return f.GetBaseUrl(ctx) + "/" + key, key, nil
	}
}

func (f *OssMinio) DeleteFile(ctx context.Context, key string) error {
	err := f.Client.RemoveObject(f.BucketName, key)
	if err != nil {
		return errs.NewInternal().WithError(err)
	}
	return nil
}

func (f *OssMinio) GetBaseUrl(ctx context.Context) string {
	if !strings.HasPrefix(f.Url, "http") {
		return "https://" + f.Url
	}
	return f.Url
}

func (f *OssMinio) GetBucketAndKeyByUrl(ctx context.Context, url string) (bucket, key string) {
	bucketKey := strings.Replace(url, f.Url+"/", "", 1)
	bucketKeySlice := strings.SplitN(bucketKey, "/", 1)
	return bucketKeySlice[0], bucketKeySlice[1]
}

func NewOssMinioClientByCfg(cfg *ncfg.FileSave) *OssMinio {
	return NewOssMinioClient(cfg.OssBucket, cfg.OssEndPoint, cfg.OssKey, cfg.OssSecret)
}

func NewOssMinioClient(bucket, endPoint, key, secret string) *OssMinio {
	if strings.HasPrefix(endPoint, "http://") {
		minioClient, err := minio.New(strings.TrimPrefix(endPoint, "http://"), key, secret, false) //endpoint, accessKeyID, secretAccessKey string, secure bool
		if err != nil {
			panic(errs.NewInternal().WithError(err))
		}
		url := fmt.Sprintf("%v%v%v", endPoint, "/", bucket)
		return &OssMinio{Client: minioClient, BucketName: bucket, Url: url}
	} else if strings.HasPrefix(endPoint, "https://") {
		minioClient, err := minio.New(strings.TrimPrefix(endPoint, "https://"), key, secret, true) //endpoint, accessKeyID, secretAccessKey string, secure bool
		if err != nil {
			panic(errs.NewInternal().WithError(err))
		}
		url := fmt.Sprintf("%v%v%v", endPoint, "/", bucket)
		return &OssMinio{Client: minioClient, BucketName: bucket, Url: url}
	} else {
		minioClient, err := minio.New(endPoint, key, secret, true) //endpoint, accessKeyID, secretAccessKey string, secure bool
		if err != nil {
			panic(errs.NewInternal().WithError(err))
		}
		url := fmt.Sprintf("%v%v%v", endPoint, "/", bucket)
		return &OssMinio{Client: minioClient, BucketName: bucket, Url: url}
	}
}
