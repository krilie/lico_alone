package ndb

import (
	"github.com/krilie/lico_alone/common/config"
	"github.com/krilie/lico_alone/common/dig"
	"github.com/krilie/lico_alone/component/nlog"
)

func init() {
	dig.MustProvide(func(cfg *config.Config, log *nlog.NLog) (ndb *NDb) {
		return NewNDb(cfg.DB, log)
	})
}
