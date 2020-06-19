package logsyshook

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

type SyslogHook struct {
}

func NewSyslogHook() *SyslogHook {
	return &SyslogHook{}
}

func (hook *SyslogHook) Fire(entry *logrus.Entry) error {
	//fields := entry.Data
	//logTime := entry.Time
	//content := entry.Message

	line, err := entry.String()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read entry, %v", err)
		return err
	}
	print("Hook************ " + line)
	return nil
	//switch entry.Level {
	//case logrus.PanicLevel:
	//	return hook.Writer.Crit(line)
	//case logrus.FatalLevel:
	//	return hook.Writer.Crit(line)
	//case logrus.ErrorLevel:
	//	return hook.Writer.Err(line)
	//case logrus.WarnLevel:
	//	return hook.Writer.Warning(line)
	//case logrus.InfoLevel:
	//	return hook.Writer.Info(line)
	//case logrus.DebugLevel:
	//	return hook.Writer.Debug(line)
	//default:
	//	return nil
	//}
}

func (hook *SyslogHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
