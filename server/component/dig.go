package component

import (
	"github.com/krilie/lico_alone/component/broker"
	"github.com/krilie/lico_alone/component/cache"
	"github.com/krilie/lico_alone/component/cron"
	"github.com/krilie/lico_alone/component/ncfg"
	"github.com/krilie/lico_alone/component/ndb"
	"github.com/krilie/lico_alone/component/nlog"
)

var DigComponentProviderFunc = []interface{}{
	ncfg.NewNConfig,
	nlog.NewLogger,
	ndb.NewNDb,
	broker.NewBroker,
	cache.NewCache,
	cron.NewCrone,
}

var DigComponentProviderFuncForTest = []interface{}{
	ncfg.NewNConfigByCfgStrFromEnvTest,
	nlog.NewLogger,
	ndb.NewNDb,
	broker.NewBroker,
	cache.NewCache,
	cron.NewCrone,
}
