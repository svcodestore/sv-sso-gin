package oauth

import (
	"github.com/gin-gonic/gin"
	"github.com/svcodestore/sv-sso-gin/model/common/response"
)

func Authorize(c *gin.Context) {
	responseType := c.Query("response_type")
	clientId := c.Query("client_id")
	redirectUri := c.Query("redirect_uri")
	scope := c.Query("scope")

	response.OkWithData(gin.H{
		"0": responseType,
		"1": clientId,
		"2": redirectUri,
		"3": scope,
	}, c)
}
