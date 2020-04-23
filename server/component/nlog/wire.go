//+build wireinject

package nlog

import (
	"github.com/google/wire"
	context_enum "github.com/krilie/lico_alone/common/model/context-enum"
	"github.com/sirupsen/logrus"
)

var NLogProviderSet = wire.NewSet(NewLogger)

func InitNLog(runEnv context_enum.RunEnv, level logrus.Level) *NLog {
	wire.Build(NLogProviderSet)
	return &NLog{}
}
