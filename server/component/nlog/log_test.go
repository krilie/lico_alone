package nlog

import (
	context_enum "github.com/krilie/lico_alone/common/model/context-enum"
	"github.com/sirupsen/logrus"
	"testing"
)

var testEnv = context_enum.RunEnv{
	AppName:    "Test",
	AppVersion: "Test",
	AppHost:    "Test",
	ClientId:   "Test",
	UserId:     "Test",
	Version:    "Test",
	BuildTime:  "Test",
	GoVersion:  "Test",
	GitCommit:  "Test",
}

func TestLog(T *testing.T) {
	log := InitNLog(testEnv, logrus.DebugLevel)
	log.Error("hello log wire")
}
