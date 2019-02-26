package pswd_md5

import (
	"crypto/md5"
	"encoding/hex"
)

//对密码进行md5加密 盐值
//原密码加盐值，返回加密过后的密码
func GetMd5Password(ori string, salt string) string {
	tagPswd := md5.Sum([]byte(ori + salt + ori))
	tagPswdHex := hex.EncodeToString(tagPswd[:])
	return tagPswdHex
}
