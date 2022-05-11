package privilege

import (
	"github.com/gin-gonic/gin"
	"github.com/svcodestore/sv-sso-gin/model/common/response"
	"github.com/svcodestore/sv-sso-gin/model/system/request"
)

func GetAccessibleApplicationsByUserId(c *gin.Context) {
	id := c.Param("id")
	apps, _, err := privilegeApplicationService.AccessibleApplications(id)

	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(apps, c)
	}
}

func GetCurrentAccessibleApplications(c *gin.Context) {
	claims, _ := c.Get("claims")
	id := claims.(*request.CustomClaims).UserId

	apps, _, err := privilegeApplicationService.AccessibleApplications(id)

	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(apps, c)
	}
}
