package test_test

import (
	"fmt"
	"github.com/krilie/lico_alone/common/dig"
	"strconv"
	"testing"
)

func TestMain(t *testing.M) {
	fmt.Println("自动测试dig测试main")
	dig.Container.MustProvide(func() *int {
		var i = 2
		return &i
	})
	run := t.Run()
	fmt.Printf("after run %v \n", run)
}

func TestAutoOne(t *testing.T) {
	fmt.Println("test one")
}

func TestAutoTwo(t *testing.T) {
	fmt.Println("test two")
	dig.Container.MustInvoke(func(i *int) {
		fmt.Println("auto test two " + strconv.Itoa(*i))
	})
}
