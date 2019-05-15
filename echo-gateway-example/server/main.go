package main

import (
	"context"
	"log"
	"net"

	pb "github.com/agguzman/grpc-examples/echo-gateway-example"
	grpc "google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) Echo(ctx context.Context, in *pb.SimpleMessage) (*pb.SimpleMessage, error) {
	return &pb.SimpleMessage{Msg: "echo... " + in.Msg}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	pb.RegisterEchoServer(srv, &server{})
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
