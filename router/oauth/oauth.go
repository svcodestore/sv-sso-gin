package oauth

import (
	"github.com/gin-gonic/gin"

	"github.com/svcodestore/sv-sso-gin/api"
	"github.com/svcodestore/sv-sso-gin/api/application"
	"github.com/svcodestore/sv-sso-gin/api/oauth"
	"github.com/svcodestore/sv-sso-gin/api/user"
)

type OAuthRoutes struct {
}

func (*OAuthRoutes) Init(r *gin.RouterGroup) {
	r.POST("register", user.RegisterUser)
	loginG := r.Group("login")
	loginG.POST("", api.Login)

	loginOauthG := loginG.Group("oauth2.0")
	loginOauthG.POST("/grant-code", oauth.GetGrantCode)
	loginOauthG.POST("/token", oauth.GetOauthCode)

	r.GET("/application/current-application", application.GetCurrentApplication)

	// 其他后台请求
	r.GET("/current-application", application.GetCurrentApplicationByClientIdAndClientSecret)
}
