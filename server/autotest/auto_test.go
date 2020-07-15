package autotest

import (
	"fmt"
	"github.com/krilie/lico_alone/common/dig"
	"strconv"
	"testing"
)

//go:generate go test -tags "auto_test" -v ./...

func TestMain(t *testing.M) {
	fmt.Println("begin auto test")
	dig.Container.MustProvide(func() *int {
		var i = 1
		return &i
	})
	t.Run()
}

func TestTest(t *testing.T) {
	t.Log("auto test one")
	dig.Container.MustInvoke(func(i *int) {
		fmt.Println("auto test one " + strconv.Itoa(*i))
	})
}
