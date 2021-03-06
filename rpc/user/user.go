package user

import (
	"context"
	"github.com/svcodestore/sv-sso-gin/model/rpc/reply"
	pb "github.com/svcodestore/sv-sso-gin/proto/user"
	"github.com/svcodestore/sv-sso-gin/service"
	"github.com/svcodestore/sv-sso-gin/utils"
	"google.golang.org/grpc"
)

var (
	userService       = service.ServiceGroup.UserService
	permissionService = service.ServiceGroup.PermissionUserService
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

func (s *UserRpcServer) GetAllUser(ctx context.Context, in *pb.GetAllUserRequest) (*pb.GetAllUserReply, error) {
	users, e := userService.AllUser()
	if e != nil {
		return nil, e
	}
	u := utils.ToRpcStruct(users)

	return &pb.GetAllUserReply{Users: u}, nil
}

func (s *UserRpcServer) GetUsersByApplicationId(ctx context.Context, in *pb.GetUsersByApplicationIdRequest) (*pb.GetUsersByApplicationIdReply, error) {
	users, e := userService.UsersWithApplicationIds(in.GetApplicationId())
	if e != nil {
		return nil, e
	}
	u := utils.ToRpcStruct(users)

	return &pb.GetUsersByApplicationIdReply{Users: u}, nil
}

func (s *UserRpcServer) GetAvailableUsers(ctx context.Context, in *pb.GetAvailableUsersRequest) (*pb.GetAvailableUsersReply, error) {
	users, e := permissionService.AvailableUsersWithApplicationIds()
	if e != nil {
		return nil, e
	}
	u := utils.ToRpcStruct(users)

	return &pb.GetAvailableUsersReply{Users: u}, nil
}

func (s *UserRpcServer) GetAvailableUsersByApplicationId(ctx context.Context, in *pb.GetAvailableUsersByApplicationIdRequest) (*pb.GetAvailableUsersByApplicationIdReply, error) {
	users, e := permissionService.AvailableUsersWithApplicationIds(in.GetApplicationId())
	if e != nil {
		return nil, e
	}
	u := utils.ToRpcStruct(reply.OkWithData(users))

	return &pb.GetAvailableUsersByApplicationIdReply{Users: u}, nil
}
