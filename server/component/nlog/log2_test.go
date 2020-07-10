package nlog

import (
	"os"
	"testing"
)

func TestNewLogger(t *testing.T) {
	os.MkdirAll("./log", os.ModePerm)
	file, e := os.OpenFile("./log/log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if e != nil {
		panic(e)
		return
	}
	defer file.Close()
}
