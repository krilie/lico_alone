package ccron

import (
	"fmt"
	"testing"
	"time"
)

func TestCron(t *testing.T) {
	c := NewCrone()
	_, _ = c.AddFunc("*/1 * * * * *", func() {
		fmt.Println("ok")
	})
	time.Sleep(time.Minute)
	Stop(c)
}
