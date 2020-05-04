package service

import (
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/dig"
	"github.com/prometheus/common/log"
	"mime"
	"strings"
	"testing"
)

func TestFileService_UploadFile(t *testing.T) {
	dig.Container.MustInvoke(func(svc *FileService) {
		uploadStr := "hello qiniu oss"
		url, bucket, key, err := svc.UploadFile(context.NewContext(), "test", "test2.txt", strings.NewReader(uploadStr), len(uploadStr))
		t.Logf("%v %v %v %v", url, bucket, key, err)
		if err != nil {
			log.Error(err)
		}
	})
}

func TestFileService_DeleteFile(t *testing.T) {
	dig.Container.MustInvoke(func(svc *FileService) {
		err := svc.DeleteFile(context.NewContext(), "", "test")
		if err != nil {
			log.Error(err)
		}
	})
}

func TestMimeType(t *testing.T) {
	extension := mime.TypeByExtension(".stream")
	t.Logf("%v", extension)
}
