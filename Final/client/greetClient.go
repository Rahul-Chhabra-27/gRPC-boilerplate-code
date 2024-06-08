package main

import (
	"context"
	"fmt"
	greetpb "grpc-dev/proto/greet"
	"log"
)

func Greet(c greetpb.GreetServiceClient) {

	fmt.Println("Unary RPC initiated")

	// Create a GreetRequest
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Ben",
			LastName:  "Stokes",
		},
	}
	
	// Call the Greet function on the client
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Greet rpc: %v \n", err)
	}

	log.Printf("Response: %v\n", res)
}
