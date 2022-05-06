package user

import (
	"github.com/gin-gonic/gin"

	"github.com/svcodestore/sv-sso-gin/model"
	"github.com/svcodestore/sv-sso-gin/model/common/response"
	"github.com/svcodestore/sv-sso-gin/model/system/request"
	"github.com/svcodestore/sv-sso-gin/service"
)

var userService = service.ServiceGroup.UserService

func CurrentUser(c *gin.Context) {
	claims, _ := c.Get("claims")
	id := claims.(*request.CustomClaims).UserId
	user, _ := userService.UserWithId(&model.Users{ID: id})
	response.OkWithData(user, c)
}

func RegisterUser(c *gin.Context) {
	loginId := c.PostForm("loginId")
	password := c.PostForm("password")
	name := c.PostForm("name")
	lang := c.PostForm("lang")
	user, err := userService.RegisterUser(&model.UsersToSave{
		LoginID:   loginId,
		Password:  password,
		Name:      name,
		Lang:      lang,
		UpdatedBy: "-1",
	})

	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(user, c)
	}
}

func CreateUser(c *gin.Context) {
	currentUserId := c.PostForm("currentUserId")
	uid := currentUserId
	loginId := c.PostForm("loginId")
	password := c.PostForm("password")
	name := c.PostForm("name")
	lang := c.PostForm("lang")

	m := &model.UsersToSave{
		LoginID:   loginId,
		Password:  password,
		CreatedBy: uid,
		UpdatedBy: uid,
	}

	if name != "" {
		m.Name = name
	}

	if lang != "" {
		m.Lang = lang
	}

	user, err := userService.CreateUser(m)

	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(user, c)
	}
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	currentUserId := c.PostForm("currentUserId")
	uid := currentUserId

	loginId := c.PostForm("loginId")
	password := c.PostForm("password")
	name := c.PostForm("name")
	alias := c.PostForm("alias")
	phone := c.PostForm("phone")
	email := c.PostForm("email")
	lang := c.PostForm("lang")
	status := c.PostForm("status")

	updatingUser := &model.UsersToSave{
		ID:        id,
		UpdatedBy: uid,
	}

	if loginId != "" {
		updatingUser.LoginID = loginId
	}
	if password != "" {
		updatingUser.Password = password
	}
	if name != "" {
		updatingUser.Name = name
	}
	if alias != "" {
		updatingUser.Alias = alias
	}
	if phone != "" {
		updatingUser.Phone = phone
	}
	if email != "" {
		updatingUser.Email = email
	}
	if lang != "" {
		updatingUser.Lang = lang
	}
	if status != "" {
		updatingUser.Status = status == "true"
	}

	user, err := userService.UpdateUser(updatingUser)

	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(user, c)
	}
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

func GetAllUser(c *gin.Context) {
	users, _ := userService.AllUser()
	response.OkWithData(users, c)
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")
	user, _ := userService.UserWithId(&model.Users{ID: id})
	response.OkWithData(user, c)
}
