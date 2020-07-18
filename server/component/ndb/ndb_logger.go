package ndb

import (
	"context"
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm/logger"
	"time"
)

type ndbLogger struct {
	*nlog.NLog
}

func (n *ndbLogger) LogMode(level logger.LogLevel) logger.Interface {
	switch level {
	case logger.Silent:
		n.NLog.Logger.SetLevel(logrus.PanicLevel)
	case logger.Error:
		n.NLog.Logger.SetLevel(logrus.ErrorLevel)
	case logger.Info:
		n.NLog.Logger.SetLevel(logrus.InfoLevel)
	case logger.Warn:
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
	n.NLog.WithField("sql", sql).WithField("rows", rows).WithField("sql_begin", begin).WithField("err", err).Trace("sql trace")
}
