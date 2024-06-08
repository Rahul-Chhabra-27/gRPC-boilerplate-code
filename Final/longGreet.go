package main

import (
	"fmt"
	"io"
	"log"
	greetpb "grpc-dev/proto/greet"
)

func (*GreetService) LongGreet(stream greetpb.GreetService_LongGreetServer) error {

	fmt.Printf("LongGreet func was invoked\n")
	result := ""
	// Receive the client stream of requests from the client and respond to each request
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Printf("Reached end of stream EOF: %v ", err)
			return stream.SendAndClose(&greetpb.LongGreetResponse{
				Result: result,
			})
		}
		if err != nil {
			log.Fatalf("Error receiving from the client stream: %v ", err)
		}
		// Get the first name from the client request
		firstName := req.GetGreeting().GetFirstName()
		result += "Hello " + firstName + "!\n"
		fmt.Printf("Received message: %v\n", firstName)
	}
}
