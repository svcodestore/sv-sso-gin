package router

import (
	"github.com/svcodestore/sv-sso-gin/router/oauth"
	"github.com/svcodestore/sv-sso-gin/router/routes"
)

type Group struct {
	OAuth oauth.OAuthRoutes
	Routes routes.Routes
}

var RouterGroup = new(Group)
