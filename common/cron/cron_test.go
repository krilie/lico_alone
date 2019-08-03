package cron

import (
	"fmt"
	"testing"
	"time"
)

func TestCron(t *testing.T) {
	//i := 0
	spec := "*/1 * * * * *" // 每分钟
	id, e := CronGlobal.AddFunc(spec, func() {
		for j := 0; j < 3; j++ {
			fmt.Print(j)
			fmt.Print("  ")
		}
		fmt.Println()
		time.Sleep(time.Second * 10)
	})

	fmt.Println(id, e)
	time.Sleep(time.Second * 5)
	Close()
	//time.Sleep(time.Second * 120)
}
