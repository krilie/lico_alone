package clog

import (
	"context"
	context2 "github.com/krilie/lico_alone/common/context"
	"github.com/sirupsen/logrus"
	"os"
)

const (
	AppName    = "app_name"
	AppVersion = "app_version"
	AppHost    = "app_host"
	TraceId    = "trace_id"
	ClientId   = "client_id"
	UserId     = "user_id"
	Module     = "module"
	Function   = "function"
	Stack      = "stack"
)

var Log = logrus.NewEntry(logrus.New())

func init() {
	//Log.Logger.SetFormatter(&logrus.JSONFormatter{TimestampFormat: "2006-01-02T15:04:05.000Z07:00"})
	Log.Logger.SetFormatter(&logrus.TextFormatter{})
	Log.Logger.SetLevel(logrus.DebugLevel)
	Log.Logger.SetOutput(os.Stdout)
	Log = Log.
		WithField(AppName, os.Getenv("BR_APP_NAME")).
		WithField(AppVersion, os.Getenv("BR_APP_VERSION")).
		WithField(AppHost, func() string {
			if name, err := os.Hostname(); err != nil {
				return name
			} else {
				return ""
			}
		})
	Log.Infoln("log init ok")
}

func SetLogField(appVersion, appName string) {
	Log = logrus.WithField(AppName, appName).WithField(AppVersion, appVersion)
}

func SetUpLogFile(f string) {
	if f == "" || f == "stdout" {
		Log.Logger.SetOutput(os.Stdout)
		Log.Logger.Warnln("set log out file to stdout")
		return
	}
	file, e := os.OpenFile(f, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if e != nil {
		panic(e)
		return
	}
	Log.Logger.SetOutput(file)
	Log.Logger.Warnln("set log out file to " + f)
}

// trace_id
func NewLog(ctx context.Context, moduleName string, functionName string) *logrus.Entry {
	bctx := context2.GetContextOrNew(ctx)
	return Log.WithFields(logrus.Fields{
		TraceId:  bctx.GetTraceId(),
		ClientId: bctx.GetClientId(),
		UserId:   bctx.GetUserId(),
		Module:   moduleName,
		Function: functionName})
}

func With(ctx context.Context, location ...string) *logrus.Entry {
	var module, funcName string
	if len(location) > 0 {
		module = location[0]
	}
	if len(location) > 1 {
		funcName = location[1]
	}
	c := context2.GetContextOrNew(ctx)
	return Log.WithFields(logrus.Fields{
		TraceId:  c.GetTraceId(),
		ClientId: c.GetClientId(),
		UserId:   c.GetUserId(),
		Module:   module,
		Function: funcName})
}
