package dao

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/krilie/lico_alone/common/ccontext"
	"github.com/krilie/lico_alone/common/cdb"
	"github.com/krilie/lico_alone/common/clog"
	"github.com/krilie/lico_alone/common/config"
	"github.com/krilie/lico_alone/module/bookkeeping/model"
)

func Migrate(d *Dao) {
	var log = clog.NewLog(ccontext.NewContext(), "alone.module.account.dao", "init")
	err := d.Db.AutoMigrate(new(model.AccountItem), new(model.AccountBill), new(model.AccountBillDetail), new(model.AccountOperatorLog)).Error
	if err != nil {
		panic(err)
	}
	log.Info("account dao init done")
}

type Dao struct {
	*cdb.Dao
}

func NewDao(cfg config.DB) *Dao {
	d := &Dao{
		Dao: cdb.NewDao(cfg),
	}
	Migrate(d)
	return d
}
