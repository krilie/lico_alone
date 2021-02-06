package ndb

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/krilie/lico_alone/common/appdig"
	"github.com/krilie/lico_alone/component/dbmigrate"
	"github.com/krilie/lico_alone/component/ncfg"
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/stretchr/testify/assert"
	"testing"
)

var container = appdig.NewAppDig()

func TestMain(m *testing.M) {

	container.MustProvide(ncfg.NewNConfigByCfgStrFromEnvJson("MYAPP_TEST_CONFIG"))
	container.MustProvide(nlog.NewLogger)
	container.MustProvide(NewNDb)

	m.Run()
}

func TestNewNDb(t *testing.T) {
	container.MustInvoke(func(db *NDb) {
		err := db.Ping()
		assert.Equal(t, nil, err, "should not err")
	})
}

func TestMigrate(t *testing.T) {
	container.MustInvoke(func(db *NDb, cfg *ncfg.NConfig) {
		err := db.Ping()
		assert.Equal(t, nil, err, "should not err")
		innerDb, _ := db.db.DB()
		dbmigrate.Migrate(innerDb, "test", "file://c://sqls", 20210206140300)
	})
}
