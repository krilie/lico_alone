package service

import (
	"github.com/krilie/lico_alone/common/com-model"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/utils/str_util"
	"github.com/krilie/lico_alone/module/module-file/model"
	"testing"
)

func TestFileModule_QueryFilePage(t *testing.T) {
	container.MustInvoke(func(svc *FileModule) {
		page, count, num, size, files, err := svc.QueryFilePage(context.EmptyAppCtx(), model.QueryFileParam{
			PageParams:     com_model.PageParams{},
			KeyNameLike:    "1",
			BucketNameLike: "2",
			UrlLike:        "3",
			UserId:         "4",
			BizType:        "5",
			ContentType:    "6",
			CreatedAtBegin: nil,
			CreatedAtEnd:   nil,
		})
		t.Log(page, count, num, size, str_util.ToJsonPretty(files), err)
	})
}
