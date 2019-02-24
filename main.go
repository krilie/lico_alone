package main

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/lico603/lico-my-site-user/common/config"
	"github.com/lico603/lico-my-site-user/common/jwt"
	"github.com/lico603/lico-my-site-user/common/log"
	"net/http"
)

func main() {

	jwt.Hello()

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})

	if err := endless.ListenAndServe(":"+config.GetString("service.port"), router); err != nil {
		log.Warningln(err)
		return
	}
}
