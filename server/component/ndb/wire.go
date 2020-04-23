//+build !wireinject

package ndb

import (
	"github.com/google/wire"
	"github.com/krilie/lico_alone/component/nlog"
)

func InitNDb(name string) (ndb *NDb, cleanUp func(), err error) {
	wire.Build(nlog.NLogProviderSet)
	return &NDb{}, nil, nil
}
