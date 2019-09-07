package config

import (
	"testing"
)

func TestGetString(t *testing.T) {
	t.Log(Cfg)
	t.Log(v.GetStringSlice("cors.allow_methods"))
}
