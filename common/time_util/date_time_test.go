package time_util

import (
	"fmt"
	"testing"
	time2 "time"
)

func TestGetBeijingZeroTime(t *testing.T) {
	fmt.Println(time2.Now())
	time := time2.Now()
	zeroTime := GetBeijingZeroTime(time)
	fmt.Println(zeroTime)
	fmt.Println(GetBeijingMonthStartTime(time2.Now()))
	fmt.Println(GetBeijingLastDateOfMonth(time2.Now()))
}
