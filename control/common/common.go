package common

import (
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/application/common"
)

var appCommon common.AppCommon

type CtlCommon struct{}

func Init(g *gin.RouterGroup) {
	var ctrl CtlCommon
	com := g.Group("/common")
	com.POST("/file/upload", ctrl.uploadFile)
	com.POST("/file/delete", ctrl.deleteFile)
}
