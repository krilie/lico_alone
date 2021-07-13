package ndb

import (
	"context"
	"fmt"
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm/logger"
	"time"
)

// ndbLogger 给gorm.Db用的log组件包装
type ndbLogger struct {
	*nlog.NLog
}

func (n *ndbLogger) LogMode(level logger.LogLevel) logger.Interface {
	switch level {
	case logger.Silent:
		n.NLog.Logger.SetLevel(logrus.PanicLevel)
	case logger.Error:
		n.NLog.Logger.SetLevel(logrus.ErrorLevel)
	case logger.Warn:
		n.NLog.Logger.SetLevel(logrus.WarnLevel)
	case logger.Info:
		n.NLog.Logger.SetLevel(logrus.TraceLevel)
	default:
		n.NLog.Logger.SetLevel(logrus.InfoLevel)
	}
	return n
}

func (n *ndbLogger) Info(ctx context.Context, s string, i ...interface{}) {
	n.NLog.Infof(s, i...)
}

func (n *ndbLogger) Warn(ctx context.Context, s string, i ...interface{}) {
	n.NLog.Warnf(s, i...)
}

func (n *ndbLogger) Error(ctx context.Context, s string, i ...interface{}) {
	n.NLog.Errorf(s, i...)
}

func (n *ndbLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	sql, rows := fc()
	if n.NLog.Logger.Level >= logrus.TraceLevel {
		fmt.Printf("sql<==  %v \n", sql)
		fmt.Printf("sql==>  %v \n", rows)
		n.NLog.WithField("sql", sql).WithField("rows", rows).WithField("sql_begin", begin).WithField("err", err).Trace("sql trace")
	}
}
