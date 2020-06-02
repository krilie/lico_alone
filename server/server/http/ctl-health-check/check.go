package ctl_health_check

import (
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/common/dig"
	"github.com/krilie/lico_alone/component/ndb"
	"github.com/krilie/lico_alone/component/nlog"
	"net/http"
	"time"
)

// Health 健康检查
// @Summary 健康检查
// @Description 健康检查
// @Tags 健康检查
// @ID 健康检查
// @Success 200 {string} string "hello"
// @Router /health [get]
func (h *HealthCheckCtrl) Hello(c *gin.Context) {
	nlog.Log.Warn("on health check")
	println("on health check")
	c.String(http.StatusOK, "hello")
}

// Health 健康检查2
// @Summary 健康检查2
// @Description 健康检查2
// @Tags 基本信息
// @ID 健康检查2
// @Success 200 {string} string "pong start time up time"
// @Router /health [get]
func (h *HealthCheckCtrl) Ping(c *gin.Context) {
	err := h.db.Ping()
	if err != nil {
		c.String(http.StatusInternalServerError, "数据库异常")
		return
	}
	nlog.Log.Warn("on ping check")
	println("on ping check")
	c.String(http.StatusOK, "pong start time "+h.startTime.String())
}

type HealthCheckCtrl struct {
	startTime time.Time
	db        *ndb.NDb
}

func NewHealthCheckCtl(db *ndb.NDb) *HealthCheckCtrl {
	return &HealthCheckCtrl{startTime: time.Now(), db: db}
}
func init() {
	dig.Container.MustProvide(NewHealthCheckCtl)
}
