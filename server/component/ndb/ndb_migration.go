package ndb

import (
	"context"
	"github.com/krilie/lico_alone/component/dbmigrate"
)

func (ndb *NDb) Migration(ctx context.Context, path string, version uint) {
	innerDb, err := ndb.db.DB()
	if err != nil {
		panic(err)
	}
	dbmigrate.Migrate(innerDb, path, version)
}
