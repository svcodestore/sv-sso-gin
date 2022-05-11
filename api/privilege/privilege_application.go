package privilege

import (
	"github.com/gin-gonic/gin"
	"github.com/svcodestore/sv-sso-gin/model/common/response"
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
