package router

import (
	"github.com/svcodestore/sv-sso-gin/router/api"
	"github.com/svcodestore/sv-sso-gin/router/oauth"
)

type Group struct {
	OAuth oauth.OAuthRoutes
	Api   api.Routes
}

var RouterGroup = new(Group)
