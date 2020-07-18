package ndb

import (
	"github.com/krilie/lico_alone/common/dig"
	"gorm.io/gorm"
	"testing"
)

type TestDb struct {
	gorm.Model
}

func (TestDb) TableName() string {
	return "tb_test_db"
}

func TestNDb_Start(t *testing.T) {
	dig.Container.MustInvoke(func(ndb *NDb) {
		err := ndb.db.AutoMigrate(&TestDb{})
		t.Log(err)
	})
}
