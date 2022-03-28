package service

import "github.com/svcodestore/sv-sso-gin/service/system"

type Group struct {
	JwtService    system.JwtService
	CasbinService system.CasbinService
	CryptoService system.CryptoService
	UserService   system.UserService
}

var ServiceGroup = new(Group)
