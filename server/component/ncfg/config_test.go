package ncfg

import (
	"github.com/krilie/lico_alone/common/utils/jsonutil"
	"github.com/mitchellh/mapstructure"
	"os"
	"testing"
)

var cfg = NewNConfig()

func TestNewNConfig2(t *testing.T) {
	eCfg := cfg.V.GetStringMap("email")
	var emailConfig = Email{}
	err := mapstructure.Decode(eCfg, &emailConfig)
	if err != nil {
		panic(err.Error())
	}
}

func TestGetString(t *testing.T) {
	t.Log(cfg.V.WriteConfig())
}

func TestGetInt(t *testing.T) {
	t.Log(jsonutil.ToJsonPretty(cfg))
	t.Log(jsonutil.ToJsonPretty(os.Environ()))
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
	err := cfg.LoadFromConfigTomlStr(getenv)
	if err != nil {
		panic(err)
	}
	err = cfg.V.WriteConfigAs("./config.toml")
	if err != nil {
		panic(err)
	}
}

func TestNewNConfigByCfgStr(t *testing.T) {
	dbCfg := cfg.GetDbCfg()
	t.Log(jsonutil.ToJsonPretty(dbCfg))
}
