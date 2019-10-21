package dao

import (
	"context"
	"github.com/krilie/lico_alone/common/cmodel"
	"github.com/krilie/lico_alone/common/config"
	"github.com/krilie/lico_alone/module/config/model"
	"testing"
	"time"
)

func TestDao_GetConfigByName(t *testing.T) {
	dao := NewDao(config.Cfg.DB)
	name, err := dao.GetConfigByName(context.Background(), model.CommonIsInitData)
	t.Log(name, err)
}

func TestDao_CreateConfig(t *testing.T) {
	dao := NewDao(config.Cfg.DB)
	err := dao.CreateConfig(context.Background(), &model.Config{
		Model: cmodel.Model{
			Id:         "123",
			CreateTime: time.Now(),
		},
		Name:  "123",
		Value: "111",
	})
	t.Log(err)
}
func TestNewDao(t *testing.T) {
	dao := NewDao(config.Cfg.DB)
	err := dao.UpdateConfig(context.Background(), &model.Config{
		Model: cmodel.Model{
			Id:         "123",
			CreateTime: time.Time{},
		},
		Name:  "1111111111",
		Value: "22222222222",
	})
	t.Log(err)
}
func TestDao_DeleteConfig(t *testing.T) {
	dao := NewDao(config.Cfg.DB)
	err := dao.DeleteConfig(context.Background(), "1111111111")
	t.Log(err)
}
