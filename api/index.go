package api

import "github.com/svcodestore/sv-sso-gin/service"

var (
	oauthService           = service.ServiceGroup.OauthService
	applicationService     = service.ServiceGroup.ApplicationService
	userLoginRecordService = service.ServiceGroup.UserLoginRecordService
)
