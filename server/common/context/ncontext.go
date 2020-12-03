package context

import (
	"context"
	"errors"
	context_enum "github.com/krilie/lico_alone/common/com-model/context-enum"
	"github.com/krilie/lico_alone/common/run_env"
	"github.com/krilie/lico_alone/common/utils/id_util"
	"time"
)

const AppCtxValuesKey = "AppCtxValuesKey"

type AppCtxValues struct {
	context_enum.ContextValues
	StartTime time.Time   //开始调用时间
	LastTime  time.Time   //调用结束时间
	Tx        interface{} // 数据库事务对象 可有可无 跟随ctx下发 线程不安全
}

func (a *AppCtxValues) CopyFrom(p *AppCtxValues) {
	a.ContextValues = p.ContextValues
	a.StartTime = p.StartTime
	a.LastTime = p.LastTime
	a.Tx = p.Tx
}

func (a *AppCtxValues) Clone(tx interface{}) *AppCtxValues {
	var v = &AppCtxValues{}
	v.CopyFrom(a)
	v.Tx = tx
	return v
}

func NewAppCtx(parent context.Context, values ...*AppCtxValues) context.Context {
	value, ok := parent.Value(AppCtxValuesKey).(*AppCtxValues)
	if ok {
		if len(values) > 0 {
			value.CopyFrom(values[0])
		}
		return parent
	} else {
		if len(values) > 0 {
			return context.WithValue(parent, AppCtxValuesKey, values[0])
		} else {
			return context.WithValue(parent, AppCtxValuesKey, NewAppCtxValues())
		}
	}
}

func GetAppValues(ctx context.Context) *AppCtxValues {
	value, ok := ctx.Value(AppCtxValuesKey).(*AppCtxValues)
	if !ok {
		return nil
	}
	return value
}

func MustGetAppValues(ctx context.Context) *AppCtxValues {
	value, ok := ctx.Value(AppCtxValuesKey).(*AppCtxValues)
	if !ok {
		panic(errors.New("no app context value"))
	}
	return value
}

func NewAppCtxValues() *AppCtxValues {
	return &AppCtxValues{
		ContextValues: context_enum.ContextValues{
			AppName:         run_env.RunEnvLocal.AppName,
			AppVersion:      run_env.RunEnvLocal.Version,
			AppHost:         run_env.RunEnvLocal.AppHost,
			CommitSha:       run_env.RunEnvLocal.GitCommit,
			TraceId:         id_util.GetUuid(),
			ClientId:        "",
			UserId:          "",
			Module:          "", // ""
			Function:        "", // ""
			Stack:           "",
			RemoteIp:        "",
			CustomerTraceId: "",
		},
		StartTime: time.Now(),
		LastTime:  time.Now(),
		Tx:        nil,
	}
}
