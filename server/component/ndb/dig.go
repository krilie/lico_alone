package ndb

import (
	"github.com/krilie/lico_alone/common/dig"
)

func init() {
	dig.Container.MustProvide(NewNDb)
}
