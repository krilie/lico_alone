package config

import (
	"github.com/krilie/lico_alone/common/utils/str_util"
	"os"
	"testing"
)

func TestGetString(t *testing.T) {
	t.Log(Cfg)
	t.Log(v.GetStringSlice("cors.allow_methods"))
	str := v.WriteConfig()
	t.Log(str)
}

func TestGetInt(t *testing.T) {
	t.Log(Cfg)
	t.Log(str_util.ToJsonPretty(Cfg))
	t.Log(str_util.ToJsonPretty(os.Environ()))
}

func TestSaveToFile(t *testing.T) {
	v.SetConfigFile("./config.yml")
	err := v.WriteConfig()
	t.Log(err)
}

func TestSetEnv(t *testing.T) {
	err := os.Setenv("APP_CONFIG_PATH", "D:\\Users\\Administrator\\IdeaProjects\\lico_alone2\\server\\config.yaml")
	t.Log(err)
	getenv := os.Getenv("APP_CONFIG_PATH")
	t.Log(getenv)
}
