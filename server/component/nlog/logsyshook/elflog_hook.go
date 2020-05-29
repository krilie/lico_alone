package logsyshook

import (
	"github.com/sirupsen/logrus"
)

type ElfLogHook struct {
	Key, Secret, Url string
}

func (e *ElfLogHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (e *ElfLogHook) Fire(entry *logrus.Entry) error {

}

func (e *ElfLogHook) PostLog() {

}

// post方法
func postLogJson(url, key, sign, data string) error {

}
