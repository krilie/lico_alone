package dao

import (
	"fmt"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/dig"
	"github.com/krilie/lico_alone/common/utils/id_util"
	"github.com/krilie/lico_alone/module/module-file/model"
	"testing"
	"time"
)

func TestFileDao_CreateFile(t *testing.T) {
	dig.Container.MustInvoke(func(dao *FileDao) {
		err := dao.CreateFile(context.NewContext(), &model.FileMaster{
			Id:          id_util.NextSnowflake(),
			CreateTime:  time.Now(),
			KeyName:     "22",
			BucketName:  "33",
			UserId:      "44",
			ContentType: "55",
			BizType:     "66",
			Size:        7,
		})
		fmt.Println(err)
	})
}
