package random

import "testing"

func TestGetAToken(t *testing.T) {
	t.Log("1: " + GetAToken())
	t.Log("2: " + GetAToken())
	t.Log("3: " + GetAToken())
	t.Log("4: " + GetAToken())
	t.Log(GetRandomNum(6))
}
