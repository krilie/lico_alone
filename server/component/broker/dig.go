package broker

import (
	"context"
	go_smq "github.com/krilie/go-smq"
	"github.com/krilie/lico_alone/common/dig"
)

type Broker struct {
	*go_smq.Smq
}

func NewBroker() *Broker {
	return &Broker{Smq: go_smq.NewSmq()}
}

func (broker *Broker) MustSend(ctx context.Context, msg go_smq.Message) {
	err := broker.Send(ctx, msg)
	if err != nil {
		panic(err)
	}
}
func (broker *Broker) MustRegister(ctx context.Context, f interface{}) {
	err := broker.Register(ctx, f)
	if err != nil {
		panic(err)
	}
}

func init() {
	dig.Container.MustProvide(NewBroker)
}

// dig provider
func DigProvider() {
	dig.Container.MustProvide(NewBroker)
}
