package health

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var startTime = time.Now()

// 健康检查 ping
func Init(engine *gin.Engine) {
	engine.GET("health/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello")
	})

	engine.GET("health/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong start time "+startTime.String())
	})
}
