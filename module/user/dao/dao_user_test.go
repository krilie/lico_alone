package dao

import (
	"context"
	"github.com/krilie/lico_alone/common/config"
	"github.com/krilie/lico_alone/common/utils/str_util"
	"testing"
)

func TestDao_GetAllValidUserId(t *testing.T) {
	dao := NewDao(&config.Cfg)
	strings, err := dao.GetAllValidUserId(context.Background())
	t.Log(str_util.ToJsonPretty(strings), err)
}
