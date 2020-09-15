package logcat

import (
	"github.com/krilie/lico_alone/common/utils/str_util"
	"testing"
	"time"
)

func TestReadLogFileInfo(t *testing.T) {
	info, err := ReadLogFileInfo("C:\\Users\\Administrator\\Desktop\\ccc\\")
	t.Log(err)
	t.Log(str_util.ToJson(info))
	t.Log(GetSumSize(info))
	t.Log(str_util.ToJson(GetLastShouldDeleteFile(info)))
}

func TestDeleteOverflowFile(t *testing.T) {
	logCat := &LogCat{
		logFile:   "C:\\Users\\Administrator\\Desktop\\ccc\\",
		sizeLimit: 330,
		duration:  time.Second * 7,
	}
	logCat.Start()
	logCat.wait.Wait()
}
