package main

import (
	"context"
	"log"
	"net"

	pb "github.com/agguzman/grpc-examples/echo-example"
	grpc "google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) Echo(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	return &pb.Response{Message: "echoing " + in.Message}, nil
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
