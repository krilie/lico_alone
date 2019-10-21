package health

import (
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/common/config"
	"github.com/krilie/lico_alone/common/utils/time_util"
	"net/http"
)

// 健康检查 ping
func Init(engine *gin.Engine) {
	engine.GET("health/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello")
	})

	engine.GET("health/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong last update time "+time_util.GetBeijingTimeString(int64(config.GetInt("info.update_time"))))
	})
}
