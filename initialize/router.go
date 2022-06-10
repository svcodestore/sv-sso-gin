package initialize

import (
	"github.com/gin-gonic/gin"

	"github.com/svcodestore/sv-sso-gin/middleware"
	"github.com/svcodestore/sv-sso-gin/router"
)

func Routers() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())

	wsGroup := r.Group("ws")
	router.RouterGroup.Websocket.Init(wsGroup)

	apiGroup := r.Group("api")

	publicApiGroup := apiGroup.Group("")
	router.RouterGroup.OAuth.Init(publicApiGroup)

	privateApiGroup := apiGroup.Group("")
	privateApiGroup.Use(middleware.JWTCheck())
	//privateGroup.Use(middleware.CasbinCheck())
	router.RouterGroup.Api.Init(privateApiGroup)

	return r
}
