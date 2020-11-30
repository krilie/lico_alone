package ncfg

import (
	"github.com/krilie/lico_alone/common/utils/str_util"
	"os"
	"testing"
)

var cfg = NewNConfig()

func TestGetString(t *testing.T) {
	t.Log(cfg.V.WriteConfig())
	t.Log(str_util.ToJson(cfg.Cfg))
}

func TestGetInt(t *testing.T) {
	t.Log(str_util.ToJsonPretty(cfg))
	t.Log(str_util.ToJsonPretty(os.Environ()))
}

func TestSaveToFile(t *testing.T) {
	cfg.V.SetConfigFile("./config.toml")
	err := cfg.V.WriteConfig()
	t.Log(err)
}

func TestSetEnv(t *testing.T) {
	err := os.Setenv("MYAPP_CONFIG_PATH", "D:\\Users\\Administrator\\IdeaProjects\\lico_alone2\\server\\config.yaml")
	t.Log(err)
	getenv := os.Getenv("MYAPP_CONFIG_PATH")
	t.Log(getenv)
}

func TestNConfig_GetInt(t *testing.T) {
	getenv := os.Getenv("MYAPP_TEST_CONFIG")
	cfg := NewNConfig()
	err := cfg.LoadFromConfigJsonStr(getenv)
	if err != nil {
		panic(err)
	}
	err = cfg.V.WriteConfigAs("./config.json")
	if err != nil {
		panic(err)
	}
}
