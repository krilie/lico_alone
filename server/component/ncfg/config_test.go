package ncfg

import (
	"github.com/krilie/lico_alone/common/dig"
	"github.com/krilie/lico_alone/common/utils/str_util"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	DigProvider()
	m.Run()
}

func TestGetString(t *testing.T) {
	dig.Container.MustInvoke(func(cfg *NConfig) {
		t.Log(cfg.V.WriteConfig())
		t.Log(str_util.ToJson(cfg.Cfg))
	})
}

func TestGetInt(t *testing.T) {
	dig.Container.MustInvoke(func(cfg *NConfig) {
		t.Log(str_util.ToJsonPretty(cfg))
		t.Log(str_util.ToJsonPretty(os.Environ()))
	})
}

func TestSaveToFile(t *testing.T) {
	dig.Container.MustInvoke(func(cfg *NConfig) {
		cfg.V.SetConfigFile("./config.toml")
		err := cfg.V.WriteConfig()
		t.Log(err)
	})
}

func TestSetEnv(t *testing.T) {
	err := os.Setenv("MYAPP_CONFIG_PATH", "D:\\Users\\Administrator\\IdeaProjects\\lico_alone2\\server\\config.yaml")
	t.Log(err)
	getenv := os.Getenv("MYAPP_CONFIG_PATH")
	t.Log(getenv)
}
