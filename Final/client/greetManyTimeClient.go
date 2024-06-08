package main

import (
	"context"
	"fmt"
	greetpb "grpc-dev/proto/greet"
	"io"
	"log"
)

func GreetManyTimes(c greetpb.GreetServiceClient) {
	fmt.Println("Server Streaming started.")

	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "John",
			LastName:  "Doe",
		},
	}

	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes rpc: %v \n", err)
	}
    // Receive the server stream of responses from the server and print each response
	// This is a server streaming function
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			fmt.Println("Reached Stream EOF")
			break
		}
		if err != nil {
			log.Fatalf("Error receiving from the server stream: %v\n", err)
		}
		log.Printf("Response from GreetManyTimes: %v\n", msg.GetResult())
	}

}
