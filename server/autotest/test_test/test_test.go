package test_test

import (
	"fmt"
	"github.com/krilie/lico_alone/common/appdig"
	"strconv"
	"testing"
)

var container = appdig.NewAppDig()

func TestMain(t *testing.M) {
	fmt.Println("自动测试dig测试main")
	container.MustProvide(func() *int {
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
	container.MustInvoke(func(i *int) {
		fmt.Println("auto test two " + strconv.Itoa(*i))
	})
}
