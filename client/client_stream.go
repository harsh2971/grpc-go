package main

import (
	"context"
	"log"
	"time"

	pb "github.com/harsh2971/grpc-go/proto"
)

func callSayHelloClientStreaming(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Client streaming started")
	// client would send a stream of requests to the server
	// server would respond with a response
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.SayHelloClientStreaming(ctx)
	if err != nil {
		log.Fatalf("clould not call SayHelloClientStreaming: %v", err)
	}

	for _, name := range names.Names {
		req:= &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err!= nil {
			log.Fatalf("error while sending request: %v", err)
		}
		log.Printf("Sent request with name: %s", name)
		time.Sleep(2*time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while closing and receiving: %v", err)
	}
	log.Printf("Response from SayHelloClientStreaming: %v", res.Messages)
	log.Printf("Client streaming finished")

}