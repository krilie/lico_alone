package log

import "github.com/sirupsen/logrus"

func init() {
	Log.SetFormatter(&logrus.TextFormatter{})
	Log.SetLevel(logrus.DebugLevel)
	Log.Infoln("log ok")
}

var Log = logrus.New()
