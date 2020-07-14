package test_test

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.M) {
	fmt.Println("for test begin")
	run := t.Run()
	fmt.Printf("after run %v \n", run)
}

func TestOne(t *testing.T) {
	fmt.Println("test one")
}

func TestTwo(t *testing.T) {
	fmt.Println("test two")
}
