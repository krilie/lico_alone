package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lico603/lico-my-site-user/common/config"
	"github.com/lico603/lico-my-site-user/common/jwt"
	"net/http"
)

func main() {

	jwt.Hello()

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})
	if err := router.Run(":" + config.GetString("service.port")); err != nil {
		return
	}
}
