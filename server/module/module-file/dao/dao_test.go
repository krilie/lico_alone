// +build !auto_test

package dao

import (
	com_model "github.com/krilie/lico_alone/common/com-model"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/dig"
	"github.com/krilie/lico_alone/common/utils/id_util"
	"github.com/krilie/lico_alone/common/utils/str_util"
	"github.com/krilie/lico_alone/component"
	"github.com/krilie/lico_alone/module/module-file/model"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	component.DigProviderTest()
	DigProvider()
	m.Run()
}

func TestFileDao_CreateFile(t *testing.T) {
	dig.Container.MustInvoke(func(dao *FileDao) {
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
		dao.GetDb(context.NewContext()).Logger.LogMode(logger.Error)
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
