package api

import "github.com/svcodestore/sv-sso-gin/service"

var (
	userService = service.ServiceGroup.UserService
	//cryptoService       = service.ServiceGroup.CryptoService
	//casbinService       = service.ServiceGroup.CasbinService
	jwtService = service.ServiceGroup.JwtService
	oauthService = service.ServiceGroup.OauthService
)
