package broker

import (
	"context"
	"testing"
	"time"
)

func TestFunc(t *testing.T) {
	err := Smq.Get("tt").Register(context.Background(), func(i interface{}) {
		t.Log(i)
	})
	t.Log(err)
	err = Smq.Get("tt").Send(context.Background(), "asd")
	t.Log(err)
	time.Sleep(time.Second * 2)
	Smq.Close()
}
