package dao

import (
	"github.com/krilie/lico_alone/common/dig"
)

func init() {
	dig.Container.MustProvide(NewMessageDao)
}

// DigProvider provider
func DigProvider() {
	dig.Container.MustProvide(NewMessageDao)
}
