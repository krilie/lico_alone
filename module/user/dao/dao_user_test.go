package dao

import (
	"context"
	"fmt"
	"github.com/krilie/lico_alone/common/config"
	"github.com/krilie/lico_alone/common/utils/str_util"
	"os"
	"runtime"
	"testing"
)

func TestDao_GetAllValidUserId(t *testing.T) {
	dao := NewDao(config.Cfg.DB)
	strings, err := dao.GetAllValidUserId(context.Background())
	t.Log(str_util.ToJsonPretty(strings), err)
}

func TestGetFilename(t *testing.T) {
	environ := os.Environ()
	fmt.Println(str_util.ToJsonPretty(environ))
	_, filename, _, _ := runtime.Caller(0)
	fmt.Println("Current test filename: " + filename)
}
