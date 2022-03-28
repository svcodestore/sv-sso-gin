package system

import "github.com/svcodestore/sv-sso-gin/global"

type JwtBlacklist struct {
	global.MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
