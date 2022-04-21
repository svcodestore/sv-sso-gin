package api

import (
	"github.com/gin-gonic/gin"
	"github.com/svcodestore/sv-sso-gin/model/common/response"
	"github.com/svcodestore/sv-sso-gin/model/system/request"
	"github.com/svcodestore/sv-sso-gin/utils"
	"strings"
)

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	loginType := c.PostForm("type")
	clientId := c.PostForm("clientId")

	if username != "" {
		if loginType == "login" {
			if user, err := userService.Login(username, password); err != nil {
				response.FailWithMessage(err.Error(), c)
			} else {
				accessToken, refreshToken, err := jwtService.GenerateToken(request.BaseClaims{
					UUID:     user.UUID,
					UserId:   user.ID,
					Username: user.Name,
					LoginId:  user.LoginID,
					ClientId: clientId,
				})

				if err == nil {
					response.OkWithData(gin.H{
						"user":         user,
						"accessToken":  accessToken,
						"refreshToken": refreshToken,
					}, c)
					return
				}

				response.FailWithMessage(err.Error(), c)
			}
		}
	}
}

func Logout(c *gin.Context) {
	t := strings.Split(c.GetHeader("Authorization"), " ")
	if len(t) > 1 {
		token := t[1]
		j := utils.NewJWT()
		claims, _ := j.ParseToken(token)
		affected, err := oauthService.DeleteAccessTokenFromRedis(claims.UserId)
		if err == nil && affected > 0 {
			response.Ok(c)
			return
		}
	}
	response.Fail(c)
}
