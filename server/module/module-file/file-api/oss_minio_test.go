package file_api

import (
	"context"
	"strings"
	"testing"
)

var api FileOperator

func TestMain(m *testing.M) {
	api = NewOssMinioClient(
		"demo",
		"sample.com",
		"admin",
		"admin123",
	)
	m.Run()
}

func TestOssMinio_GetBucketAndKeyByUrl(t *testing.T) {
	url, key, err := api.UploadFile(context.Background(), "111", strings.NewReader("123"), -1)
	t.Log(url)
	t.Log(key)
	t.Log(err)
}

func TestOssMinio_DeleteFile(t *testing.T) {
	err := api.DeleteFile(context.Background(), "cd652cbb-bf50-42cb-b13a-54a6ac08815b111")
	t.Log(err)
}
