package user

import (
	"github.com/gin-gonic/gin"

	"github.com/svcodestore/sv-sso-gin/model"
	"github.com/svcodestore/sv-sso-gin/model/common/response"
	"github.com/svcodestore/sv-sso-gin/service"
)

var userService = service.ServiceGroup.UserService

func GetAllUser(c *gin.Context) {
	users, _ := userService.AllUser()
	response.OkWithData(users, c)
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")
	user, _ := userService.UserWithId(&model.Users{ID: id})
	response.OkWithData(user, c)
}

func DeleteUserById(c *gin.Context) {
	id := c.Param("id")
	isDeleted := userService.DeleteUserWithId(&model.Users{ID: id})
	if isDeleted {
		response.Ok(c)
	} else {
		response.Fail(c)
	}
}

func CreateUser(c *gin.Context) {
	loginId := c.PostForm("loginId")
	password := c.PostForm("password")
	name := c.PostForm("name")
	lang := c.PostForm("lang")
	user, err := userService.CreateUser(&model.UsersToSave{
		LoginID:  loginId,
		Password: password,
		Name:     name,
		Lang:     lang,
	})

	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(user, c)
	}
}
