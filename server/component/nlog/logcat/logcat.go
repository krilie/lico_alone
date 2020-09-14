package logcat

import (
	"io/ioutil"
	"path/filepath"
	"time"
)

// /data/log/myapplog-*.txt
// 目标 限制日志文件夹的大小 根据删除策略有计划删除日志文件

type LogFileInfo struct {
	FullName   string
	Name       string
	CreateTime time.Time
	UpdateTime time.Time
	Host       string
	Size       int64
}

func GetLogDir(logFilePath string) string {
	dir := filepath.Dir(logFilePath)
	if !(dir == "." || dir == "" || dir == "/") {
		return dir
	} else {
		return dir
	}
}

// 读取所有文件
func ReadLogFileInfo(logFilePath string) ([]LogFileInfo, error) {
	dir := GetLogDir(logFilePath)
	readDir, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

}
