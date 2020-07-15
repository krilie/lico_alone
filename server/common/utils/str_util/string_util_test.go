// +build auto_test

package str_util

import "testing"

func TestToJsonPretty(t *testing.T) {
	pretty := ToJsonPretty(123231)
	t.Log(pretty)
}
