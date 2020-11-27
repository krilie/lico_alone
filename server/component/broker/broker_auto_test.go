package broker

import (
	"context"
	"github.com/krilie/lico_alone/common/dig"
	"github.com/krilie/lico_alone/component/broker/messages"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMain(m *testing.M) {
	DigProvider()
	m.Run()
}

func TestAutoFunc(t *testing.T) {
	count := 0
	dig.Container.MustInvoke(func(Smq *Broker) {
		// register smq
		err := Smq.Register(context.Background(), func(i *messages.TestMessage) {
			count++
			assert.Equal(t, "for test test", i.Test, "should equal")
		})
		assert.Equal(t, nil, err, "should nil")
		// send message
		for i := 0; i < 777; i++ {
			err = Smq.Send(context.Background(), &messages.TestMessage{Test: "for test test"})
			assert.Equal(t, nil, err, "should nil")
			if err != nil {
				t.FailNow()
			}
		}
		// check send message
		Smq.Close()
		assert.Equal(t, 777, count, "no match count")
	})
}
