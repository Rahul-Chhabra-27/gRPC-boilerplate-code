package main

import (
	"context"
	"fmt"
	greetpb "grpc-dev/proto/greet"
	"io"
	"log"
	"time"
)

func GreetEveryone(c greetpb.GreetServiceClient) {
	fmt.Println("BiDi Streaming initiated")

	requests := []*greetpb.GreetEveryoneRequest{
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Stuart",
				LastName:  "Broad",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Paul",
				LastName:  "Collingwood",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Ben",
				LastName:  "Stokes",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Ben",
				LastName:  "Foakes",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Mark",
				LastName:  "Wood",
			},
		},
	}

	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while calling GreetEveryone RPC: %v \n", err)
		return
	}

	waitChannel := make(chan struct{})

	// Send messages to client
	go func() {
		for _, req := range requests {
			fmt.Println("[SENDING] Message: ", req)
			stream.Send(req)
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()
	// Receive messages from client
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				log.Println("Reached EOF")
				break
			}
			if err != nil {
				log.Fatalf("Error while reading server stream: %v ", err)
				break
			}
			fmt.Println("[RECEIVING] Message: ", res)
		}
		close(waitChannel)
	}()

	<-waitChannel
}
