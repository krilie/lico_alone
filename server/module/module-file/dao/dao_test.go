package dao

import (
	"github.com/krilie/lico_alone/common/appdig"
	com_model "github.com/krilie/lico_alone/common/com-model"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/utils/id_util"
	"github.com/krilie/lico_alone/common/utils/jsonutil"
	"github.com/krilie/lico_alone/component"
	"github.com/krilie/lico_alone/module/module-file/model"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"testing"
	"time"
)

var container = appdig.
	NewAppDig().
	MustProvides(component.DigComponentProviderAllForTest).
	MustProvide(NewFileDao)

func TestFileDao_CreateFile(t *testing.T) {
	container.MustInvoke(func(dao *FileDao) {
		file := &model.FileMaster{
			Model: com_model.Model{
				Id:        id_util.GetUuid(),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: gorm.DeletedAt{
					Time:  time.Time{},
					Valid: false,
				},
			},
			KeyName:     id_util.NextSnowflake(),
			BucketName:  id_util.NextSnowflake(),
			Url:         id_util.NextSnowflake(),
			UserId:      "44",
			ContentType: "55",
			BizType:     "66",
			Size:        7,
		}
		err := dao.CreateFile(context.EmptyAppCtx(), file)
		if err != nil {
			t.Error("err db insert" + jsonutil.ToJsonPretty(file))
			t.FailNow()
		}
		err = dao.DeleteFile(context.EmptyAppCtx(), file.Id)
		if err != nil {
			t.Error("err db delete" + jsonutil.ToJsonPretty(file))
			t.FailNow()
		}
	})
}

func BenchmarkNewFileDao(b *testing.B) {
	container.MustInvoke(func(dao *FileDao) {
		dao.GetDb(context.EmptyAppCtx()).Logger.LogMode(logger.Error)
		file := &model.FileMaster{
			Model: com_model.Model{
				Id:        id_util.GetUuid(),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: gorm.DeletedAt{
					Time:  time.Time{},
					Valid: false,
				},
			},
			KeyName:     id_util.NextSnowflake(),
			BucketName:  id_util.NextSnowflake(),
			Url:         id_util.NextSnowflake(),
			UserId:      "44",
			ContentType: "55",
			BizType:     "66",
			Size:        7,
		}
		err := dao.CreateFile(context.EmptyAppCtx(), file)
		if err != nil {
			b.Error("err db insert" + jsonutil.ToJsonPretty(file))
			b.FailNow()
		}
		err = dao.DeleteFile(context.EmptyAppCtx(), file.Id)
		if err != nil {
			b.Error("err db delete" + jsonutil.ToJsonPretty(file))
			b.FailNow()
		}
	})
}
