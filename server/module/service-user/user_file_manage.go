package service_user

import (
	"context"
	"io"
)

func (a *UserService) UploadFile(ctx context.Context, userId, fileName string, file io.ReadSeeker, size int) (url, bucket, key string, err error) {
	return a.moduleFile.UploadFile(ctx, userId, fileName, file, size)
}

func (a *UserService) DeleteFile(ctx context.Context, key string) (err error) {
	return a.moduleFile.DeleteFile(ctx, "", key)
}
