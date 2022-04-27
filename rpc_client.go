package main

import (
	"context"
	pb "github.com/svcodestore/sv-sso-gin/proto/application"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"time"
)

func main() {
	address     := ":50088"
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewApplicationClient(conn)

	// Contact the server and print out its response.
	id := "0"
	if len(os.Args) > 1 {
		id = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GetApplicationById(ctx, &pb.GetApplicationByIdRequest{Id: id})
	if err != nil {
		log.Fatalf("could not get user: %v", err)
	}
	reply := r.GetApplication()
	b,_:=reply.MarshalJSON()
	log.Printf("user: %v", string(b))
}