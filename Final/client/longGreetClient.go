package main

import (
	"context"
	"fmt"
	greetpb "grpc-dev/proto/greet"
	"log"
	"time"
)

func LongGreet(c greetpb.GreetServiceClient) {
	fmt.Println("Client Streaming initiated")
	// Create a slice of requests to send to the server using the client stream
	// This is a client streaming function and we are sending multiple requests to the server
	requests := []*greetpb.LongGreetRequest{
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Abhishek",
				LastName:  "Kumar",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Ben",
				LastName:  "Stokes",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Ben",
				LastName:  "Foakes",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Adam",
				LastName:  "Jonas",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "John",
				LastName:  "Doe",
			},
		},
	}

	clientStream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while calling LongGreet RPC: %v \n", err)
	}
	// Send each request to the server
	// This is a client streaming function and we are sending multiple requests to the server
	for _, req := range requests {
		clientStream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}

	res, err := clientStream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from LongGreet: %v \n", err)
	}
	fmt.Println("Long greet response: ", res)

}
