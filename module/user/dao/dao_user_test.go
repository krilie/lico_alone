package dao

import (
	"context"
	"github.com/krilie/lico_alone/common/config"
	"github.com/krilie/lico_alone/common/utils/str_util"
	"os"
	"testing"
)

func TestDao_GetAllValidUserId(t *testing.T) {
	t.Log(os.Getwd())
	dao := NewDao(config.Cfg.DB)
	strings, err := dao.GetAllValidUserId(context.Background())
	t.Log(str_util.ToJsonPretty(strings), err)
}
