package middleware

import (
	"fmt"
	"testing"
)

func TestOpsLimit(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("defer2 %v \n", err)
		}
	}()
	defer println("-0----")
	defer func() {
		defer println("defer in defer")
		if err := recover(); err != nil {
			fmt.Printf("defer %v \n", err)
			panic(err)
		}
	}()
	TPanic()
}
func TPanic() {
	panic("okkkk")
}

//func TestRateLimit(t *testing.T) {
//	engine := gin.New()
//	engine.POST("/upload", RateLimit(), func(c *gin.Context) {
//		err := c.Request.ParseMultipartForm(1024 * 1024 * 5) // 5mb
//		if err != nil {
//			ginutil.ReturnFailure(c, errs.ErrorInternal, err.Error())
//			return
//		}
//		file, err := c.FormFile("file")
//		if err != nil {
//			ginutil.ReturnFailure(c, errs.ErrorParam, err.Error())
//			return
//		}
//		open, err := file.Open()
//		if err != nil {
//			ginutil.ReturnFailure(c, errs.ErrorParam, err.Error())
//			return
//		}
//		defer open.Close()
//		all, err := ioutil.ReadAll(open)
//		t.Log(string(all))
//	})
//	err := engine.Run(":80")
//	t.Log(err)
//}
//
//func TestRateLimit2(t *testing.T) {
//	engine := gin.New()
//	engine.POST("/upload", RateLimit(), func(c *gin.Context) {
//		reader, err := c.Request.MultipartReader()
//		if err != nil {
//			ginutil.ReturnWithErr(c, err)
//			return
//		}
//		for {
//			p, err := reader.NextPart()
//			if err == io.EOF {
//				ginutil.ReturnWithErr(c, errs.NewParamError().WithMsg("no file"))
//				return
//			}
//			if err != nil {
//				ginutil.ReturnWithErr(c, err)
//				return
//			}
//
//			name := p.FormName()
//			if name == "" {
//				continue
//			}
//			if name != "file" {
//				continue
//			}
//
//			filename := p.FileName()
//			if filename == "" {
//				continue
//			}
//
//			return
//		}
//	})
//	err := engine.Run(":80")
//	t.Log(err)
//}
