package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/svcodestore/sv-sso-gin/model/common/response"
)

func GetAuthMenusByApplicationIdAndUserId(c *gin.Context) {
	applicationId := c.Query("applicationId")
	userId := c.Query("userId")

	reply, err := authService.GetAuthMenusWithApplicationIdAndUserId(applicationId, userId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(reply["data"], c)
	}
}
