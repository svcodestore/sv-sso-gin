package user

import (
	"context"
	pb "github.com/svcodestore/sv-sso-gin/proto/user"
	"github.com/svcodestore/sv-sso-gin/service"
	"github.com/svcodestore/sv-sso-gin/utils"
	"google.golang.org/grpc"
)

var (
	userService = service.ServiceGroup.UserService
)

type UserRpcServer struct {
	pb.UnimplementedUserServer
}

func RegisterUserRpcServer(s *grpc.Server) {
	pb.RegisterUserServer(s, &UserRpcServer{})
}

func (s *UserRpcServer) GetUserById(ctx context.Context, in *pb.GetUserByIdRequest) (*pb.GetUserByIdReply, error) {
	user, e := userService.UserWithId(in.GetId())
	if e != nil {
		return nil, e
	}
	u := utils.ToRpcStruct(user)

	return &pb.GetUserByIdReply{User: u}, nil
}
