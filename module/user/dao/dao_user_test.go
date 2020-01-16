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
	t.Log(os.Getwd())
	dao := NewDao(config.Cfg.DB)
	strings, err := dao.GetAllValidUserId(context.Background())
	t.Log(str_util.ToJsonPretty(strings), err)
}

func TestGetFilename(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	fmt.Println("Current test filename: " + filename)
}
