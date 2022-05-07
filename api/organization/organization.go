package organization

import (
	"github.com/gin-gonic/gin"
	"github.com/svcodestore/sv-sso-gin/model"
	"github.com/svcodestore/sv-sso-gin/model/common/response"
	"github.com/svcodestore/sv-sso-gin/service"
)

var organizationService = service.ServiceGroup.OrganizationService

func CreateOrganization(c *gin.Context) {
	code := c.PostForm("code")
	name := c.PostForm("name")
	currentUserId := c.PostForm("currentUserId")
	uid := currentUserId

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
	status := c.PostForm("status")
	currentUserId := c.PostForm("currentUserId")
	uid := currentUserId

	o := &model.Organizations{
		ID:        id,
		UpdatedBy: uid,
	}

	isOnlyUpdateStatus := true

	if code != "" {
		isOnlyUpdateStatus = false
		o.Code = code
	}

	if name != "" {
		isOnlyUpdateStatus = false
		o.Name = name
	}

	var organization model.Organizations
	var err error

	if !isOnlyUpdateStatus {
		organization, err = organizationService.UpdateOrganizationWithId(o)
	}

	if err == nil {
		if status == "1" || status == "0" {
			if status == "1" {
				organization, err = organizationService.UpdateOrganizationStatusWithId(true, id, currentUserId)
			} else if status == "0" {
				organization, err = organizationService.UpdateOrganizationStatusWithId(false, id, currentUserId)
			}
		}
		if err == nil {
			response.OkWithData(organization, c)
			return
		}
	}

	response.FailWithMessage(err.Error(), c)
}

func GetAllOrganization(c *gin.Context) {
	organizations, _ := organizationService.AllOrganization()
	response.OkWithData(organizations, c)
}

func GetOrganizationById(c *gin.Context) {
	id := c.Param("id")
	organization, _ := organizationService.OrganizationWithId(id)
	response.OkWithData(organization, c)
}
