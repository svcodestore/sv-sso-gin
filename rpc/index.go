package rpc

import (
	"github.com/svcodestore/sv-sso-gin/rpc/user"
	"google.golang.org/grpc"
)

func RegisterServer(s *grpc.Server)  {
	user.RegisterUserRpcServer(s)
}
