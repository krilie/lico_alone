// +build !auto_test

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

func TestGzipBytes(t *testing.T) {
	var img = "H4sIAAAAAAAA/3L3dLOwTExhSGH4zsBQVlZWUlJSU1PT0dExbdq0pUuXrlmz5sCBAxcvXrx3796rV6++fPny58+fv3///hsFo4BK4P///+TJ4tdIqoL///8r/uf2cw0JdnYMcDXSM2BmZGBgUPzJwsnIcIlBh4GBgQGUTRg4/jEs5ZTx2LDgYDOHsFbMiQ0LD7cLKHvN8di46Gi3hHHWnRMbFx/vV3DukvHctOTkZI3gVTEnNy09Pd0g+dQcz83Lzs62KH515+Tm5efnOzRzyXptWXFxscdkrdgsw4NMy0R4OTgUFjBc6WBbwcchkHAgtJ2bRUBgRsKrG+u5uH1sDTbkMLBmCIs9yGp4ELOtUPCc6ISEhLA9lax6BVoF7xtZOFg4OGwmHAjZWPC8aCmXxftrjMxy7sv+TPnYyFAhmWO7e7v+47XLlraKXqjUv3NgBecxPZ0HpffWxLGnvdOR0EpgWibpIztB4eDzzRmcLBt2iW64fXLr512x0RMaGc9MfbzL9ZzFw2srw9i/33qn93bd7Q9PiydPb2vQqcg8fl16jt/U7TUf+7hv7D+/f////wxsPy4y3pC8cq8q2nF6VPETzfIIpvu8S8KyKluXzjPh+XYtlilWsMWh/bumgyhzbuONu40SZa5KNQsmbr/ayGXTJhmqV8Q6g4mTzeFARaGpukMZl8qB5H7dfz/up1dwBkx30KibKhvJzDnhdvoLx62hU9k8Tz5q3bf1xV2G+YWVoen+XFNmdjaebV06M4Pt282iz4XX787Orkud41TCpjZnOsfTmddnp/J9eTapkLFLPX7ezRUXujo/cS5YyOQktWBJf9HVJxOTuFa8XDan7NXL5QvatF6tWDKt65L9fxYWBmsAAAAA//8="
	imgbytes, err := base64.StdEncoding.DecodeString(img)
	if err != nil {
		panic(err)
	}
	imgbytes = MustUnGzipBytes(imgbytes)
	markbLizo = imgbytes
}
