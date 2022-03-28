package api

import (
	"github.com/svcodestore/sv-sso-gin/model/common/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login ref: https://swaggo.github.io/swaggo.io/declarative_comments_format/api_operation.html
// @Summary login
// @Description login
// @Accept  interface
// @Produce  json
// @Param string true "username ID"
// @Param string true "password ID"
// @Param string true "type ID"
// @Success 200 {object} model.Account
// @Failure 400 {object} model.HTTPError
// @Router /api/login [post]
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	loginType := c.PostForm("type")

	if loginType == "login" {
		if _, err := userService.Login(username, password); err != nil {
			response.FailWithMessage(err.Error(), c)
		} else {
			response.Ok(c)
		}
	}
}

// Logout ref: https://swaggo.github.io/swaggo.io/declarative_comments_format/api_operation.html
// @Summary logout
// @Description logout
// @Success 200 {object} model.Account
// @Failure 400 {object} model.HTTPError
// @Router /logout [post]
func Logout(c *gin.Context) {
	response.Ok(c)
}

// currentUser ref: https://swaggo.github.io/swaggo.io/declarative_comments_format/api_operation.html
// @Summary Show an account
// @Description get string by ID
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param id path string true "Account ID"
// @Success 200 {object} model.Account
// @Failure 400 {object} model.HTTPError
// @Router /accounts/{id} [get]
func CurrentUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"name":   "Serati Ma",
			"avatar": "https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png",
			"access": "admin",
		},
	})
}
