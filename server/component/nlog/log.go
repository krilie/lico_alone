package nlog

import (
	"context"
	"github.com/krilie/lico_alone/common/config"
	context2 "github.com/krilie/lico_alone/common/context"
	context_enum "github.com/krilie/lico_alone/common/model/context-enum"
	"github.com/sirupsen/logrus"
	"os"
)

type NLog struct {
	*logrus.Entry
}

func NewLogger(runEnv context_enum.RunEnv, cfg config.Config) *NLog {
	var Log = logrus.NewEntry(logrus.New())
	Log.Logger.SetFormatter(&logrus.TextFormatter{})
	Log.Logger.SetLevel(logrus.Level(cfg.LogLevel))
	Log.Logger.SetOutput(os.Stdout)
	Log = Log.
		WithField(context_enum.AppName, runEnv.AppName).
		WithField(context_enum.AppVersion, runEnv.AppVersion).
		WithField(context_enum.AppHost, runEnv.AppHost)
	Log.Infoln("log init ok")
	log := &NLog{Log}
	log.SetUpLogFile(cfg.LogFile)
	return log
}

func (nlog *NLog) SetUpLogFile(f string) {
	if f == "" || f == "stdout" {
		nlog.Logger.SetOutput(os.Stdout)
		nlog.Logger.Warnln("set log out file to stdout")
		return
	}
	file, e := os.OpenFile(f, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if e != nil {
		panic(e)
		return
	}
	nlog.Logger.SetOutput(file)
	nlog.Logger.Warnln("set log out file to " + f)
}

// trace_id
func (nlog *NLog) NewLog(ctx context.Context, moduleName string, functionName string) *logrus.Entry {
	nCtx := context2.GetContextOrNew(ctx)
	return nlog.WithFields(logrus.Fields{
		context_enum.TraceId:  nCtx.GetTraceId(),
		context_enum.ClientId: nCtx.GetClientId(),
		context_enum.UserId:   nCtx.GetUserId(),
		context_enum.Module:   moduleName,
		context_enum.Function: functionName})
}

func (nlog *NLog) NewWithCtx(ctx context.Context, location ...string) *logrus.Entry {
	var module, funcName string
	if len(location) > 0 {
		module = location[0]
	}
	if len(location) > 1 {
		funcName = location[1]
	}
	c := context2.GetContextOrNew(ctx)
	return nlog.WithFields(logrus.Fields{
		context_enum.TraceId:  c.GetTraceId(),
		context_enum.ClientId: c.GetClientId(),
		context_enum.UserId:   c.GetUserId(),
		context_enum.Module:   module,
		context_enum.Function: funcName})
}
