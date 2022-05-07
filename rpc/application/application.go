package application

import (
	jsoniter "github.com/json-iterator/go"
	pb "github.com/svcodestore/sv-sso-gin/proto/application"
	"github.com/svcodestore/sv-sso-gin/service"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/structpb"
	"log"
)

var (
	applicationService = service.ServiceGroup.ApplicationService
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

type ApplicationRpcServer struct {
	pb.UnimplementedApplicationServer
}

func RegisterApplicationRpcServer(s *grpc.Server) {
	pb.RegisterApplicationServer(s, &ApplicationRpcServer{})
}

func (s *ApplicationRpcServer) GetApplicationById(ctx context.Context, in *pb.GetApplicationByIdRequest) (*pb.GetApplicationByIdReply, error) {
	application, e := applicationService.ApplicationWithId(in.GetId())
	if e != nil {
		return nil, e
	}
	b, e := json.Marshal(application)
	if e != nil {
		return nil, e
	}
	var m map[string]interface{}
	e = json.Unmarshal(b, &m)
	if e != nil {
		return nil, e
	}
	a, err := structpb.NewStruct(m)
	if err != nil {
		return nil, e
	}

	log.Printf("Received application id: %v", in.GetId())
	return &pb.GetApplicationByIdReply{Application: a}, nil
}