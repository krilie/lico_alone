package ndb

import (
	"github.com/krilie/lico_alone/common/dig"
	"testing"
)

func TestNDb_Start(t *testing.T) {
	dig.MustInvoke(func(ndb *NDb) {
		err := ndb.db.Close()
		t.Log(err)
	})
}
