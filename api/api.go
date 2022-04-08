package api

import (
	"github.com/gin-gonic/gin"
	"github.com/svcodestore/sv-sso-gin/model/common/response"
	"github.com/svcodestore/sv-sso-gin/model/system/request"
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

	if username != "" {
		if loginType == "login" {
			if user, err := userService.Login(username, password); err != nil {
				response.FailWithMessage(err.Error(), c)
			} else {
				token, expireAt, err := jwtService.GenerateToken(request.BaseClaims{
					UUID:        user.UUID,
					ID:          user.ID,
					Username:    user.Name,
					AuthorityId: user.LoginID,
				})
				if err != nil {
					response.FailWithMessage(err.Error(), c)
				} else {
					response.OkWithData(gin.H{
						"user":        user,
						"accessToken": token,
						"expireAt":    expireAt,
					}, c)
				}
			}
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
