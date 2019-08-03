package broker

import (
	"fmt"
	"testing"
	"time"
)

func TestNewStartedBroker(t *testing.T) {
	brokerTest := NewStartedBroker("test", 3) // 0 相当于同步调用
	brokerTest.Register(func(i interface{}) {
		fmt.Println(i)
		j := 0
		m := 6 / j
		fmt.Print(m)
	})
	brokerTest.Register(func(i interface{}) {
		fmt.Println("seconed", i)
	})
	brokerTest.Register(func(i interface{}) {
		fmt.Println("seconedseconedseconedseconedseconed", i)
		time.Sleep(time.Second * 2)
	})
	brokerTest.Send("assssss")
	brokerTest.Stop()
	err := brokerTest.Send("123")
	if err != nil {
		t.Log(err)
	}
}
