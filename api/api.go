package api

import (
	"github.com/gin-gonic/gin"
	"github.com/svcodestore/sv-sso-gin/model/common/response"
	"github.com/svcodestore/sv-sso-gin/model/system/request"
	"strings"
)

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	loginType := c.PostForm("type")
	ClientId := c.PostForm("clientId")

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
					ClientId: ClientId,
				})
				if err != nil {
					response.FailWithMessage(err.Error(), c)
				} else {
					response.OkWithData(gin.H{
						"user":         user,
						"accessToken":  accessToken,
						"refreshToken": refreshToken,
					}, c)
				}
			}
		}
	}
}

func Logout(c *gin.Context) {
	t := strings.Split(c.GetHeader("Authorization"), " ")
	if len(t) > 1 {
		token := t[1]
		_, err := oauthService.DeleteAccessTokenFromRedis(token)
		if err == nil {
			response.Ok(c)
			return
		}
	}
	response.Fail(c)
}
