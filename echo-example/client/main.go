package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/agguzman/grpc-examples/echo-example"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewEchoClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Echo(ctx, &pb.Request{Message: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf(r.Message)
}
