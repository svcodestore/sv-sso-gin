package oauth

import (
	"github.com/gin-gonic/gin"

	"github.com/svcodestore/sv-sso-gin/api"
	"github.com/svcodestore/sv-sso-gin/api/oauth"
)

type OAuthRoutes struct {
}

func (*OAuthRoutes) Init(r *gin.RouterGroup) {
	apiG := r.Group("login")

	apiG.POST("", api.Login)

	apiG.POST("oauth2.0", oauth.Login)

	r.POST("logout", api.Logout)

	r.GET("currentUser", api.CurrentUser)
}
