package dao

import (
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/dig"
	"github.com/krilie/lico_alone/common/utils/id_util"
	"github.com/krilie/lico_alone/common/utils/str_util"
	"github.com/krilie/lico_alone/module/module-file/model"
	"testing"
	"time"
)

func TestFileDao_CreateFile(t *testing.T) {
	dig.Container.MustInvoke(func(dao *FileDao) {
		file := &model.FileMaster{Id: id_util.NextSnowflake(), CreateTime: time.Now(), KeyName: id_util.NextSnowflake(), BucketName: id_util.NextSnowflake(), Url: id_util.NextSnowflake(), UserId: "44", ContentType: "55", BizType: "66", Size: 7}
		err := dao.CreateFile(context.NewContext(), file)
		if err != nil {
			t.Error("err db insert" + str_util.ToJsonPretty(file))
			t.FailNow()
		}
		err = dao.DeleteFile(context.NewContext(), file.Id)
		if err != nil {
			t.Error("err db delete" + str_util.ToJsonPretty(file))
			t.FailNow()
		}
	})
}

func BenchmarkNewFileDao(b *testing.B) {
	dig.Container.MustInvoke(func(dao *FileDao) {
		dao.GetDb(context.NewContext()).LogMode(false)
		file := &model.FileMaster{Id: id_util.NextSnowflake(), CreateTime: time.Now(), KeyName: id_util.NextSnowflake(), BucketName: id_util.NextSnowflake(), Url: id_util.NextSnowflake(), UserId: "44", ContentType: "55", BizType: "66", Size: 7}
		err := dao.CreateFile(context.NewContext(), file)
		if err != nil {
			b.Error("err db insert" + str_util.ToJsonPretty(file))
			b.FailNow()
		}
		err = dao.DeleteFile(context.NewContext(), file.Id)
		if err != nil {
			b.Error("err db delete" + str_util.ToJsonPretty(file))
			b.FailNow()
		}
	})
}
