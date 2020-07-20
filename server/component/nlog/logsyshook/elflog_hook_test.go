// +build !auto_test

package logsyshook

import (
	"encoding/json"
	"testing"
	"time"
)

func TestElfLogHook_PostLog(t *testing.T) {
	var val = struct {
		T time.Time
	}{T: time.Now()}
	marshal, err := json.Marshal(val)
	t.Log(err)
	t.Log(string(marshal))
}
