package component

import (
	"github.com/krilie/lico_alone/component/broker"
	"github.com/krilie/lico_alone/component/cache"
	"github.com/krilie/lico_alone/component/cron"
	"github.com/krilie/lico_alone/component/ncfg"
	"github.com/krilie/lico_alone/component/ndb"
	"github.com/krilie/lico_alone/component/nlog"
)

var DigComponentProviderAll = []interface{}{
	ncfg.NewNConfigByFileFromEnv("APP_CONFIG_PATH"),
	nlog.NewLogger,
	ndb.NewNDb,
	broker.NewBroker,
	cache.NewCache,
	cron.NewCrone,
}

var DigComponentProviderAllForTest = []interface{}{
	ncfg.NewNConfigByCfgStrFromEnvJson("MYAPP_TEST_CONFIG"),
	nlog.NewLogger,
	ndb.NewNDb,
	broker.NewBroker,
	cache.NewCache,
	cron.NewCrone,
}
