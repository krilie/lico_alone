package random

import "testing"

func TestGetAToken(t *testing.T) {
	t.Log("1: " + GetAToken())
	t.Log("2: " + GetAToken())
	t.Log("3: " + GetAToken())
	t.Log("4: " + GetAToken())
}

func TestGetRandomNum(t *testing.T) {
	t.Log(GetRandomNum(6))
	t.Log(GetRandomNum(5))
	t.Log(GetRandomNum(3))
}
