package comlog

import (
	c2 "github.com/krilie/lico_alone/common/context"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestLog(T *testing.T) {
	Log.Info("okokok")
	log := Log.WithFields(logrus.Fields{"ok": "23"})
	log.Info("ok", "ok", "okkk")
}

func TestLogNewLog(T *testing.T) {
	log := NewLog(c2.NewContext(), "ok", "")
	log.Info("info info")
	log.WithField("function", "testlognewlog").Info("ok info info 2")
}
