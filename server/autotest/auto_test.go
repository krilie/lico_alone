package autotest

import (
	"fmt"
	"github.com/krilie/lico_alone/common/dig"
	"strconv"
	"testing"
)

//go:generate go test -tags "auto_test" -v ./...

func TestMain(t *testing.M) {
	fmt.Println("自动测试外层TestMain")
	dig.Container.MustProvide(func() *int {
		var i = 1
		return &i
	})
	t.Run()
}

func TestAutoTest(t *testing.T) {
	t.Log("自动测试外层测试")
	dig.Container.MustInvoke(func(i *int) {
		fmt.Println("自动测试外层测试一 " + strconv.Itoa(*i))
	})
}
