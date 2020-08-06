// +build !auto_test

package middleware

import (
	"fmt"
	"testing"
)

func TestOpsLimit(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("defer2 %v \n", err)
		}
	}()
	defer println("-0----")
	defer func() {
		defer println("defer in defer")
		if err := recover(); err != nil {
			fmt.Printf("defer %v \n", err)
			panic(err)
		}
	}()
	TPanic()
}
func TPanic() {
	panic("okkkk")
}
