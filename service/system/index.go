package system

import jsoniter "github.com/json-iterator/go"

var (
	applicationService             = ApplicationService{}
	jwtService                     = JwtService{}
	userService                    = UserService{}
	organizationService            = OrganizationService{}
	organizationApplicationService = OrganizationApplicationService{}
	oauthService                   = OauthService{}
	privilegeApplicationService    = PrivilegeApplicationService{}
	json                           = jsoniter.ConfigCompatibleWithStandardLibrary
)
