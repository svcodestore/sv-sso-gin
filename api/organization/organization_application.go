package organization

import (
	"github.com/gin-gonic/gin"
	"github.com/svcodestore/sv-sso-gin/model"
	"github.com/svcodestore/sv-sso-gin/model/common/response"
	"github.com/svcodestore/sv-sso-gin/service"
)

var organizationApplicationService = service.ServiceGroup.OrganizationApplicationService

func CreateOrganizationApplication(c *gin.Context) {
	organizationId := c.PostForm("organizationId")
	applicationId := c.PostForm("applicationId")
	currentUserId := c.PostForm("currentUserId")
	uid := currentUserId

	organizationApplication, err := organizationApplicationService.CreateOrganizationApplication(&model.OrganizationApplication{
		OrganizationID: organizationId,
		ApplicationID:  applicationId,
		CreatedBy:      uid,
		UpdatedBy:      uid,
	})

	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(organizationApplication, c)
	}
}

func DeleteOrganizationApplicationById(c *gin.Context) {
	organizationId := c.Query("organizationId")
	applicationId := c.Query("applicationId")
	isDeleted := organizationApplicationService.DeleteOrganizationApplicationWithId(&model.OrganizationApplication{
		OrganizationID: organizationId,
		ApplicationID:  applicationId,
	})
	if isDeleted {
		response.Ok(c)
	} else {
		response.Fail(c)
	}
}

func UpdateOrganizationApplicationById(c *gin.Context) {
	organizationId := c.PostForm("organizationId")
	applicationId := c.PostForm("applicationId")
	currentUserId := c.PostForm("currentUserId")
	uid := currentUserId

	organizationApplication, err := organizationApplicationService.UpdateOrganizationApplicationWithId(&model.OrganizationApplication{
		OrganizationID: organizationId,
		ApplicationID:  applicationId,
		CreatedBy:      uid,
		UpdatedBy:      uid,
	})

	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(organizationApplication, c)
	}
}

func GetAllOrganizationApplication(c *gin.Context) {
	organizationApplications, err := organizationApplicationService.AllOrganizationApplication()

	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(organizationApplications, c)
	}
}

func GetOrganizationApplicationById(c *gin.Context) {
	organizationId := c.Query("organizationId")
	applicationId := c.Query("applicationId")
	organizationApplication, err := organizationApplicationService.OrganizationApplicationWithId(&model.OrganizationApplication{
		OrganizationID: organizationId,
		ApplicationID:  applicationId,
	})

	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(organizationApplication, c)
	}
}
