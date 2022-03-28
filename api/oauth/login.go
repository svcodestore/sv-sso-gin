package oauth

import (
	"github.com/gin-gonic/gin"
	"github.com/svcodestore/sv-sso-gin/model/common/response"
)

func Login(c *gin.Context) {
	response.Ok(c)
}
