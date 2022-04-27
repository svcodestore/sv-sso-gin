package initialize

import (
	"fmt"
	"github.com/svcodestore/sv-sso-gin/global"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"log"
	"net"
	"time"
)

func InitGrpc(registerPb func(server *grpc.Server)) (s *grpc.Server, err error) {
	addr := fmt.Sprintf(":%d", global.CONFIG.System.RpcAddr)
	fmt.Println(addr, "rpc addr")
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}
	s = grpc.NewServer(grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle: 5 * time.Minute,
	}))
	registerPb(s)
	log.Printf("rpc server listening at %v", lis.Addr())
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	return
}
