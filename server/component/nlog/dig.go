package nlog

import (
	"github.com/krilie/lico_alone/common/dig"
)

func init() {
	dig.Container.MustProvide(NewLogger)
}

// DigProvider provider
func DigProvider() {
	dig.Container.MustProvide(NewLogger)
}
