package router

import (
	"github.com/svcodestore/sv-sso-gin/router/api"
	"github.com/svcodestore/sv-sso-gin/router/oauth"
	"github.com/svcodestore/sv-sso-gin/router/websocket"
)

type Group struct {
	OAuth oauth.OAuthRoutes
	Api   api.Routes
	Websocket websocket.WebsocketRoutes
}

var RouterGroup = new(Group)
