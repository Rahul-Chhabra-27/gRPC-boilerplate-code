package main

import (
	"fmt"
	greetpb "grpc-dev/proto/greet"
	"strconv"
	"time"
)

func (*GreetService) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {

	fmt.Println("GreetManyTimes function was invoked with a request", req)
	// Get the first name from the client request
	firstName := req.GetGreeting().GetFirstName()
	lastName := req.GetGreeting().GetLastName()

	// Send a response to the client for 10 times
	for i := 0; i < 10; i++ {
		result := "Hello, " + firstName + " " + lastName + " | " + strconv.Itoa(i)
		res := &greetpb.GreetManyTimesResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}
