package ctl_health_check

import (
	"github.com/gin-gonic/gin"
	context_enum "github.com/krilie/lico_alone/common/com-model/context-enum"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/dig"
	"github.com/krilie/lico_alone/common/utils/str_util"
	"github.com/krilie/lico_alone/component/ndb"
	"github.com/krilie/lico_alone/component/nlog"
	"net/http"
	"time"
)

// Health 健康检查Hello
// @Summary 健康检查Hello
// @Description 健康检查Hello 返回hello字样
// @Tags 健康检查
// @ID 健康检查Hello
// @Success 200 {string} string "hello"
// @Router /health [get]
func (h *HealthCheckCtrl) Hello(c *gin.Context) {
	log := h.log.Get(context.NewContext(), "HealthCheckCtrl", "Hello")
	log.Trace("on health check")
	log.Infof("headers: %v", str_util.ToJson(c.Request.Header))
	log.Infof("remote addr: %v %v", c.Request.RemoteAddr)
	println("on health check")
	c.String(http.StatusOK, "hello")
}

// Health 健康检查Ping
// @Summary 健康检查Ping
// @Description 健康检查Ping 检查数据库是否正常 并返回启动时间
// @Tags 健康检查
// @ID 健康检查Ping
// @Success 200 {string} string "pong start time up time"
// @Router /health/ping [get]
func (h *HealthCheckCtrl) Ping(c *gin.Context) {
	log := h.log.Get(context.NewContext(), "HealthCheckCtrl", "Ping")
	err := h.db.Ping()
	if err != nil {
		log.Errorf("health ping db error %v", err)
		c.String(http.StatusInternalServerError, "数据库异常")
		return
	}
	h.log.Trace("on ping check")
	println("on ping check")
	c.String(http.StatusOK, "pong start time "+h.startTime.String())
}

type HealthCheckCtrl struct {
	startTime time.Time
	db        *ndb.NDb
	log       *nlog.NLog
}

func NewHealthCheckCtl(db *ndb.NDb, log *nlog.NLog) *HealthCheckCtrl {
	log = log.WithField(context_enum.Module.Str(), "health Ctrl")
	return &HealthCheckCtrl{startTime: time.Now(), db: db, log: log}
}

func init() {
	dig.Container.MustProvide(NewHealthCheckCtl)
}
