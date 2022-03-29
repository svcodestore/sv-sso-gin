package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/svcodestore/sv-sso-gin/api/organization"

	"github.com/svcodestore/sv-sso-gin/api/user"
)

type Routes struct {
}

func (*Routes) Init(g *gin.RouterGroup) {
	usersG := g.Group("users")
	usersG.GET("", user.GetAllUser)
	userG := g.Group("user")
	userG.GET("/:id", user.GetUserById)
	userG.POST("", user.CreateUser)
	userG.DELETE("/:id", user.DeleteUserById)

	organizationsG := g.Group("organizations")
	organizationsG.GET("", organization.GetAllOrganization)
	organizationG := g.Group("organization")
	organizationG.POST("", organization.CreateOrganization)
	organizationG.DELETE("/:id", organization.DeleteOrganizationById)
	organizationG.GET("/:id", organization.GetOrganizationById)
	organizationG.PATCH("/:id", organization.UpdateOrganizationById)
}
