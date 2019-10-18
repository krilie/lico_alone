package dao

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/krilie/lico_alone/common/ccontext"
	"github.com/krilie/lico_alone/common/cdb"
	"github.com/krilie/lico_alone/common/clog"
	"github.com/krilie/lico_alone/common/config"
	"github.com/krilie/lico_alone/module/user/model"
)

func Migrate(d *Dao) {
	var log = clog.NewLog(ccontext.NewContext(), "alone.module.account.dao", "init")
	err := d.Db.AutoMigrate(new(model.Permission), new(model.RolePermission), new(model.Role), new(model.UserRole), new(model.UserMaster)).Error
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
