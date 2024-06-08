package main

import (
	"context"
	"fmt"
	greetpb "grpc-dev/proto/greet"
)

func (*GreetService) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {

	fmt.Printf("Greet func was invoked with %v: \n", req)

	// Get the first name from the client request
	firstName := req.GetGreeting().GetFirstName()
	// Get the last name from the client request
	lastName := req.GetGreeting().GetLastName()
	// Create the response
	result := "Hello, " + firstName + " " + lastName
	res := &greetpb.GreetResponse{
		Result: result,
	}
	// Return the response
	return res, nil
}
