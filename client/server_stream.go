package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/harsh2971/grpc-go/proto"
)

// callSayHelloServerStreaming is a function to call the SayHelloServerStreaming method
func callSayHelloServerStreaming(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Streaming started")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) // Increased timeout to handle all messages
	defer cancel()

	stream, err := client.SayHelloServerStreaming(ctx, names)
	if err != nil {
		log.Fatalf("error while calling SayHelloServerStreaming: %v", err)
	}

	for {
		// receive a stream of responses from the server
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while receiving message: %v", err)
		}
		log.Printf("Received: %s", message.Message)
	}
	log.Printf("Streaming finished")

}
