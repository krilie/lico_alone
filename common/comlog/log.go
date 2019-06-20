package comlog

import (
	"context"
	"github.com/krilie/lico_alone/common/comstruct/errs"
	context2 "github.com/krilie/lico_alone/common/context"
	"github.com/sirupsen/logrus"
	"os"
)

var Log = logrus.NewEntry(logrus.New())

func init() {
	Log.Logger.SetFormatter(&logrus.JSONFormatter{})
	Log.Logger.SetLevel(logrus.DebugLevel)
	Log = Log.
		WithField("app_name", os.Getenv("BR_APP_NAME")).
		WithField("app_version", os.Getenv("BR_APP_VERSION")).
		WithField("app_host", os.Getenv("HOST_NAME"))
	Log.Infoln("log init ok")
}

// trace_id
func NewLog(ctx context.Context, moduleName string) *logrus.Entry {
	lctx := ctx.(*context2.Context)
	if lctx == nil {
		panic(errs.ErrInternal.NewWithMsg("can not get context on NewLog"))
		return nil
	}
	return Log.WithFields(logrus.Fields{
		"trace_id":  lctx.GetTraceId(),
		"client_id": lctx.GetClientId(),
		"user_id":   lctx.GetUserId(),
		"module":    moduleName})
}
