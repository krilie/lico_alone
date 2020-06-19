package broker

import (
	go_smq "github.com/krilie/go-smq"
	"github.com/krilie/lico_alone/common/dig"
)

type Broker struct {
	*go_smq.Smq
}

func NewBroker() *Broker {
	return &Broker{Smq: go_smq.NewSmq()}
}

func init() {
	dig.Container.MustProvide(NewBroker)
}
