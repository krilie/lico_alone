package strutil

import (
	"github.com/krilie/lico_alone/common/utils/jsonutil"
	"testing"
)

func TestToJsonPretty(t *testing.T) {
	pretty := jsonutil.ToJsonPretty(123231)
	t.Log(pretty)
}
