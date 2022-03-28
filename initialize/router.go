package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/svcodestore/sv-sso-gin/router"
	"net/http"
)

func Routers() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "/")
	})

	api := r.Group("api")

	publicGroup := api.Group("")
	router.RouterGroup.OAuth.Init(publicGroup)

	privateGroup := api.Group("")
	//privateGroup.Use(middleware.JWTCheck())
	//privateGroup.Use(middleware.CasbinCheck())
	router.RouterGroup.Routes.Init(privateGroup)

	return r
}
