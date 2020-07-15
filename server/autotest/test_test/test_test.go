package test_test

import (
	"fmt"
	"github.com/krilie/lico_alone/common/dig"
	"strconv"
	"testing"
)

func TestMain(t *testing.M) {
	fmt.Println("for test begin")
	dig.Container.MustProvide(func() *int {
		var i = 2
		return &i
	})
	run := t.Run()
	fmt.Printf("after run %v \n", run)
}

func TestOne(t *testing.T) {
	fmt.Println("test one")
}

func TestTwo(t *testing.T) {
	fmt.Println("test two")
	dig.Container.MustInvoke(func(i *int) {
		fmt.Println("auto test two " + strconv.Itoa(*i))
	})
}
