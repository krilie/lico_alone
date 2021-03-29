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

	container.MustProvide(ncfg.NewNConfig)
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
		dbmigrate.Migrate(innerDb, 20210206140300)
	})
}

func TestGetDbNameFromConnectStr(t *testing.T) {
	println(GetDbNameFromConnectStr("test:123456@tcp(lizo.top:3306)/1?charset=utf8mb4&parseTime=True&loc=Asia%2FShanghai&multiStatements=true"))
	println(GetDbNameFromConnectStr("test:123456@tcp(lizo.top:3306)/12?charset=utf8mb4&parseTime=True&loc=Asia%2FShanghai&multiStatements=true"))
	println(GetDbNameFromConnectStr("test:123456@tcp(lizo.top:3306)/abc?charset=utf8mb4&parseTime=True&loc=Asia%2FShanghai&multiStatements=true"))
}
