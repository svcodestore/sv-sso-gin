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

func (s *OauthRpcServer) GetOauthCode(ctx context.Context, in *pb.GetOauthCodeRequest) (*pb.GetOauthCodeReply, error) {
	grantType := in.GetGrantType()
	clientId := in.GetClientId()
	clientSecret := in.GetClientSecret()
	code := in.GetCode()
	redirectUri := in.GetRedirectUri()

	if grantType != "authorization_code" {
		return &pb.GetOauthCodeReply{OauthInfo: utils.ToRpcStruct(reply.FailWithDetail(nil, "grant type error"))}, nil
	}
	if code == "" {
		return &pb.GetOauthCodeReply{OauthInfo: utils.ToRpcStruct(reply.FailWithDetail(nil, "empty code"))}, nil
	}
	accessToken, refreshToken, user, err := oauthService.DoGenerateOauthCode(clientId, clientSecret, code, redirectUri)
	if err != nil {
		return &pb.GetOauthCodeReply{OauthInfo: utils.ToRpcStruct(reply.FailWithDetail(nil, err.Error()))}, nil
	}
	info := map[string]interface{}{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
		"user":         user,
	}
	return &pb.GetOauthCodeReply{OauthInfo: utils.ToRpcStruct(reply.OkWithData(info))}, nil
}

func (s *OauthRpcServer) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginReply, error) {
	username := in.GetUsername()
	password := in.GetPassword()
	loginType := in.GetLoginType()
	clientId := in.GetClientId()

	accessToken, refreshToken, user, err := oauthService.DoOauthLogin(username, password, loginType, clientId)

	if err != nil {
		return &pb.LoginReply{OauthInfo: utils.ToRpcStruct(reply.FailWithDetail(nil, err.Error()))}, nil
	}
	info := map[string]interface{}{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
		"user":         user,
	}
	return &pb.LoginReply{OauthInfo: utils.ToRpcStruct(reply.OkWithData(info))}, nil
}

func (s OauthRpcServer) Logout(ctx context.Context, in *pb.LogoutRequest) (*pb.LogoutReply, error) {
	token := in.GetAccessToken()
	j := utils.NewJWT()
	claims, _ := j.ParseToken(token)
	affected, err := oauthService.DeleteAccessTokenFromRedis(claims.UserId)
	if err == nil && affected > 0 {
		return &pb.LogoutReply{LogoutResult: utils.ToRpcStruct(reply.Ok())}, nil
	}
	return &pb.LogoutReply{LogoutResult: utils.ToRpcStruct(reply.Fail())}, nil
}
