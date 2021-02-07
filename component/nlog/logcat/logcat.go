package logcat

import (
	"github.com/krilie/lico_alone/common/utils/timeutil"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

// /data/log/myapplog-{host}-{createTime}.txt
// 目标 限制日志文件夹的大小 根据删除策略有计划删除日志文件
type LogFileInfo struct {
	FullName   string
	Name       string
	CreateTime time.Time
	UpdateTime time.Time
	Host       string
	Size       int64
	Exclude    bool
}

func GetSumSize(infos []LogFileInfo) (sum int64) {
	for i := range infos {
		sum += infos[i].Size
	}
	return sum
}

func GetHostCount(infos []LogFileInfo, host string) int {
	var count = 0
	for _, info := range infos {
		if info.Host == host {
			count++
		}
	}
	return count
}

func GetLastShouldDeleteFile(infos []LogFileInfo) *LogFileInfo {
	var info *LogFileInfo
	// 最后修改的
	for i := range infos {
		if !infos[i].Exclude {
			if info == nil {
				info = &infos[i]
			} else {
				if infos[i].UpdateTime.Before(info.UpdateTime) {
					info = &infos[i]
				}
			}
		}
	}
	// 不在使用的最后一个
	if info != nil {
		if GetHostCount(infos, info.Host) == 1 && info.UpdateTime.After(time.Now().Add(-time.Hour*4)) {
			info.Exclude = true
			return GetLastShouldDeleteFile(infos)
		} else {
			return info
		}
	} else {
		return info
	}
}

func GetLogDir(logFilePath string) string {
	dir := filepath.Dir(logFilePath)
	if !(dir == "." || dir == "" || dir == "/") {
		return dir
	} else {
		return dir
	}
}

func GetHostNameFromFileName(fileName string) string {
	fields := strings.FieldsFunc(fileName, func(r rune) bool {
		if r == '-' || r == '.' {
			return true
		}
		return false
	})
	if len(fields) >= 2 {
		return fields[1]
	}
	return ""
}

func GetCreateTimeFromFileName(fileName string) time.Time {
	fields := strings.FieldsFunc(fileName, func(r rune) bool {
		if r == '-' || r == '.' {
			return true
		}
		return false
	})
	if len(fields) >= 2 {
		t, err := time.Parse(timeutil.StringFormat, fields[2])
		if err != nil {
			return time.Now()
		}
		return t
	}
	return time.Now()
}

// 读取所有文件
func ReadLogFileInfo(logFilePath string) (fileInfo []LogFileInfo, err error) {
	dir := GetLogDir(logFilePath)
	readDir, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, info := range readDir {
		if info.IsDir() || !strings.HasPrefix(info.Name(), "myapplog-") {
			continue
		}
		fileInfo = append(fileInfo, LogFileInfo{
			FullName:   dir + string(os.PathSeparator) + info.Name(),
			Name:       info.Name(),
			CreateTime: GetCreateTimeFromFileName(info.Name()),
			UpdateTime: info.ModTime(),
			Host:       GetHostNameFromFileName(info.Name()),
			Size:       info.Size(),
		})
	}
	return fileInfo, nil
}

func DeleteOverflowFile(filePath string, sizeLimit int64) {
	if sizeLimit < 0 {
		return
	}
	info, err := ReadLogFileInfo(filePath)
	if err != nil {
		return
	}
	sumSize := GetSumSize(info)
	if sizeLimit < sumSize {
		file := GetLastShouldDeleteFile(info)
		if len(info) == 1 || file == nil {
			return
		}
		err := os.Remove(file.FullName)
		if err != nil {
			return
		}
		if sumSize-sizeLimit > file.Size {
			DeleteOverflowFile(filePath, sizeLimit)
		}
	}
}

type LogCat struct {
	logFile   string
	sizeLimit int64
	duration  time.Duration
	onceStart sync.Once
	onceStop  sync.Once
	exitChan  chan interface{}
	wait      sync.WaitGroup
}

func (l *LogCat) Start() {
	if l.exitChan == nil {
		l.exitChan = make(chan interface{})
	}
	l.wait.Add(1)
	go func() {
		defer l.wait.Done()
		ticker := time.NewTicker(l.duration)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				println("ticker here")
				DeleteOverflowFile(l.logFile, l.sizeLimit)
			case <-l.exitChan:
				return
			}
		}
	}()
}

func (l *LogCat) Stop() {
	close(l.exitChan)
	l.wait.Wait()
}

func BeginLogFileLimit(mb int64, filePath string, duration time.Duration) (close func()) {
	if filePath == "stdout" || filePath == "stderr" || filePath == "" {
		return func() {}
	}
	logCat := &LogCat{
		logFile:   filePath,
		sizeLimit: mb * 1024 * 1024,
		duration:  duration,
	}
	logCat.Start()
	return func() {
		logCat.Stop()
	}
}
