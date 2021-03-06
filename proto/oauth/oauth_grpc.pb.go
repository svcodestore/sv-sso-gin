// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: oauth.proto

package oauth

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OauthClient is the client API for Oauth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OauthClient interface {
	GetGrantCode(ctx context.Context, in *GetGrantCodeRequest, opts ...grpc.CallOption) (*GetGrantCodeReply, error)
	GetOauthCode(ctx context.Context, in *GetOauthCodeRequest, opts ...grpc.CallOption) (*GetOauthCodeReply, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error)
	Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutReply, error)
	IsUserLogin(ctx context.Context, in *IsUserLoginRequest, opts ...grpc.CallOption) (*IsUserLoginReply, error)
}

type oauthClient struct {
	cc grpc.ClientConnInterface
}

func NewOauthClient(cc grpc.ClientConnInterface) OauthClient {
	return &oauthClient{cc}
}

func (c *oauthClient) GetGrantCode(ctx context.Context, in *GetGrantCodeRequest, opts ...grpc.CallOption) (*GetGrantCodeReply, error) {
	out := new(GetGrantCodeReply)
	err := c.cc.Invoke(ctx, "/oauth.Oauth/GetGrantCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oauthClient) GetOauthCode(ctx context.Context, in *GetOauthCodeRequest, opts ...grpc.CallOption) (*GetOauthCodeReply, error) {
	out := new(GetOauthCodeReply)
	err := c.cc.Invoke(ctx, "/oauth.Oauth/GetOauthCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oauthClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := c.cc.Invoke(ctx, "/oauth.Oauth/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oauthClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutReply, error) {
	out := new(LogoutReply)
	err := c.cc.Invoke(ctx, "/oauth.Oauth/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oauthClient) IsUserLogin(ctx context.Context, in *IsUserLoginRequest, opts ...grpc.CallOption) (*IsUserLoginReply, error) {
	out := new(IsUserLoginReply)
	err := c.cc.Invoke(ctx, "/oauth.Oauth/IsUserLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OauthServer is the server API for Oauth service.
// All implementations must embed UnimplementedOauthServer
// for forward compatibility
type OauthServer interface {
	GetGrantCode(context.Context, *GetGrantCodeRequest) (*GetGrantCodeReply, error)
	GetOauthCode(context.Context, *GetOauthCodeRequest) (*GetOauthCodeReply, error)
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	Logout(context.Context, *LogoutRequest) (*LogoutReply, error)
	IsUserLogin(context.Context, *IsUserLoginRequest) (*IsUserLoginReply, error)
	mustEmbedUnimplementedOauthServer()
}

// UnimplementedOauthServer must be embedded to have forward compatible implementations.
type UnimplementedOauthServer struct {
}

func (UnimplementedOauthServer) GetGrantCode(context.Context, *GetGrantCodeRequest) (*GetGrantCodeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGrantCode not implemented")
}
func (UnimplementedOauthServer) GetOauthCode(context.Context, *GetOauthCodeRequest) (*GetOauthCodeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOauthCode not implemented")
}
func (UnimplementedOauthServer) Login(context.Context, *LoginRequest) (*LoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedOauthServer) Logout(context.Context, *LogoutRequest) (*LogoutReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedOauthServer) IsUserLogin(context.Context, *IsUserLoginRequest) (*IsUserLoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsUserLogin not implemented")
}
func (UnimplementedOauthServer) mustEmbedUnimplementedOauthServer() {}

// UnsafeOauthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OauthServer will
// result in compilation errors.
type UnsafeOauthServer interface {
	mustEmbedUnimplementedOauthServer()
}

func RegisterOauthServer(s grpc.ServiceRegistrar, srv OauthServer) {
	s.RegisterService(&Oauth_ServiceDesc, srv)
}

func _Oauth_GetGrantCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGrantCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OauthServer).GetGrantCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/oauth.Oauth/GetGrantCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OauthServer).GetGrantCode(ctx, req.(*GetGrantCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Oauth_GetOauthCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOauthCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OauthServer).GetOauthCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/oauth.Oauth/GetOauthCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OauthServer).GetOauthCode(ctx, req.(*GetOauthCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Oauth_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OauthServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/oauth.Oauth/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OauthServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Oauth_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OauthServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/oauth.Oauth/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OauthServer).Logout(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Oauth_IsUserLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsUserLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OauthServer).IsUserLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/oauth.Oauth/IsUserLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OauthServer).IsUserLogin(ctx, req.(*IsUserLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Oauth_ServiceDesc is the grpc.ServiceDesc for Oauth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Oauth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "oauth.Oauth",
	HandlerType: (*OauthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGrantCode",
			Handler:    _Oauth_GetGrantCode_Handler,
		},
		{
			MethodName: "GetOauthCode",
			Handler:    _Oauth_GetOauthCode_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Oauth_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _Oauth_Logout_Handler,
		},
		{
			MethodName: "IsUserLogin",
			Handler:    _Oauth_IsUserLogin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "oauth.proto",
}
