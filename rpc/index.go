package rpc

import (
	"github.com/svcodestore/sv-sso-gin/rpc/application"
	"github.com/svcodestore/sv-sso-gin/rpc/oauth"
	"github.com/svcodestore/sv-sso-gin/rpc/privilege"
	"github.com/svcodestore/sv-sso-gin/rpc/user"
	"google.golang.org/grpc"
)

func RegisterServer(s *grpc.Server) {
	user.RegisterUserRpcServer(s)
	application.RegisterApplicationRpcServer(s)
	oauth.RegisterOauthRpcServer(s)
	privilege.RegisterPrivilegeRpcServer(s)
}
