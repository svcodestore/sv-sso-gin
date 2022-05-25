package oauth

import "github.com/svcodestore/sv-sso-gin/service"

var (
	oauthService = service.ServiceGroup.OauthService
	applicationService = service.ServiceGroup.ApplicationService
	privilegeService = service.ServiceGroup.PrivilegeService
)

