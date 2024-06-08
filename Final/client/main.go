package main

import (
	"fmt"
	greetpb "grpc-dev/proto/greet"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello, I'm a client")

	// Listener....
	connection, err := grpc.Dial("localhost:50055", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("could not connect %v", err)
		return
	}
	defer connection.Close()

	// GRPC client server..
	grpcClientServer := greetpb.NewGreetServiceClient(connection)
	
	//Greet(grpcClientServer);
	//GreetManyTimes(grpcClientServer);
	//GreetEveryone(grpcClientServer);
	LongGreet(grpcClientServer);
}
