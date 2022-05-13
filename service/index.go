package service

import "github.com/svcodestore/sv-sso-gin/service/system"

type Group struct {
	JwtService                     system.JwtService
	CasbinService                  system.CasbinService
	CryptoService                  system.CryptoService
	UserService                    system.UserService
	OrganizationService            system.OrganizationService
	ApplicationService             system.ApplicationService
	OrganizationApplicationService system.OrganizationApplicationService
	ApplicationUserService         system.ApplicationUserService
	OauthService                   system.OauthService
	PrivilegeService               system.PrivilegeService
	PrivilegeApplicationService    system.PrivilegeApplicationService
	PermissionUserService          system.PermissionUserService
}

var ServiceGroup = new(Group)
