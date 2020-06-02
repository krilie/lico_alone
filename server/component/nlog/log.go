package nlog

import (
	"context"
	"fmt"
	context_enum "github.com/krilie/lico_alone/common/com-model/context-enum"
	"github.com/krilie/lico_alone/common/config"
	context2 "github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/run_env"
	"github.com/krilie/lico_alone/component/nlog/logsyshook"
	"github.com/sirupsen/logrus"
	"os"
)

type NLog struct {
	*logrus.Entry
	hook *logsyshook.ElfLogHook
}

func NewLogger(runEnv *run_env.RunEnv, cfg *config.Config, hook *logsyshook.ElfLogHook) *NLog {
	var Log = logrus.NewEntry(logrus.New())
	Log.Logger.SetFormatter(&logrus.TextFormatter{})
	Log.Logger.SetLevel(logrus.Level(cfg.LogLevel))
	Log.Logger.AddHook(hook)
	Log.Logger.SetOutput(os.Stdout)
	Log = Log.
		WithField(context_enum.AppName.Str(), runEnv.AppName).
		WithField(context_enum.AppVersion.Str(), runEnv.Version).
		WithField(context_enum.AppHost.Str(), runEnv.AppHost).
		WithField(context_enum.CommitSha.Str(), runEnv.GetShortGitCommitSha()).
		WithField(context_enum.TraceId.Str(), "")
	Log.Infoln("log init ok")
	log := &NLog{Entry: Log, hook: hook}
	log.SetUpLogFile(cfg.LogFile)
	return log
}

func (nlog *NLog) SetUpLogFile(f string) {
	if f == "" || f == "stdout" {
		nlog.Logger.SetOutput(os.Stdout)
		nlog.Warnln("set log out file to stdout")
		return
	}
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
		context_enum.Function.Str(): funcName}), hook: nlog.hook}
}

func (nlog *NLog) WithField(key string, value interface{}) *NLog {
	return &NLog{Entry: nlog.Entry.WithField(key, value), hook: nlog.hook}
}

func (nlog *NLog) CloseAndWait(ctx context.Context) {
	nlog.hook.StopPushLogWorker(ctx)
}
