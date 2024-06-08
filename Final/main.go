package main

import (
	"fmt"
	greetpb "grpc-dev/proto/greet"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GreetService struct {
	greetpb.UnimplementedGreetServiceServer
}

func main() {

	fmt.Println("Hello, I'm the server")

	// Create a listener on TCP port 50055
	lis, err := net.Listen("tcp", "0.0.0.0:50055")
	if err != nil {
		log.Fatalf("Failed to listen: %v ", err)
	}

	// Create a gRPC server
	s := grpc.NewServer()
	// Register the greet service
	greetpb.RegisterGreetServiceServer(s, &GreetService{})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	// Serve the gRPC server
	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
