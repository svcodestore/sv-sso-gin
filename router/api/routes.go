package api

import (
	"github.com/gin-gonic/gin"
	"github.com/svcodestore/sv-sso-gin/api/application"
	"github.com/svcodestore/sv-sso-gin/api/organization"
	"github.com/svcodestore/sv-sso-gin/api/privilege"
	"github.com/svcodestore/sv-sso-gin/api/user"
)

type Routes struct {
}

func (*Routes) Init(g *gin.RouterGroup) {
	myG := g.Group("my")
	myG.GET("/applications", privilege.GetCurrentAccessibleApplications)

	usersG := g.Group("users")
	usersG.GET("", user.GetAllUser)
	userG := g.Group("user")
	userG.POST("", user.CreateUser)
	userG.DELETE("/:id", user.DeleteUserById)
	userG.GET("/:id", user.GetUserById)
	userG.PATCH("/:id", user.UpdateUser)
	userG.GET("/current-user", user.CurrentUser)

	organizationsG := g.Group("organizations")
	organizationsG.GET("", organization.GetAllOrganization)
	organizationG := g.Group("organization")
	organizationG.POST("", organization.CreateOrganization)
	organizationG.DELETE("/:id", organization.DeleteOrganizationById)
	organizationG.GET("/:id", organization.GetOrganizationById)
	organizationG.PATCH("/:id", organization.UpdateOrganizationById)

	organizationApplicationsG := g.Group("organization-applications")
	organizationApplicationsG.GET("", organization.GetAllOrganizationApplication)
	organizationApplicationG := g.Group("organization-application")
	organizationApplicationG.POST("", organization.CreateOrganizationApplication)
	organizationApplicationG.DELETE("", organization.DeleteOrganizationApplicationById)
	organizationApplicationG.GET("", organization.GetOrganizationApplicationById)
	organizationApplicationG.PATCH("", organization.UpdateOrganizationApplicationById)

	applicationsG := g.Group("applications")
	applicationsG.GET("", application.GetAllApplication)
	applicationG := g.Group("application")
	applicationG.POST("", application.CreateApplication)
	applicationG.DELETE("/:id", application.DeleteApplicationById)
	applicationG.GET("/:id", application.GetApplicationById)
	applicationG.PATCH("/:id", application.UpdateApplicationById)

	applicationUsersG := g.Group("application-users")
	applicationUsersG.GET("", application.GetAllApplicationUser)
	applicationUserG := g.Group("application-user")
	applicationUserG.POST("", application.CreateApplicationUser)
	applicationUserG.DELETE("", application.DeleteApplicationUserById)
	applicationUserG.GET("", application.GetApplicationUserById)
	applicationUserG.PATCH("", application.UpdateApplicationUserById)

	privilegeApplicationG := g.Group("privilege-application")
	privilegeApplicationG.GET("/:id", privilege.GetAccessibleApplicationsByUserId)
}
