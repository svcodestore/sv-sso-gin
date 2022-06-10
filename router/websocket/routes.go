package websocket

import (
	"github.com/gin-gonic/gin"
	"github.com/svcodestore/sv-sso-gin/controller/websocket"
)

type WebsocketRoutes struct {
}

func (*WebsocketRoutes) Init(g *gin.RouterGroup) {
	g.GET("/user-activation/:id", websocket.UserActivation)
}

