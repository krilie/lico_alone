package brlog

import (
	"context"
	context2 "github.com/krilie/lico_alone/common/context"
	"github.com/sirupsen/logrus"
	"os"
)

var (
	app_name    = "app_name"
	app_version = "app_version"
	app_host    = "app_host"
	trace_id    = "trace_id"
	client_id   = "client_id"
	user_id     = "user_id"
	module      = "module"
	function    = "function"
	stack       = "stack"
)

var Log = logrus.NewEntry(logrus.New())

func init() {
	Log.Logger.SetFormatter(&logrus.JSONFormatter{TimestampFormat: "2006-01-02T15:04:05.000Z07:00"})
	Log.Logger.SetLevel(logrus.DebugLevel)
	//
	//file, e := os.OpenFile("./log.txt", os.O_CREATE|os.O_APPEND, os.ModeAppend)
	//if e != nil {
	//	panic(e)
	//	return
	//}
	Log.Logger.SetOutput(os.Stdout)
	Log = Log.
		WithField(app_name, os.Getenv("BR_APP_NAME")).
		WithField(app_version, os.Getenv("BR_APP_VERSION")).
		WithField(app_host, os.Getenv("HOST_NAME"))
	Log.Infoln("log init ok")
}

// trace_id
func NewLog(ctx context.Context, moduleName string, functionName string) *logrus.Entry {
	bctx := ctx.(*context2.Context)
	return Log.WithFields(logrus.Fields{
		trace_id:  bctx.GetTraceId(),
		client_id: bctx.GetClientId(),
		user_id:   bctx.GetUserId(),
		module:    moduleName,
		function:  functionName})
}
