package main

import (
	"fmt"
	greetpb "grpc-dev/proto/greet"
	"io"
	"log"
)

func (*GreetService) GreetEveryone(stream greetpb.GreetService_GreetEveryoneServer) error {
	fmt.Printf("GreetEveryone func invoked\n")
	// Receive the client stream of requests from the client and respond to each request
	// This is a server streaming function
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Println("Reached EOF")
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v ", err)
			return err
		}
		firstName := req.GetGreeting().GetFirstName()
		result := "Hello, " + firstName + ". "

		sendErr := stream.Send(&greetpb.GreetEveryoneResponse{
			Result: result,
		})
		if sendErr != nil {
			log.Fatalf("Error while sending data to client: %v ", err)
			return err
		}
	}
}
