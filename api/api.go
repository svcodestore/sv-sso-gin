package api

import (
	"github.com/gin-gonic/gin"

	"github.com/svcodestore/sv-sso-gin/model/common/response"
	"github.com/svcodestore/sv-sso-gin/utils"
)

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	loginType := c.PostForm("type")
	clientId := c.PostForm("clientId")

	if username != "" {
		accessToken, refreshToken, user, err := oauthService.DoOauthLogin(username, password, loginType, clientId)
		if err == nil {
			response.OkWithData(gin.H{
				"user":         user,
				"accessToken":  accessToken,
				"refreshToken": refreshToken,
			}, c)
			return
		}
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Fail(c)
}

func Logout(c *gin.Context) {
	userId := utils.GetUserID(c)
	affected, err := oauthService.DeleteAccessTokenFromRedis(userId)

	if err == nil && affected > 0 {
		response.Ok(c)
		return
	}

	response.Fail(c)
}
