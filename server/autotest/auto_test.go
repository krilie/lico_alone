package autotest

import (
	"fmt"
	"testing"
)

//go:generate go test -v ./...

func TestMain(t *testing.M) {
	fmt.Println("begin auto test")
	t.Run()
}

func TestTest(t *testing.T) {
	t.Log("auto test one")
}
