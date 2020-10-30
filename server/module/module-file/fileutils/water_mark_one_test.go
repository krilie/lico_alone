// +build !auto_test

package fileutils

import (
	"context"
	"os"
	"testing"
)

func TestWaterMarkOne(t *testing.T) {
	open, err := os.Open("C:\\Users\\Administrator\\Desktop\\qiang.jpg")
	if err != nil {
		panic(err)
	}
	defer open.Close()
	create, err := os.Create("./res.jpg")
	if err != nil {
		panic(err)
	}
	defer create.Close()
	err = WaterMarkOne(context.Background(), open, create)
	if err != nil {
		panic(err)
	}
}
