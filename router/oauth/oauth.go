package oauth

import (
	"github.com/gin-gonic/gin"
	"github.com/svcodestore/sv-sso-gin/api/oauth"
	"github.com/svcodestore/sv-sso-gin/api/user"

	"github.com/svcodestore/sv-sso-gin/api"
)

type OAuthRoutes struct {
}

func (*OAuthRoutes) Init(r *gin.RouterGroup) {
	r.POST("register", user.RegisterUser)
	apiG := r.Group("login")
	apiG.POST("", api.Login)
	r.POST("logout", api.Logout)

	oauthG := apiG.Group("oauth2.0")
	oauthG.GET("/authorize", oauth.Authorize)

}
