package pswd_md5

import "testing"

func TestGetMd5Password(t *testing.T) {
	t.Log(GetMd5Password("12345678", "234343"))
	t.Log(GetMd5Password("12345678", "234343"))
	t.Log(GetMd5Password("12345678", "234342"))
}
