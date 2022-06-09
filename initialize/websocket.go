package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/svcodestore/sv-sso-gin/middleware"
	"github.com/svcodestore/sv-sso-gin/router"
)

func WebsocketRouters() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())

	apiGroup := r.Group("ws")

	privateApiGroup := apiGroup.Group("")
	privateApiGroup.Use(middleware.JWTCheck())
	//privateGroup.Use(middleware.CasbinCheck())
	router.RouterGroup.Api.Init(privateApiGroup)

	return r
}
