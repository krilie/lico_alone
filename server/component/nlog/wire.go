//+build wireinject

package nlog

import (
	"github.com/google/wire"
	const_val "github.com/krilie/lico_alone/common/const-val"
)

var NLogProviderSet = wire.NewSet(NewLogger)

func InitNLog() *NLog {
	wire.Build(NLogProviderSet, const_val.WireProviderRunEnvAndConfig)
	return &NLog{}
}
