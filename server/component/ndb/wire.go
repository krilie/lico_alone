//+build !wireinject

package ndb

import (
	"github.com/google/wire"
	"github.com/krilie/lico_alone/component/nlog"
)

func InitNDb() (ndb *NDb, cleanUp func(), err error) {
	wire.Build(nlog.InitNLog, NewNDb)
	return &NDb{}, nil, nil
}
