package broker

import (
	"context"
	"github.com/krilie/lico_alone/component/broker/messages"
	"testing"
	"time"
)

func TestFunc(t *testing.T) {
	err := Smq.Register(context.Background(), func(i *messages.TestMessage) {
		t.Log("ffffffffff" + i.Test)
	})
	t.Log(err)
	err = Smq.Send(context.Background(), &messages.TestMessage{Test: "for test test"})
	t.Log(err)
	time.Sleep(time.Second * 2)
	Smq.Close()
}
