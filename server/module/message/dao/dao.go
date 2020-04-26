package dao

import (
	"github.com/krilie/lico_alone/common/config"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/krilie/lico_alone/module/message/model"
)

func AutoMigrate(d *Dao) {
	var log = nlog.NewLog(context.NewContext(), "alone.module.user.model", "init")
	err := d.Db.AutoMigrate(new(model.MessageEmail), new(model.MessageSms), new(model.MessageValidCode)).Error
	if err != nil {
		panic(err)
	}
	log.Info("user dao init done")
}

func NewDao(cfg config.DB) *Dao {
	d := &Dao{
		Dao: cdb.NewDao(cfg),
	}
	AutoMigrate(d)
	return d
}

type Dao struct {
	*cdb.Dao
}
