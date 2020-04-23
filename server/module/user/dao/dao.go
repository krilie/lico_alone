package dao

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/krilie/lico_alone/common/cdb"
	"github.com/krilie/lico_alone/common/config"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/component/clog"
	"github.com/krilie/lico_alone/module/user/model"
)

func AutoMigrate(d *Dao) {
	var log = clog.NewLog(context.NewContext(), "alone.module.user.model", "init")
	err := d.Db.AutoMigrate(new(model.Permission), new(model.RolePermission), new(model.Role), new(model.UserRole), new(model.UserMaster)).Error
	if err != nil {
		panic(err)
	}
	log.Info("user dao init done")
}

type Dao struct {
	*cdb.Dao
}

func NewDao(cfg config.DB) *Dao {
	d := &Dao{
		Dao: cdb.NewDao(cfg),
	}
	AutoMigrate(d)
	return d
}
