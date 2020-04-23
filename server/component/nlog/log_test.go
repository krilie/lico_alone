package nlog

import (
	"testing"
)

func TestLog(T *testing.T) {
	log := InitNLog()
	log.Error("hello log wire")
}
