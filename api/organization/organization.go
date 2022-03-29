package organization

import (
	"github.com/gin-gonic/gin"
	"github.com/svcodestore/sv-sso-gin/model"
	"github.com/svcodestore/sv-sso-gin/model/common/response"
	"github.com/svcodestore/sv-sso-gin/service"
	"strconv"
)

var organizationService = service.ServiceGroup.OrganizationService

func CreateOrganization(c *gin.Context) {
	code := c.PostForm("code")
	name := c.PostForm("name")
	currentUserId := c.PostForm("currentUserId")
	uid, _ := strconv.ParseInt(currentUserId, 10, 64)

	organization, err := organizationService.CreateOrganization(&model.Organizations{
		Code:      code,
		Name:      name,
		CreatedBy: uid,
		UpdatedBy: uid,
	})

	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(organization, c)
	}
}

func DeleteOrganizationById(c *gin.Context) {
	id := c.Param("id")
	isDeleted := organizationService.DeleteOrganizationWithId(&model.Organizations{ID: id})
	if isDeleted {
		response.Ok(c)
	} else {
		response.Fail(c)
	}
}

func UpdateOrganizationById(c *gin.Context) {
	id := c.Param("id")
	code := c.PostForm("code")
	name := c.PostForm("name")
	currentUserId := c.PostForm("currentUserId")
	uid, _ := strconv.ParseInt(currentUserId, 10, 64)

	organization, err := organizationService.UpdateOrganizationWithId(&model.Organizations{
		ID:        id,
		Code:      code,
		Name:      name,
		UpdatedBy: uid,
	})

	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(organization, c)
	}
}

func GetAllOrganization(c *gin.Context) {
	organizations, _ := organizationService.AllOrganization()
	response.OkWithData(organizations, c)
}

func GetOrganizationById(c *gin.Context) {
	id := c.Param("id")
	organization, _ := organizationService.OrganizationWithId(&model.Organizations{ID: id})
	response.OkWithData(organization, c)
}
