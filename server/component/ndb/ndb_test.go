package ndb

import (
	"github.com/krilie/lico_alone/common/appdig"
	"github.com/krilie/lico_alone/component/ncfg"
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/stretchr/testify/assert"
	"testing"
)

var container = appdig.NewAppDig()

func TestMain(m *testing.M) {

	container.MustInvoke(ncfg.NewNConfig)
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
