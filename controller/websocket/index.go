package websocket

import "github.com/svcodestore/sv-sso-gin/service"

var (
	oauthService = service.ServiceGroup.OauthService
	userService = service.ServiceGroup.UserService
	websocketClientManagerService = service.ServiceGroup.WebsocketClientManagerService
)

