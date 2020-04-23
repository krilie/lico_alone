//+build wireinject

package nlog

import (
	"github.com/google/wire"
	"github.com/krilie/lico_alone/common/config"
	const_val "github.com/krilie/lico_alone/common/const-val"
)

var NLogProviderSet = wire.NewSet(NewLogger)

func InitNLog() *NLog {
	wire.Build(NLogProviderSet, wire.Value(const_val.RunEnv), wire.Value(config.Cfg))
	return &NLog{}
}
