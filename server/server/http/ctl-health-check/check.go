package ctl_health_check

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var startTime = time.Now()

// Health 健康检查
// @Summary 健康检查
// @Description 健康检查
// @Tags 健康检查
// @ID 健康检查
// @Success 200 {string} string "hello"
// @Router /health [post]
func Hello(c *gin.Context) {
	c.String(http.StatusOK, "hello")
}

// Health 健康检查2
// @Summary 健康检查2
// @Description 健康检查2
// @Tags 健康检查2
// @ID 健康检查2
// @Success 200 {string} string "pong start time up time"
// @Router /health [post]
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong start time "+startTime.String())
}

// 健康检查 ping
func Init(engine *gin.Engine) {
	engine.GET("health/", Hello)
	engine.GET("health/ping", Ping)
}
