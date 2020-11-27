package cron

import (
	"fmt"
	"github.com/krilie/lico_alone/common/context"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	DigProvider()
	m.Run()
}

func TestCron(t *testing.T) {
	c := NewCrone()
	_, _ = c.AddFunc("*/1 * * * * *", func() {
		fmt.Println("ok")
	})
	_, _ = c.AddFunc("@every 2s", func() {
		fmt.Println("ok 2")
	})
	time.Sleep(time.Second * 10)
	c.StopAndWait(context.NewContext())
}
