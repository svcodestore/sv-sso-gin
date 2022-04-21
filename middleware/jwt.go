package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/svcodestore/sv-sso-gin/model/common/response"
	"github.com/svcodestore/sv-sso-gin/service"
	"github.com/svcodestore/sv-sso-gin/utils"
)

var oauthService = service.ServiceGroup.OauthService

func JWTCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")

		if token == "" {
			response.UnAuth(c)
			c.Abort()
			return
		}
		if strings.HasPrefix(token, "Bearer") {
			t := strings.Split(token, " ")
			if len(t) != 2 || t[1] == "" {
				response.UnAuth(c)
				c.Abort()
				return
			}
			token = t[1]
		}
		j := utils.NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == utils.TokenExpired {
				response.UnAuthWithMessage(err.Error(), c)
				c.Abort()
				return
			}
			response.UnAuth(c)
			c.Abort()
			return
		}

		isLogin := oauthService.IsUserLogin(claims.UserId)

		if !isLogin {
			response.UnAuth(c)
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}
