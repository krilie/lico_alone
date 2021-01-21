package id_util

import "testing"

func TestRandomStrWithCharset(t *testing.T) {
	println(RandomInt())
	println(RandomInt())
	println(RandomInt())
	println(RandomInt())
	println(RandomStr(1))
	println(RandomStr(3))
	println(RandomStr(9))
	println(RandomStr(30))
	println(RandomStr(300))
	println(RandomStr(300))
	println(RandomStr(300))
	println(RandomStr(300))
	println(RandomStr(300))
	println(RandomStrWithCharset(1, "1232123"))
	println(RandomStrWithCharset(10, "123456789"))
	println(RandomStrWithCharset(100, "abc123"))
	println(RandomStrWithCharset(23, "abc"))
	println(RandomStrWithCharset(45, "ab"))
	println(RandomIntStr())
	println(RandomIntStr())
	println(RandomIntStr())
	println(RandomIntStr())
	println(RandomIntStr())
}
