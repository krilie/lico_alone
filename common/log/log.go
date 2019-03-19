package log

import "github.com/sirupsen/logrus"

func init() {
	Log.SetFormatter(&logrus.TextFormatter{})
	Log.SetLevel(logrus.DebugLevel)
	Log.Infoln("log ok")
}

var Log = logrus.New()

func Debug(v ...interface{}) {
	Log.Debug(v)
}
func Debugf(format string, v ...interface{}) {
	Log.Debugf(format, v)
}
func Debugln(v ...interface{}) {
	Log.Debugln(v)
}

func Error(v ...interface{}) {
	Log.Error(v)
}
func Errorf(format string, v ...interface{}) {
	Log.Errorf(format, v)
}
func Errorln(v ...interface{}) {
	Log.Errorln(v)
}
func Info(v ...interface{}) {
	Log.Info(v)
}
func Infof(format string, v ...interface{}) {
	Log.Infof(format, v)
}
func Infoln(v ...interface{}) {
	Log.Infoln(v)
}

func Warning(v ...interface{}) {
	Log.Warning(v)
}
func Warningf(format string, v ...interface{}) {
	Log.Warningf(format, v)
}
func Warningln(v ...interface{}) {
	Log.Warningln(v)
}
func Fatal(v ...interface{}) {
	Log.Fatal(v)
}
func Fatalf(format string, v ...interface{}) {
	Log.Fatalf(format, v)
}
func Fatalln(v ...interface{}) {
	Log.Fatalln(v)
}
func Panic(v ...interface{}) {
	Log.Panic(v)
}
func Panicf(format string, v ...interface{}) {
	Log.Panicf(format, v)
}
func Panicln(v ...interface{}) {
	Log.Panicln(v)
}
