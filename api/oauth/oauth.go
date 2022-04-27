package oauth

import (
	"github.com/gin-gonic/gin"
	"github.com/svcodestore/sv-sso-gin/utils"
	"strings"

	"github.com/svcodestore/sv-sso-gin/model/common/response"
	"github.com/svcodestore/sv-sso-gin/service"
)

var (
	oauthService       = service.ServiceGroup.OauthService
	applicationService = service.ServiceGroup.ApplicationService
	jwtService         = service.ServiceGroup.JwtService
	userService        = service.ServiceGroup.UserService
)

func GetGrantCode(c *gin.Context) {
	responseType := c.PostForm("responseType")
	clientId := c.PostForm("clientId")
	//redirectUri := c.PostForm("redirectUri")
	//scope := c.PostForm("scope")
	//state := c.PostForm("state")

	if responseType == "code" {
		authorization := c.GetHeader("Authorization")
		if authorization != "" {
			t := strings.Split(authorization, " ")
			token := t[1]
			j := utils.NewJWT()
			claims, err := j.ParseToken(token)
			grantedCode, err := oauthService.DoGenerateGrantCode(claims.UserId, clientId)

			if err == nil {
				response.OkWithData(gin.H{
					"code": grantedCode,
				}, c)
				return
			}
			response.FailWithMessage(err.Error(), c)
		}
	}
}

func GetOauthCode(c *gin.Context) {
	grantType := c.Query("grant_type")
	clientId := c.Query("client_id")
	clientSecret := c.Query("client_secret")
	code := c.Query("code")
	redirectUri := c.Query("redirect_uri")

	if grantType == "authorization_code" {
		if code == "" {
			response.UnAuthWithMessage("empty code", c)
			return
		}
		accessToken, refreshToken, user, err := oauthService.DoGenerateOauthCode(clientId, clientSecret, code, redirectUri)
		if err == nil {
			response.OkWithData(gin.H{
				"accessToken":  accessToken,
				"refreshToken": refreshToken,
				"user":         user,
			}, c)
			return
		}

		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.UnAuthWithMessage("client id and secret error", c)
	}
}
