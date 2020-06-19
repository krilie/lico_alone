package cron

import (
	"fmt"
	"github.com/krilie/lico_alone/common/context"
	"testing"
	"time"
)

func TestCron(t *testing.T) {
	c := NewCrone()
	_, _ = c.AddFunc("*/1 * * * * *", func() {
		fmt.Println("ok")
	})
	_, _ = c.AddFunc("@every 2s", func() {
		fmt.Println("ok 2")
	})
	time.Sleep(time.Minute)
	c.StopAndWait(context.NewContext())
}
