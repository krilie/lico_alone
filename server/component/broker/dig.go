package broker

import (
	go_smq "github.com/krilie/go-smq"
	"github.com/krilie/lico_alone/common/dig"
)

func init() {
	dig.Container.MustProvide(func() *go_smq.Smq {
		return Smq
	})
}
