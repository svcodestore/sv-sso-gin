package application

import (
	"github.com/gin-gonic/gin"
	"github.com/svcodestore/sv-sso-gin/model"
	"github.com/svcodestore/sv-sso-gin/model/common/response"
	"github.com/svcodestore/sv-sso-gin/service"
	"strconv"
)

var applicationService = service.ServiceGroup.ApplicationService

func CreateApplication(c *gin.Context) {
	currentUserId := c.PostForm("currentUserId")
	uid, _ := strconv.ParseInt(currentUserId, 10, 64)

	code := c.PostForm("code")
	name := c.PostForm("name")
	internalUrl := c.PostForm("internalUrl")
	homepageUrl := c.PostForm("homepageUrl")
	status := c.PostForm("status")
	redirectUris := c.PostForm("redirectUris")
	tokenFormat := c.DefaultPostForm("tokenFormat", "JWT")

	application, err := applicationService.CreateApplication(&model.Applications{
		Code:        code,
		Name:        name,
		InternalURL: internalUrl,
		HomepageURL: homepageUrl,
		Status:      status == "true",
		RedirectURIs: redirectUris,
		TokenFormat:  tokenFormat,
		CreatedBy:    uid,
		UpdatedBy:    uid,
	})

	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(application, c)
	}
}

func DeleteApplicationById(c *gin.Context) {
	id := c.Param("id")
	isDeleted := applicationService.DeleteApplication(&model.Applications{ID: id})
	if isDeleted {
		response.Ok(c)
	} else {
		response.Fail(c)
	}
}

func UpdateApplicationById(c *gin.Context) {
	id := c.Param("id")
	currentUserId := c.PostForm("currentUserId")
	uid, _ := strconv.ParseInt(currentUserId, 10, 64)

	code := c.PostForm("code")
	name := c.PostForm("name")
	internalUrl := c.PostForm("internalUrl")
	homepageUrl := c.PostForm("homepageUrl")
	status := c.PostForm("status")
	redirectUris := c.PostForm("redirectUris")
	tokenFormat := c.DefaultPostForm("tokenFormat", "JWT")

	updatingApplication := &model.Applications{
		ID: id,
		UpdatedBy: uid,
	}

	if code != "" {
		updatingApplication.Code = code
	}
	if name != "" {
		updatingApplication.Name = name
	}
	if internalUrl != "" {
		updatingApplication.InternalURL = internalUrl
	}
	if homepageUrl != "" {
		updatingApplication.HomepageURL = homepageUrl
	}
	if status != "" {
		updatingApplication.Status = status == "true"
	}
	if redirectUris != "" {
		updatingApplication.RedirectURIs = redirectUris
	}
	if tokenFormat != "" {
		updatingApplication.TokenFormat = tokenFormat
	}

	application, err := applicationService.UpdateApplicationWithId(updatingApplication)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(application, c)
	}
}

func AllApplication(c *gin.Context) {
	applications, err := applicationService.AllApplication()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(applications, c)
	}
}

func GetApplicationById(c *gin.Context) {
	id := c.Param("id")
	application, err := applicationService.ApplicationWithId(&model.Applications{ID: id})
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(application, c)
	}
}
