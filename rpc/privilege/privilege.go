package privilege

import (
	"context"
	pb "github.com/svcodestore/sv-sso-gin/proto/privilege"
	"github.com/svcodestore/sv-sso-gin/utils"
	"google.golang.org/grpc"
)

type PrivilegeRpcServer struct {
	pb.UnimplementedPrivilegeServer
}

func RegisterPrivilegeRpcServer(s *grpc.Server) {
	pb.RegisterPrivilegeServer(s, &PrivilegeRpcServer{})
}

func (receiver PrivilegeRpcServer) GetAccessibleApplicationsByUserId(ctx context.Context, in *pb.GetAccessibleApplicationsByUserIdRequest) (*pb.GetAccessibleApplicationsByUserIdReply, error) {
	applications, _, e := privilegeApplicationService.AccessibleApplications(in.GetUserId())
	if e != nil {
		return nil, e
	}
	apps := utils.ToRpcStruct(applications)

	return &pb.GetAccessibleApplicationsByUserIdReply{Applications: apps}, nil
}
