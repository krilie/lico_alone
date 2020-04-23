//+build wireinject

package ndb

import (
	"github.com/google/wire"
	"github.com/krilie/lico_alone/common/config"
	const_val "github.com/krilie/lico_alone/common/const-val"
	"github.com/krilie/lico_alone/component/nlog"
)

func InitNDb() (ndb *NDb, cleanUp func(), err error) {
	wire.Build(NewNDb, nlog.InitNLog, wire.FieldsOf(&config.Config{}, "DB"), const_val.WireProviderRunEnvAndConfig)
	return &NDb{}, nil, nil
}
