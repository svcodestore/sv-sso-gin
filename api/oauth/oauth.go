package oauth

import (
	"github.com/gin-gonic/gin"
	"github.com/svcodestore/sv-sso-gin/model/common/response"
	"net/http"
)

func Authorize(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		schema := "https"
		if c.Request.TLS == nil {
			schema = "http"
		}
		returnTo := schema + "://" + c.Request.Host + c.Request.RequestURI
		location := "http://localhost:8000/login?return_to=" + returnTo
		c.Redirect(http.StatusFound, location)
	} else {
		responseType := c.Query("response_type")
		clientId := c.Query("client_id")
		redirectUri := c.Query("redirect_uri")
		scope := c.Query("scope")
		state := c.Query("state")

		response.OkWithData(gin.H{
			"responseType": responseType,
			"clientId":     clientId,
			"redirectUri":  redirectUri,
			"scope":        scope,
			"state":        state,
		}, c)
	}
}
