package application

import (
	"github.com/gin-gonic/gin"
	"github.com/svcodestore/sv-sso-gin/model"
	"github.com/svcodestore/sv-sso-gin/model/common/response"
)

func CreateApplicationUser(c *gin.Context) {
	applicationId := c.PostForm("applicationId")
	userId := c.PostForm("userId")
	currentUserId := c.PostForm("currentUserId")
	uid := currentUserId

	applicationUser, err := applicationUserService.CreateApplicationUser(&model.ApplicationUser{
		ApplicationID: applicationId,
		UserID:        userId,
		CreatedBy:     uid,
		UpdatedBy:     uid,
	})

	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(applicationUser, c)
	}
}

func DeleteApplicationUserById(c *gin.Context) {
	applicationId := c.Query("applicationId")
	userId := c.Query("userId")

	isDeleted := applicationUserService.DeleteApplicationUserWithId(&model.ApplicationUser{
		ApplicationID: applicationId,
		UserID:        userId,
	})

	if isDeleted {
		response.Ok(c)
	} else {
		response.Fail(c)
	}
}

func UpdateApplicationUserById(c *gin.Context) {
	applicationId := c.PostForm("applicationId")
	userId := c.PostForm("userId")
	status := c.DefaultPostForm("status", "true")
	currentUserId := c.PostForm("currentUserId")
	uid := currentUserId

	applicationUser, err := applicationUserService.UpdateApplicationUserWithId(&model.ApplicationUser{
		ApplicationID: applicationId,
		UserID:        userId,
		Status:        status == "true",
		CreatedBy:     uid,
		UpdatedBy:     uid,
	})

	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(applicationUser, c)
	}
}

func GetAllApplicationUser(c *gin.Context) {
	applicationUsers, err := applicationUserService.AllApplicationUser()

	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(applicationUsers, c)
	}
}

func GetApplicationUserById(c *gin.Context) {
	applicationId := c.Query("applicationId")
	userId := c.Query("userId")
	applicationUser, err := applicationUserService.ApplicationUserWithId(&model.ApplicationUser{
		ApplicationID: applicationId,
		UserID: userId,
	})

	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(applicationUser, c)
	}
}
