package utils

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/svcodestore/sv-sso-gin/global"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/structpb"
	"log"
	"time"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

func ToRpcStruct(data interface{}) *structpb.Struct {
	b, e := json.Marshal(data)
	if e != nil {
		return nil
	}
	var m map[string]interface{}
	e = json.Unmarshal(b, &m)
	if e != nil {
		return nil
	}
	s, err := structpb.NewStruct(m)
	if err != nil {
		return nil
	}
	return s
}

func CallAuthRpcService(cb func(conn *grpc.ClientConn, ctx context.Context) (reply *structpb.Struct, e error)) (data map[string]interface{}, err error) {
	addr := global.CONFIG.System.AuthRpcAddr
	if addr == "" {
		panic("sso server address is empty")
	}

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	defer conn.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return
	}

	reply, e := cb(conn, ctx)
	if e != nil {
		return
	}
	data = reply.AsMap()

	return
}
