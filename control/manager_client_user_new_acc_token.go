package control

import (
	"github.com/gin-gonic/gin"
	"github.com/lico603/lico-my-site-user/control/gin_util"
)

// manager user create new acc token
func ManagerClientUserNewAccToken(c *gin.Context) {
	abort := gin_util.GetApplicationContextOrAbort(c)
	if abort == nil {
		return
	}
}
