package component

import (
	"github.com/krilie/lico_alone/component/broker"
	"github.com/krilie/lico_alone/component/cache"
	"github.com/krilie/lico_alone/component/cron"
	"github.com/krilie/lico_alone/component/ncfg"
	"github.com/krilie/lico_alone/component/ndb"
	"github.com/krilie/lico_alone/component/nlog"
)

// dig provider
func DigProvider() {
	ncfg.DigProvider()
	nlog.DigProvider()
	ndb.DigProvider()
	broker.DigProvider()
	cache.DigProvider()
	cron.DigProvider()
}

// dig provider
func DigProviderTest() {
	ncfg.DigProviderByCfgStrFromEnv()
	nlog.DigProvider()
	ndb.DigProvider()
	broker.DigProvider()
	cache.DigProvider()
	cron.DigProvider()
}
