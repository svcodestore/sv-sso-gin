package websocket

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/svcodestore/sv-sso-gin/service"
)

var (
	oauthService           = service.ServiceGroup.OauthService
	userService            = service.ServiceGroup.UserService
	applicationService     = service.ServiceGroup.ApplicationService
	userLoginRecordService = service.ServiceGroup.UserLoginRecordService
	json                   = jsoniter.ConfigCompatibleWithStandardLibrary
)
