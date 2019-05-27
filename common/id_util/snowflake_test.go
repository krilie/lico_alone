package id_util

import (
	"fmt"
	"testing"
)

func TestNextSnowflakeId(t *testing.T) {
	for i := 0; i < 100; i++ {
		fmt.Println(NextSnowflakeId().String())
	}
}
