// +build !auto_test

package fileutils

import (
	"context"
	"github.com/issue9/watermark"
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

func TestWaterMarkOne2(t *testing.T) {
	w, err := watermark.New("./beijing8.gif", 2, watermark.TopLeft)
	if err != nil {
		panic(err)
	}
	create, err := os.Create("./res.jpg")
	if err != nil {
		panic(err)
	}
	defer create.Close()
	err = w.Mark(create, "C:\\Users\\Administrator\\Desktop\\qiang.jpg")
	println(err.Error())
}
