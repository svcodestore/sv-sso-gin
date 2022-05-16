package application

import (
	pb "github.com/svcodestore/sv-sso-gin/proto/application"
	"github.com/svcodestore/sv-sso-gin/service"
	"github.com/svcodestore/sv-sso-gin/utils"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	applicationService          = service.ServiceGroup.ApplicationService
	privilegeApplicationService = service.ServiceGroup.PrivilegeApplicationService
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
	a := utils.ToRpcStruct(application)

	return &pb.GetApplicationByIdReply{Application: a}, nil
}

func (s *ApplicationRpcServer) GetApplicationSecretByClientId(ctx context.Context, in *pb.GetApplicationSecretByClientIdRequest) (*pb.GetApplicationSecretByClientIdReply, error) {
	clientSecret, e := applicationService.ApplicationClientSecretWithClientId(in.GetClientId())
	if e != nil {
		return nil, e
	}

	return &pb.GetApplicationSecretByClientIdReply{ClientSecret: clientSecret}, nil
}

func (s *ApplicationRpcServer) GetApplicationsByOrganizationId(ctx context.Context, in *pb.GetApplicationsByOrganizationIdRequest) (*pb.GetApplicationsByOrganizationIdReply, error) {
	applications, e := applicationService.ApplicationsWithOrganizationIds(in.GetOrganizationId())
	if e != nil {
		return nil, e
	}

	a := utils.ToRpcStruct(map[string]interface{}{
		"data": applications,
	})

	return &pb.GetApplicationsByOrganizationIdReply{
		Applications: a,
	}, nil
}

func (s *ApplicationRpcServer) GetAvailableApplications(ctx context.Context, in *pb.GetAvailableApplicationsRequest) (*pb.GetAvailableApplicationsReply, error) {
	applications, e := applicationService.AvailableApplications()
	if e != nil {
		return nil, e
	}

	a := utils.ToRpcStruct(map[string]interface{}{
		"data": applications,
	})
	return &pb.GetAvailableApplicationsReply{
		Applications: a,
	}, nil
}
