package api

import (
	"github.com/gin-gonic/gin"
	"github.com/svcodestore/sv-sso-gin/api/application"
	"github.com/svcodestore/sv-sso-gin/api/organization"
	"github.com/svcodestore/sv-sso-gin/api/user"
)

type Routes struct {
}

func (*Routes) Init(g *gin.RouterGroup) {
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

	applicationsG := g.Group("applications")
	applicationsG.GET("", application.AllApplication)
	applicationG := g.Group("application")
	applicationG.POST("", application.CreateApplication)
	applicationG.DELETE("/:id", application.DeleteApplicationById)
	applicationG.GET("/:id", application.GetApplicationById)
	applicationG.PATCH("/:id", application.UpdateApplicationById)
}
