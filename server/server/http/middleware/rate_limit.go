package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"io/ioutil"
)

// RateLimit 限制网速
func RateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 限制速度 100kb/s
		oriBody := c.Request.Body
		limitBucket := ratelimit.NewBucketWithRate(100*1024, 100*1024)
		c.Request.Body = ioutil.NopCloser(ratelimit.Reader(oriBody, limitBucket))
		defer func() { c.Request.Body = oriBody }()
		c.Next()
	}
}
