package middleware

import (
	"bufio"
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/server/http/ginutil"
	"io/ioutil"
)

// RateLimit 限制网速
func RateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 限制速度 100kb/s
		oriBody := c.Request.Body
		limitBucket := ratelimit.NewBucketWithRate(100*1024, 100*1024)
		c.Request.Body = ioutil.NopCloser(ratelimit.Reader(bufio.NewReaderSize(oriBody, 100*1024), limitBucket)) // 100kb
		defer func() { c.Request.Body = oriBody }()
		c.Next()
	}
}

func OpsLimit(ops int) gin.HandlerFunc {
	var limitOps = ops
	return func(c *gin.Context) {
		limitOps--
		defer func() { limitOps++ }()
		if limitOps < 0 {
			ginutil.AbortWithAppErr(c, errs.NewNormal().WithMsg("操作太频繁"))
			return
		}
		c.Next()
	}
}
