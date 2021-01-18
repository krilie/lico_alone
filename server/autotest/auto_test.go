package autotest

import (
	"fmt"
	"strconv"
	"testing"
)

//go:generate go test -run Auto -v ./...

func TestMain(t *testing.M) {
	t.Run()
}

func TestAutoTest(t *testing.T) {
	fmt.Println("自动测试外层测试一 " + strconv.Itoa(12))
}
