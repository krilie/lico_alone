// +build !auto_test

package ndb

import (
	"github.com/krilie/lico_alone/common/dig"
	"github.com/krilie/lico_alone/component/ncfg"
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMain(m *testing.M) {
	ncfg.DigProvider()
	nlog.DigProvider()
	DigProvider()
	m.Run()
}

func TestNewNDb(t *testing.T) {
	dig.Container.MustInvoke(func(db *NDb) {
		err := db.Ping()
		assert.Equal(t, nil, err, "should not err")
	})
}
