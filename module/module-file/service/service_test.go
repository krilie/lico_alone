package service

import (
	"github.com/krilie/lico_alone/common/appdig"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/utils/id_util"
	"github.com/krilie/lico_alone/component"
	"github.com/prometheus/common/log"
	"mime"
	"strings"
	"testing"
)

var container = appdig.
	NewAppDig().
	MustProvides(component.DigComponentProviderAll).
	MustProvides(DigModuleFileProviderAll)

func TestFileService_UploadFile(t *testing.T) {
	container.MustInvoke(func(svc *FileModule) {
		uploadStr := "hello qiniu oss"
		url, bucket, key, err := svc.UploadFile(context.EmptyAppCtx(), "test", "tts/"+id_util.GetUuid()+"test2.txt", strings.NewReader(uploadStr), len(uploadStr))
		t.Logf("%v %v %v %v", url, bucket, key, err)
		if err != nil {
			log.Error(err)
		}
	})
}

func TestFileService_DeleteFile(t *testing.T) {
	container.MustInvoke(func(svc *FileModule) {
		err := svc.DeleteFile(context.EmptyAppCtx(), "", "test")
		if err != nil {
			log.Error(err)
		}
	})
}

func TestMimeType(t *testing.T) {
	extension := mime.TypeByExtension(".stream")
	t.Logf("%v", extension)
}
