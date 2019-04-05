package main

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/lico603/lico_user/common/config"
	"github.com/lico603/lico_user/common/log"
	"github.com/lico603/lico_user/control"
	"net/http"
)

func main() {

	control.LocalRouter.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})

	//是否有ssl.public_key ssl.private_key
	pubKey := config.GetString("ssl.public_key")
	priKey := config.GetString("ssl.private_key")
	if pubKey == "" || priKey == "" {
		if err := endless.ListenAndServe(":"+config.GetString("service.port"), control.LocalRouter); err != nil {
			log.Warningln(err)
			return
		}
	} else {
		if err := endless.ListenAndServeTLS(":"+config.GetString("service.port"), pubKey, priKey, control.LocalRouter); err != nil {
			log.Warningln(err)
			return
		}
	}
}
