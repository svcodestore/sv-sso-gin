package oauth

import (
	"github.com/svcodestore/sv-sso-gin/model/rpc/reply"
	pb "github.com/svcodestore/sv-sso-gin/proto/oauth"
	"github.com/svcodestore/sv-sso-gin/service"
	"github.com/svcodestore/sv-sso-gin/utils"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	oauthService = service.ServiceGroup.OauthService
)

type OauthRpcServer struct {
	pb.UnimplementedOauthServer
}

func RegisterOauthRpcServer(s *grpc.Server) {
	pb.RegisterOauthServer(s, &OauthRpcServer{})
}

func (s *OauthRpcServer) GetGrantCode(ctx context.Context, in *pb.GetGrantCodeRequest) (*pb.GetGrantCodeReply, error) {
	token := in.GetToken()
	j := utils.NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		return &pb.GetGrantCodeReply{GrantCode: utils.ToRpcStruct(reply.FailWithDetail(nil, err.Error()))}, nil
	}
	grantedCode, err := oauthService.DoGenerateGrantCode(claims.UserId, in.GetClientId())
	if err != nil {
		return &pb.GetGrantCodeReply{GrantCode: utils.ToRpcStruct(reply.FailWithDetail(nil, err.Error()))}, nil
	}
	return &pb.GetGrantCodeReply{GrantCode: utils.ToRpcStruct(reply.OkWithData(grantedCode))}, nil
}


