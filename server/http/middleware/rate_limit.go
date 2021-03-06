package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	com_model "github.com/krilie/lico_alone/common/com-model"
	"github.com/krilie/lico_alone/common/errs"
	"io/ioutil"
	"time"
)

// RateLimit 限制网速
func RateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 限制速度 100kb/s
		oriBody := c.Request.Body
		limitBucket := ratelimit.NewBucketWithRate(100*1024, 100*1024)
		c.Request.Body = ioutil.NopCloser(ratelimit.Reader(oriBody, limitBucket)) // 100kb
		defer func() { c.Request.Body = oriBody }()
		c.Next()
	}
}

// RequestOpsLimit RateLimit 限制网速
func RequestOpsLimit() gin.HandlerFunc {
	limitBucket := ratelimit.NewBucket(time.Millisecond*100, 30) // 0.1秒三次
	return func(c *gin.Context) {
		// 限制速度
		limitBucket.Wait(1)
		c.Next()
	}
}

func OpsLimit(ops int) gin.HandlerFunc {
	var limitOps = ops
	return func(c *gin.Context) {
		limitOps--
		defer func() { limitOps++ }()
		if limitOps < 0 {
			c.AbortWithStatusJSON(200, com_model.NewRetFromErr(errs.NewNormal().WithMsg("操作太频繁")))
			return
		}
		c.Next()
	}
}
