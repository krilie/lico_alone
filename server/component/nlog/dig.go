package nlog

import (
	"github.com/krilie/lico_alone/common/dig"
)

var GLog *NLog

// DigProvider provider
func DigProvider() {
	dig.Container.MustProvide(NewLogger)
}
