package dao

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/krilie/lico_alone/common/config"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/krilie/lico_alone/module/bookkeeping/model"
)

func Migrate(d *Dao) {
	var log = nlog.NewLog(context.NewContext(), "alone.module.account.dao", "init")
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
