package fileutils

import (
	"encoding/base64"
	"io/ioutil"
	"testing"
)

func TestWaterMarkOne3(t *testing.T) {
	file, _ := ioutil.ReadFile("./lizotop.gif")
	println(base64.StdEncoding.EncodeToString(MustGzipBytes(file)))
	println(base64.StdEncoding.EncodeToString(file))
}
