package dao

import (
	"github.com/krilie/lico_alone/common/ccontext"
	"github.com/krilie/lico_alone/common/cdb"
	"github.com/krilie/lico_alone/common/clog"
	"github.com/krilie/lico_alone/common/config"
	"github.com/krilie/lico_alone/module/config/model"
)

func init() {
	var log = clog.NewLog(ccontext.NewContext(), "alone.module.user.model", "init")
	err := cdb.Db.AutoMigrate(new(model.Config)).Error
	if err != nil {
		panic(err)
	}
	log.Info("user dao init done")
}

type Dao struct {
	*cdb.Dao
}

func NewDao(cfg config.DB) *Dao {
	return &Dao{
		Dao: cdb.NewDao(cfg),
	}
}
