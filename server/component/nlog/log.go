package nlog

import (
	"context"
	"fmt"
	context_enum "github.com/krilie/lico_alone/common/com-model/context-enum"
	context2 "github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/utils/time_util"
	"github.com/krilie/lico_alone/component/ncfg"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type NLog struct {
	*logrus.Entry
}

func NewLogger(cfg *ncfg.NConfig) *NLog {

	logCfg := cfg.Cfg.Log

	var Log = logrus.NewEntry(logrus.New())
	Log.Logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat:  time.RFC3339Nano,
		DisableTimestamp: false,
		DataKey:          "",
		FieldMap:         nil,
		CallerPrettyfier: nil,
		PrettyPrint:      false,
	})
	Log.Logger.SetLevel(logrus.Level(logCfg.LogLevel))
	Log.Logger.SetOutput(os.Stdout)
	Log = Log.
		WithField(context_enum.AppName.Str(), cfg.RunEnv.AppName).
		WithField(context_enum.AppVersion.Str(), cfg.RunEnv.Version).
		WithField(context_enum.AppHost.Str(), cfg.RunEnv.AppHost).
		WithField(context_enum.CommitSha.Str(), cfg.RunEnv.GetShortGitCommitSha()).
		WithField(context_enum.TraceId.Str(), "")
	Log.Infoln("log init ok")
	log := &NLog{Entry: Log}
	log.SetUpLogFile(cfg.Cfg.Log.LogFile)
	return log
}

func (nlog *NLog) SetUpLogFile(f string) {
	if f == "" || f == "stdout" {
		nlog.Logger.SetOutput(os.Stdout)
		nlog.Warnln("set log out file to stdout")
		return
	}
	dir := filepath.Dir(f)
	if !(dir == "." || dir == "" || dir == "/") {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			panic(err)
			return
		}
	}
	hostname, _ := os.Hostname()
	var tagStr = hostname + "-" + time.Now().Format(time_util.StringFormat)
	f = strings.ReplaceAll(f, "*", tagStr)
	file, e := os.OpenFile(f, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if e != nil {
		panic(e)
		return
	}
	nlog.Logger.SetOutput(file)
	nlog.Warnln("set log out file to " + f)
}

func (nlog *NLog) Get(ctx context.Context, location ...string) *NLog {
	var module, funcName string
	val, ok := nlog.Entry.Data[context_enum.Module.Str()]
	if ok {
		module = fmt.Sprint(val)
	}
	val, ok = nlog.Entry.Data[context_enum.Function.Str()]
	if ok {
		funcName = fmt.Sprint(val)
	}
	nCtx := context2.GetContextOrNew(ctx)
	if nCtx.Module != "" {
		module = nCtx.Module
	}
	if nCtx.Function != "" {
		funcName = nCtx.Function
	}
	if len(location) > 0 {
		module = location[0]
	}
	if len(location) > 1 {
		funcName = location[1]
	}
	return &NLog{Entry: nlog.WithFields(logrus.Fields{
		//context_enum.AppName.Str():    nCtx.AppName,
		//context_enum.AppVersion.Str(): nCtx.AppVersion,
		//context_enum.AppHost.Str():    nCtx.AppHost,
		//context_enum.CommitSha.Str(): nCtx.CommitSha,
		context_enum.TraceId.Str():  nCtx.GetTraceId(),
		context_enum.ClientId.Str(): nCtx.GetClientId(),
		context_enum.UserId.Str():   nCtx.GetUserId(),
		context_enum.Stack.Str():    nCtx.Stack,
		context_enum.RemoteIp.Str(): nCtx.RemoteIp,
		context_enum.Module.Str():   module,
		context_enum.Function.Str(): funcName})}
}

func (nlog *NLog) WithField(key string, value interface{}) *NLog {
	return &NLog{Entry: nlog.Entry.WithField(key, value)}
}

func (nlog *NLog) WithFuncName(value interface{}) *NLog {
	return &NLog{Entry: nlog.Entry.WithField(context_enum.Function.Str(), value)}
}

func (nlog *NLog) WithError(value interface{}) *NLog {
	return &NLog{Entry: nlog.Entry.WithField(context_enum.Err.Str(), value)}
}

func (nlog *NLog) CloseAndWait(ctx context.Context) {

}
