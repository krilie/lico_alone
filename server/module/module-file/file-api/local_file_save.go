package file_api

import (
	"context"
	"fmt"
	"github.com/krilie/s-file"
	"io"
)

type LocalFileSave struct {
	File *s_file.SFile
	Url  string
}

func (o *LocalFileSave) GetBucketName(ctx context.Context) string {
	return ""
}

func (o *LocalFileSave) UploadFile(ctx context.Context, fileName string, fileStream io.ReadSeeker, fileSize int64) (url, key string, err error) {
	_, key, err = o.File.SaveFile(ctx, fileName, fileStream)
	if err != nil {
		return "", "", err
	}
	url, err = o.GetUrl(ctx, true, key)
	if err != nil {
		return "", "", err
	}
	return url, key, nil
}

func (o *LocalFileSave) DeleteFile(ctx context.Context, fileKey string) error {
	return o.File.DeleteFile(ctx, fileKey)

}

func (o *LocalFileSave) GetUrl(ctx context.Context, isPub bool, fileKey string) (url string, err error) {
	return fmt.Sprintf("%v/%v", o.GetBaseUrl(ctx), fileKey), nil
}

func (o *LocalFileSave) GetBaseUrl(ctx context.Context) string {
	return o.Url
}

func NewLocalFileSave(saveDir string, baseUrl string) *LocalFileSave {
	return &LocalFileSave{
		File: s_file.NewSFile(saveDir),
		Url:  baseUrl,
	}
}
