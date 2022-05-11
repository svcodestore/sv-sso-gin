package application

import (
	"github.com/gin-gonic/gin"
	"github.com/svcodestore/sv-sso-gin/global"
	"github.com/svcodestore/sv-sso-gin/model"
	"github.com/svcodestore/sv-sso-gin/model/common/response"
)

func CreateApplication(c *gin.Context) {
	currentUserId := c.PostForm("currentUserId")
	uid := currentUserId

	code := c.PostForm("code")
	name := c.PostForm("name")
	internalUrl := c.PostForm("internalUrl")
	homepageUrl := c.PostForm("homepageUrl")
	status := c.PostForm("status")
	redirectUris := c.PostForm("redirectUris")
	tokenFormat := c.DefaultPostForm("tokenFormat", "JWT")

	application, err := applicationService.CreateApplication(&model.Applications{
		Code:         code,
		Name:         name,
		InternalURL:  internalUrl,
		HomepageURL:  homepageUrl,
		Status:       status == "true",
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
	uid := currentUserId

	code := c.PostForm("code")
	name := c.PostForm("name")
	internalUrl := c.PostForm("internalUrl")
	homepageUrl := c.PostForm("homepageUrl")
	status := c.PostForm("status")
	redirectUris := c.PostForm("redirectUris")
	tokenFormat := c.DefaultPostForm("tokenFormat", "JWT")

	updatingApplication := &model.Applications{
		ID:        id,
		UpdatedBy: uid,
	}

	isOnlyUpdateStatus := true

	if code != "" {
		isOnlyUpdateStatus = false
		updatingApplication.Code = code
	}
	if name != "" {
		isOnlyUpdateStatus = false
		updatingApplication.Name = name
	}
	if internalUrl != "" {
		isOnlyUpdateStatus = false
		updatingApplication.InternalURL = internalUrl
	}
	if homepageUrl != "" {
		isOnlyUpdateStatus = false
		updatingApplication.HomepageURL = homepageUrl
	}

	if redirectUris != "" {
		isOnlyUpdateStatus = false
		updatingApplication.RedirectURIs = redirectUris
	}
	if tokenFormat != "" {
		isOnlyUpdateStatus = false
		updatingApplication.TokenFormat = tokenFormat
	}

	var application model.Applications
	var err error

	if !isOnlyUpdateStatus {
		application, err = applicationService.UpdateApplicationWithId(updatingApplication)
	}

	if err == nil {
		if status == "1" || status == "0" {
			if status == "1" {
				application, err = applicationService.UpdateApplicationStatusWithId(true, id, currentUserId)
			} else if status == "0" {
				application, err = applicationService.UpdateApplicationStatusWithId(false, id, currentUserId)
			}
		}
		if err == nil {
			response.OkWithData(application, c)
			return
		}
	}

	response.FailWithMessage(err.Error(), c)
}

func GetAllApplication(c *gin.Context) {
	applications, err := applicationService.AllApplication()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(applications, c)
	}
}

func GetApplicationById(c *gin.Context) {
	id := c.Param("id")
	application, err := applicationService.ApplicationWithId(id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(application, c)
	}
}

func GetCurrentApplication(c *gin.Context) {
	id := global.CONFIG.System.Id
	application, err := applicationService.ApplicationWithId(id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		application.ClientSecret = "***"
		application.CreatedByUser = model.UsersWithoutModInfo{}
		application.UpdatedByUser = application.CreatedByUser
		response.OkWithData(application, c)
	}
}
