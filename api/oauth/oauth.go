package oauth

import (
	"github.com/gin-gonic/gin"
	"github.com/svcodestore/sv-sso-gin/model/system/request"
	"github.com/svcodestore/sv-sso-gin/utils"
	"strings"

	"github.com/svcodestore/sv-sso-gin/model"
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

			application, err := applicationService.ApplicationWithClientId(&model.Applications{ClientID: clientId})
			if application.ClientID == clientId {
				grantedCode := oauthService.GenerateGrantCode()
				_, err = oauthService.SaveGrantedCodeToRedis(claims.UserId, clientId, grantedCode)
				if err == nil {
					response.OkWithData(gin.H{
						"code": grantedCode,
					}, c)
					return
				}
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
		application, err := applicationService.ApplicationWithClientId(&model.Applications{ClientID: clientId})
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
		if redirectUri != application.RedirectURIs {
			response.UnAuthWithMessage("redirect uri "+redirectUri+" error", c)
			return
		}

		if clientId == application.ClientID && clientSecret == application.ClientSecret {
			userId, grantedCode, expireAt, err := oauthService.GetGrantedCodeFromRedisByClientId(clientId)
			if err != nil {
				response.UnAuthWithMessage(err.Error(), c)
				return
			}

			if grantedCode != code {
				response.UnAuthWithMessage("code err", c)
				return
			}

			if utils.IsExpire(expireAt) {
				oauthService.DeleteGrantCodeByClientId(clientId)
				response.UnAuthWithMessage("expired code", c)
				return
			}
			user, _ := userService.UserWithId(&model.Users{ID: userId})
			accessToken, refreshToken, err := jwtService.GenerateToken(request.BaseClaims{
				UserId:   user.ID,
				UUID:     user.UUID,
				LoginId:  user.LoginID,
				Username: user.Name,
				ClientId: clientId,
			})

			if err == nil {
				oauthService.DeleteGrantCodeByClientId(clientId)
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
}
