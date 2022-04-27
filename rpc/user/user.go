package user

import (
	"context"
	jsoniter "github.com/json-iterator/go"
	"github.com/svcodestore/sv-sso-gin/model"
	pb "github.com/svcodestore/sv-sso-gin/proto/user"
	"github.com/svcodestore/sv-sso-gin/service"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/structpb"
	"log"
)

var (
	userService = service.ServiceGroup.UserService
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

type UserRpcServer struct {
	pb.UnimplementedUserServer
}

func RegisterUserRpcServer(s *grpc.Server) {
	pb.RegisterUserServer(s, &UserRpcServer{})
}

func (s *UserRpcServer) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserReply, error) {
	user, e := userService.UserWithId(&model.Users{
		ID: in.GetId(),
	})
	if e != nil {
		return nil, e
	}
	b, e := json.Marshal(user)
	if e != nil {
		return nil, e
	}
	var m map[string]interface{}
	json.Unmarshal(b, &m)
	u, err := structpb.NewStruct(m)
	if err != nil {
		return nil, e
	}

	log.Printf("Received user id: %v", in.GetId())
	return &pb.GetUserReply{Reply: u}, nil
}
