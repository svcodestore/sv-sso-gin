package organization

import (
	"github.com/gin-gonic/gin"
	"github.com/svcodestore/sv-sso-gin/model/common/response"
	"github.com/svcodestore/sv-sso-gin/service"
)

var organizationApplicationService = service.ServiceGroup.OrganizationApplicationService

func GetAllOrganizationApplication(c *gin.Context) {
	organizationApplications, _ := organizationApplicationService.AllOrganizationApplication()
	response.OkWithData(organizationApplications, c)
}
